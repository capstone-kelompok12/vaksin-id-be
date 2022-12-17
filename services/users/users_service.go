package services

import (
	"errors"
	"sort"
	"time"
	"vaksin-id-be/dto/payload"
	"vaksin-id-be/dto/response"
	m "vaksin-id-be/middleware"
	"vaksin-id-be/model"
	mysqla "vaksin-id-be/repository/mysql/addresses"
	mysqlb "vaksin-id-be/repository/mysql/bookings"
	mysqlh "vaksin-id-be/repository/mysql/health_facilities"
	mysqlhs "vaksin-id-be/repository/mysql/histories"
	mysqls "vaksin-id-be/repository/mysql/sessions"
	mysqlu "vaksin-id-be/repository/mysql/users"
	"vaksin-id-be/util"

	"github.com/google/uuid"
)

type UserService interface {
	RegisterUser(payloads payload.RegisterUser) error
	LoginUser(payloads payload.Login) (response.Login, error)
	GetUserDataByNik(nik string) (response.UserProfile, error)
	GetUserDataByNikNoAddress(nik string) (response.UserProfile, error)
	GetUserHistory(nik string) (response.UserHistory, error)
	UpdateUserProfile(payloads payload.UpdateUser, nik string) (response.UpdateUser, error)
	DeleteUserProfile(nik string) error
	NearbyHealthFacilities(payloads payload.NearbyHealth, nik string) (response.UserNearbyHealth, error)
}

type userService struct {
	UserRepo    mysqlu.UserRepository
	AddressRepo mysqla.AddressesRepository
	HealthRepo  mysqlh.HealthFacilitiesRepository
	HistoryRepo mysqlhs.HistoriesRepository
	BookingRepo mysqlb.BookingRepository
	SessionRepo mysqls.SessionsRepository
}

func NewUserService(
	userRepo mysqlu.UserRepository,
	addressRepo mysqla.AddressesRepository,
	healthRepo mysqlh.HealthFacilitiesRepository,
	historyRepo mysqlhs.HistoriesRepository,
	bookingRepo mysqlb.BookingRepository,
	sessionRepo mysqls.SessionsRepository,
) *userService {
	return &userService{
		UserRepo:    userRepo,
		AddressRepo: addressRepo,
		HealthRepo:  healthRepo,
		HistoryRepo: historyRepo,
		BookingRepo: bookingRepo,
		SessionRepo: sessionRepo,
	}
}

func (u *userService) RegisterUser(payloads payload.RegisterUser) error {

	hashPass, err := util.HashPassword(payloads.Password)
	if err != nil {
		return err
	}

	if payloads.Gender != "P" && payloads.Gender != "L" {
		return errors.New("input gender must P or L")
	}

	defaultVaccineCount := 0

	dateBirth, err := time.Parse("2006-01-02", payloads.BirthDate)
	if err != nil {
		return err
	}
	data, _ := u.UserRepo.CheckExistNik(payloads.NikUser)

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
		ProfileImage: nil,
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

	errAddr := u.AddressRepo.CreateAddress(userAddr)
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
		BirthDate:    getData.BirthDate,
		Age:          ageUser.Age,
		Address:      getData.Address,
	}

	return responseUser, nil
}

func (u *userService) GetUserDataByNikNoAddress(nik string) (response.UserProfile, error) {
	var responseUser response.UserProfile

	getData, err := u.UserRepo.GetUserDataByNikNoAddress(nik)
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
		BirthDate:    getData.BirthDate,
		Age:          ageUser.Age,
	}

	return responseUser, nil
}

func (u *userService) GetUserHistory(nik string) (response.UserHistory, error) {
	var historyUser response.UserHistory

	nikUser, err := m.GetUserNik(nik)
	if err != nil {
		return historyUser, err
	}

	getData, err := u.UserRepo.GetUserHistoryByNik(nikUser)
	if err != nil {
		return historyUser, err
	}

	getDataUser, err := u.UserRepo.GetUserDataByNikNoAddress(nikUser)
	if err != nil {
		return historyUser, err
	}

	ageUser, err := u.UserRepo.GetAgeUser(getDataUser)
	if err != nil {
		return historyUser, err
	}

	countHistoryUser, err := u.HistoryRepo.GetHistoriesByNIK(nikUser)
	if err != nil {
		return historyUser, err
	}

	historyList := make([]response.HistoryCustomUser, len(countHistoryUser))

	for i, val := range countHistoryUser {
		dataBooking, err := u.BookingRepo.GetBooking(val.IdBooking)
		if err != nil {
			return historyUser, err
		}

		GetHealthFacil, err := u.HealthRepo.GetHealthFacilitiesById(dataBooking.Session.Vaccine.IdHealthFacilities)
		if err != nil {
			return historyUser, err
		}

		dataHealthFacilities := response.HealthFacilitiesCustomUser{
			ID:        GetHealthFacil.ID,
			Email:     GetHealthFacil.Email,
			PhoneNum:  GetHealthFacil.PhoneNum,
			Name:      GetHealthFacil.Name,
			Image:     GetHealthFacil.Image,
			CreatedAt: GetHealthFacil.CreatedAt,
			UpdatedAt: GetHealthFacil.UpdatedAt,
			Address:   GetHealthFacil.Address,
		}

		bookingLoop := response.BookingHistoryLoop{
			ID:               dataBooking.ID,
			IdSession:        dataBooking.IdSession,
			NikUser:          dataBooking.NikUser,
			Queue:            &dataBooking.Queue,
			Status:           dataBooking.Status,
			CreatedAt:        dataBooking.CreatedAt,
			UpdatedAt:        dataBooking.UpdatedAt,
			Session:          *dataBooking.Session,
			HealthFacilities: dataHealthFacilities,
		}

		historyList[i] = response.HistoryCustomUser{
			ID:         val.ID,
			IdBooking:  val.IdBooking,
			NikUser:    val.NikUser,
			IdSameBook: val.IdSameBook,
			Status:     val.Status,
			CreatedAt:  val.CreatedAt,
			UpdatedAt:  val.UpdatedAt,
			Booking:    bookingLoop,
		}
	}

	historyUser = response.UserHistory{
		NIK:          getData.NIK,
		Email:        getData.Email,
		Fullname:     getData.Fullname,
		PhoneNum:     getData.PhoneNum,
		Gender:       getData.Gender,
		VaccineCount: getData.VaccineCount,
		BirthDate:    getData.BirthDate,
		Age:          ageUser.Age,
		Address:      getData.Address,
		History:      historyList,
	}

	return historyUser, nil
}

