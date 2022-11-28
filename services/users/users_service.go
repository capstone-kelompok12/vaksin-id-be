package services

import (
	"errors"
	"time"
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
	GetAddressUser(nik string) (model.Addresses, error)
	UpdateUserAddress(payloads payload.UpdateAddress, nik string) error
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

func (u *userService) RegisterUser(payloads payload.RegisterUser) error {

	hashPass, err := util.HashPassword(payloads.Password)
	if err != nil {
		return err
	}

	defaultVaccineCount := 0

	dateBirth, err := time.Parse("2006-01-02", payloads.BirthDate)
	if err != nil {
		return err
	}
	data, err := u.UserRepo.CheckExistNik(payloads.NikUser)
	if err != nil {
		return err
	}

	if data.NIK != "" && data.DeletedAt != nil {
		userModel := model.Users{
			Email:     payloads.Email,
			Password:  hashPass,
			Fullname:  payloads.Fullname,
			PhoneNum:  payloads.PhoneNum,
			Gender:    payloads.Gender,
			BirthDate: dateBirth,
		}
		err = u.UserRepo.ReactivatedUser(payloads.NikUser)
		if err != nil {
			return err
		}
		err = u.UserRepo.ReactivatedUpdateUser(userModel, payloads.NikUser)
		if err != nil {
			return err
		}
		err = u.UserRepo.ReactivatedAddress(payloads.NikUser)
		if err != nil {
			return err
		}
		return nil
	}

	userModel := model.Users{
		NIK:          payloads.NikUser,
		Email:        payloads.Email,
		Password:     hashPass,
		Fullname:     payloads.Fullname,
		PhoneNum:     payloads.PhoneNum,
		Gender:       payloads.Gender,
		VaccineCount: defaultVaccineCount,
		BirthDate:    dateBirth,
	}

	errRegis := u.UserRepo.RegisterUser(userModel)
	if errRegis != nil {
		return errRegis
	}

	idAddr := uuid.NewString()

	userAddr := model.Addresses{
		ID:                 idAddr,
		IdHealthFacilities: nil,
		NikUser:            &payloads.NikUser,
	}

	errAddr := u.UserRepo.CreateAddress(userAddr)
	if errAddr != nil {
		return errAddr
	}

	return nil
}

func (u *userService) LoginUser(payloads payload.Login) (response.Login, error) {
	var loginResponse response.Login

	userModel := model.Users{
		Email:    payloads.Email,
		Password: payloads.Password,
	}

	userData, err := u.UserRepo.LoginUser(userModel)
	if err != nil {
		return loginResponse, err
	}

	isValid := util.CheckPasswordHash(payloads.Password, userData.Password)
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

	getUserNik, err := m.GetUserNik(nik)
	if err != nil {
		return responseUser, err
	}

	getData, err := u.UserRepo.GetUserDataByNik(getUserNik)
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

	dateBirth, err := time.Parse("2006-01-02", payloads.BirthDate)
	if err != nil {
		return err
	}

	dataUser := model.Users{
		Fullname:  payloads.Fullname,
		NIK:       userNik,
		Email:     payloads.Email,
		Gender:    payloads.Gender,
		PhoneNum:  payloads.PhoneNum,
		BirthDate: dateBirth,
	}

	if err := u.UserRepo.UpdateUserProfile(dataUser); err != nil {
		return err
	}

	return nil
}

func (u *userService) GetAddressUser(nik string) (model.Addresses, error) {
	var address model.Addresses

	getUserNik, err := m.GetUserNik(nik)
	if err != nil {
		return address, err
	}

	dataAddress, err := u.UserRepo.GetAddress(getUserNik)
	if err != nil {
		return address, err
	}

	return dataAddress, nil
}

func (u *userService) UpdateUserAddress(payloads payload.UpdateAddress, nik string) error {

	getUserNik, err := m.GetUserNik(nik)
	if err != nil {
		return err
	}

	newAddress := model.Addresses{
		CurrentAddress: payloads.CurrentAddress,
		District:       payloads.District,
		City:           payloads.City,
		Province:       payloads.Province,
		Longitude:      payloads.Longitude,
		Latitude:       payloads.Latitude,
	}

	if err := u.UserRepo.UpdateAddress(newAddress, getUserNik); err != nil {
		return err
	}
	return nil
}

func (u *userService) DeleteUserProfile(nik string) error {
	getUserNik, err := m.GetUserNik(nik)
	if err != nil {
		return err
	}

	if err := u.UserRepo.DeleteAddress(getUserNik); err != nil {
		return err
	}

	if err := u.UserRepo.DeleteUser(getUserNik); err != nil {
		return err
	}

	return nil
}
