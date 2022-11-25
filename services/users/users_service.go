package services

import (
	"errors"
	"vaksin-id-be/dto/payload"
	"vaksin-id-be/dto/response"
	m "vaksin-id-be/middleware"
	"vaksin-id-be/model"
	mysql "vaksin-id-be/repository/mysql/users"
	"vaksin-id-be/util"

	"github.com/google/uuid"
)

type UserService interface {
	RegisterUser(payloads payload.RegisterUser) error
	LoginUser(payloads payload.Login) (response.Login, error)
	GetUserDataByNik(nik string) (response.UserProfile, error)
	UpdateUserProfile(payloads payload.UpdateUser, nik string) error
	DeleteUserProfile(nik string) error
}

type userService struct {
	UserRepo mysql.UserRepository
}

func NewUserService(userRepo mysql.UserRepository) *userService {
	return &userService{
		UserRepo: userRepo,
	}
}

func (u *userService) RegisterUser(payload payload.RegisterUser) error {

	hashPass, err := util.HashPassword(payload.Password)
	if err != nil {
		return err
	}

	defaultVaccineCount := 0

	userModel := model.Users{
		NIK:          payload.NikUser,
		Email:        payload.Email,
		Password:     hashPass,
		Fullname:     payload.Fullname,
		PhoneNum:     payload.PhoneNum,
		Gender:       payload.Gender,
		VaccineCount: defaultVaccineCount,
		BirthDate:    payload.BirthDate,
	}

	errRegis := u.UserRepo.RegisterUser(userModel)
	if errRegis != nil {
		return errRegis
	}

	idAddr := uuid.NewString()

	userAddr := model.Addresses{
		ID:      idAddr,
		NikUser: payload.NikUser,
	}

	errAddr := u.UserRepo.CreateAddress(userAddr)
	if errAddr != nil {
		return errAddr
	}

	return nil
}

func (u *userService) LoginUser(payload payload.Login) (response.Login, error) {
	var loginResponse response.Login

	userModel := model.Users{
		Email:    payload.Email,
		Password: payload.Password,
	}

	userData, err := u.UserRepo.LoginUser(userModel)
	if err != nil {
		return loginResponse, err
	}

	isValid := util.CheckPasswordHash(payload.Password, userData.Password)
	if !isValid {
		return loginResponse, errors.New("wrong password")
	}

	token, errToken := m.CreateToken(userData.NIK, userData.Email)

	if errToken != nil {
		return loginResponse, err
	}

	loginResponse = response.Login{
		Token: token,
	}

	return loginResponse, nil
}

func (u *userService) GetUserDataByNik(nik string) (response.UserProfile, error) {
	var responseUser response.UserProfile

	getData, err := u.UserRepo.GetUserDataByNik(nik)
	if err != nil {
		return responseUser, err
	}

	ageUser, err := u.UserRepo.GetAgeUser(getData)
	if err != nil {
		return responseUser, err
	}

	responseUser = response.UserProfile{
		NIK:          getData.NIK,
		Email:        getData.Email,
		Fullname:     getData.Fullname,
		PhoneNum:     getData.PhoneNum,
		Gender:       getData.Gender,
		VaccineCount: getData.VaccineCount,
		Age:          ageUser.Age,
	}

	return responseUser, nil
}

func (u *userService) UpdateUserProfile(payloads payload.UpdateUser, nik string) error {
	userNik, err := m.GetUserNik(nik)
	if err != nil {
		return err
	}

	dataUser := model.Users{
		Fullname:  payloads.Fullname,
		NIK:       userNik,
		Email:     payloads.Email,
		Gender:    payloads.Gender,
		PhoneNum:  payloads.PhoneNum,
		BirthDate: payloads.BirthDate,
	}

	if err := u.UserRepo.UpdateUserProfile(dataUser); err != nil {
		return err
	}

	return nil
}

func (u *userService) DeleteUserProfile(nik string) error {
	if err := u.UserRepo.DeleteAddress(nik); err != nil {
		return err
	}

	if err := u.UserRepo.DeleteUser(nik); err != nil {
		return err
	}

	return nil
}