func (u *userService) UpdateUserProfile(payloads payload.UpdateUser, nik string) (response.UpdateUser, error) {
	var dataResp response.UpdateUser

	getNikUser, err := m.GetUserNik(nik)
	if err != nil {
		return dataResp, err
	}

	if payloads.Gender != "P" && payloads.Gender != "L" {
		return dataResp, errors.New("input gender must P or L")
	}

	dateBirth, err := time.Parse("2006-01-02", payloads.BirthDate)
	if err != nil {
		return dataResp, err
	}

	hashPass, err := util.HashPassword(payloads.Password)
	if err != nil {
		return dataResp, err
	}

	data, err := u.GetUserDataByNikNoAddress(getNikUser)
	if err != nil {
		return dataResp, err
	}

	dataUser := model.Users{
		NIK:       payloads.NikUser,
		Email:     payloads.Email,
		Password:  hashPass,
		Fullname:  payloads.Fullname,
		PhoneNum:  payloads.PhoneNum,
		Gender:    payloads.Gender,
		BirthDate: dateBirth,
	}

	if err := u.UserRepo.UpdateUserProfile(dataUser); err != nil {
		return dataResp, err
	}

	dataResp = response.UpdateUser{
		Fullname:  data.Fullname,
		NikUser:   data.NIK,
		Email:     data.Email,
		Password:  hashPass,
		PhoneNum:  data.PhoneNum,
		Gender:    data.Gender,
		BirthDate: dateBirth,
	}

	return dataResp, nil
}

func (u *userService) DeleteUserProfile(nik string) error {
	getUserNik, err := m.GetUserNik(nik)
	if err != nil {
		return err
	}

	if err := u.AddressRepo.DeleteAddressUser(getUserNik); err != nil {
		return err
	}

	if err := u.UserRepo.DeleteUser(getUserNik); err != nil {
		return err
	}

	return nil
}

func (u *userService) NearbyHealthFacilities(payloads payload.NearbyHealth, nik string) (response.UserNearbyHealth, error) {
	var result response.UserNearbyHealth
	var tempData []response.HealthResponse
	var tempRes []response.HealthResponse
	getUserNik, err := m.GetUserNik(nik)
	if err != nil {
		return result, err
	}

	userProfile, err := u.UserRepo.GetUserDataByNik(getUserNik)
	if err != nil {
		return result, err
	}

	allHealthFacilities, err := u.HealthRepo.GetAllHealthFacilities()
	tempData = make([]response.HealthResponse, len(allHealthFacilities))
	if err != nil {
		return result, err
	}

	for i, val := range allHealthFacilities {
		newRanges := util.FindRange(payloads.Latitude, payloads.Longitude, val.Address.Latitude, val.Address.Longitude)
		if newRanges < 10 {
			tempData[i] = response.HealthResponse{
				ID:       val.ID,
				Email:    val.Email,
				PhoneNum: val.PhoneNum,
				Name:     val.Name,
				Image:    val.Image,
				Ranges:   newRanges,
				Address:  *val.Address,
				Vaccine:  val.Vaccine,
			}
		}
	}

	sort.Slice(tempData, func(i, j int) bool {
		return tempData[i].Ranges < tempData[j].Ranges
	})

	length := 0
	for _, val := range tempData {
		if val.ID == "" {
			length += 1
		}
	}
	tempRes = append(tempRes, tempData[length:]...)

	ageUser, err := u.UserRepo.GetAgeUser(userProfile)
	if err != nil {
		return result, err
	}

	result = response.UserNearbyHealth{
		User: response.UserProfile{
			NIK:          userProfile.NIK,
			Email:        userProfile.Email,
			Fullname:     userProfile.Fullname,
			PhoneNum:     userProfile.PhoneNum,
			Gender:       userProfile.Gender,
			VaccineCount: userProfile.VaccineCount,
			Age:          ageUser.Age,
			Address:      userProfile.Address,
		},
		HealthFacilities: tempRes,
	}

	return result, nil
}
