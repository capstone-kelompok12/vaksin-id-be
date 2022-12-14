package services

import (
	"errors"
	"fmt"
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
	mysqlv "vaksin-id-be/repository/mysql/vaccines"
	"vaksin-id-be/util"

	"github.com/google/uuid"
)

type UserService interface {
	RegisterUser(payloads payload.RegisterUser) error
	LoginUser(payloads payload.Login) (response.Login, error)
	GetUserDataByNik(nik string) (response.UserProfile, error)
	GetUserDataByNikNoAddress(nik string) (response.UserProfile, error)
	GetUserHistory(nik string) (response.UserHistory, error)
	GetUserRegisteredDashboard() (response.RegisterStatistic, error)
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
	VaccineRepo mysqlv.VaccinesRepository
}

func NewUserService(
	userRepo mysqlu.UserRepository,
	addressRepo mysqla.AddressesRepository,
	healthRepo mysqlh.HealthFacilitiesRepository,
	historyRepo mysqlhs.HistoriesRepository,
	bookingRepo mysqlb.BookingRepository,
	sessionRepo mysqls.SessionsRepository,
	vaccineRepo mysqlv.VaccinesRepository,
) *userService {
	return &userService{
		UserRepo:    userRepo,
		AddressRepo: addressRepo,
		HealthRepo:  healthRepo,
		HistoryRepo: historyRepo,
		BookingRepo: bookingRepo,
		SessionRepo: sessionRepo,
		VaccineRepo: vaccineRepo,
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

	getData, err := u.UserRepo.GetUserHistoryByNik(nik)
	if err != nil {
		return historyUser, err
	}

	getDataUser, err := u.UserRepo.GetUserDataByNikNoAddress(nik)
	if err != nil {
		return historyUser, err
	}

	ageUser, err := u.UserRepo.GetAgeUser(getDataUser)
	if err != nil {
		return historyUser, err
	}

	countHistoryUser, err := u.HistoryRepo.GetHistoriesByNIK(nik)
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

func (u *userService) GetUserRegisteredDashboard() (response.RegisterStatistic, error) {
	var responseDash response.RegisterStatistic
	RegisteredData := make([]response.DashboardForm, 3)

	var Name string
	var FirstDose int
	var SecondDose int
	var ThirdDose int
	var Kosong int

	getData, err := u.UserRepo.GetAllUser()
	if err != nil {
		return responseDash, err
	}

	for _, val := range getData {
		ageUser, err := u.UserRepo.GetAgeUser(val)
		if err != nil {
			return responseDash, err
		}
		RegisteredData[0].Name = "12 - 17 Tahun"
		RegisteredData[1].Name = "18 - 59 Tahun"
		RegisteredData[2].Name = "60 Tahun Ke atas"
		if ageUser.Age >= 12 && ageUser.Age <= 17 {
			if val.VaccineCount == 1 {
				FirstDose += 1
			} else if val.VaccineCount == 2 {
				SecondDose += 1
			} else if val.VaccineCount == 3 {
				ThirdDose += 1
			} else {
				Kosong = 0
				fmt.Print(Kosong)
			}
			Name = "12 - 17 Tahun"
			RegisteredData[0] = response.DashboardForm{
				Name:      Name,
				DoseOne:   FirstDose,
				DoseTwo:   SecondDose,
				DoseThree: ThirdDose,
			}
		} else if ageUser.Age >= 18 && ageUser.Age <= 59 {
			if val.VaccineCount == 1 {
				FirstDose += 1
			} else if val.VaccineCount == 2 {
				SecondDose += 1
			} else if val.VaccineCount == 3 {
				ThirdDose += 1
			} else {
				Kosong = 0
				fmt.Print(Kosong)
			}
			Name = "18 - 59 Tahun"
			RegisteredData[1] = response.DashboardForm{
				Name:      Name,
				DoseOne:   FirstDose,
				DoseTwo:   SecondDose,
				DoseThree: ThirdDose,
			}
		} else {
			if val.VaccineCount == 1 {
				FirstDose += 1
			} else if val.VaccineCount == 2 {
				SecondDose += 1
			} else if val.VaccineCount == 3 {
				ThirdDose += 1
			} else {
				Kosong = 0
				fmt.Print(Kosong)
			}
			Name = "60 Tahun Ke atas"
			RegisteredData[2] = response.DashboardForm{
				Name:      Name,
				DoseOne:   FirstDose,
				DoseTwo:   SecondDose,
				DoseThree: ThirdDose,
			}
		}
	}
	responseData := response.RegisterStatistic{
		RegisteredStat: RegisteredData,
	}

	return responseData, nil
}

func (u *userService) UpdateUserProfile(payloads payload.UpdateUser, nik string) (response.UpdateUser, error) {
	var dataResp response.UpdateUser

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

	dataUser := model.Users{
		Email:     payloads.Email,
		Password:  hashPass,
		Fullname:  payloads.Fullname,
		PhoneNum:  payloads.PhoneNum,
		Gender:    payloads.Gender,
		BirthDate: dateBirth,
	}

	if err := u.UserRepo.UpdateUserProfile(dataUser, nik); err != nil {
		return dataResp, err
	}

	data, err := u.GetUserDataByNikNoAddress(nik)
	if err != nil {
		return dataResp, err
	}

	dataResp = response.UpdateUser{
		Fullname:  data.Fullname,
		NikUser:   data.NIK,
		Email:     data.Email,
		PhoneNum:  data.PhoneNum,
		Gender:    data.Gender,
		BirthDate: dateBirth,
	}

	return dataResp, nil
}

func (u *userService) DeleteUserProfile(nik string) error {

	if err := u.AddressRepo.DeleteAddressUser(nik); err != nil {
		return err
	}

	if err := u.UserRepo.DeleteUser(nik); err != nil {
		return err
	}

	return nil
}

func (u *userService) NearbyHealthFacilities(payloads payload.NearbyHealth, nik string) (response.UserNearbyHealth, error) {
	var result response.UserNearbyHealth
	var tempData []response.HealthResponse
	var tempRes []response.HealthResponse

	userProfile, err := u.UserRepo.GetUserDataByNik(nik)
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
