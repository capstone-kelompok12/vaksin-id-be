package services

import (
	"errors"
	"fmt"
	"time"
	"vaksin-id-be/dto/payload"
	"vaksin-id-be/dto/response"
	"vaksin-id-be/model"
	mysqlb "vaksin-id-be/repository/mysql/bookings"
	mysqls "vaksin-id-be/repository/mysql/sessions"
	mysqlu "vaksin-id-be/repository/mysql/users"
	mysqlv "vaksin-id-be/repository/mysql/vaccines"

	"github.com/google/uuid"
)

type SessionsService interface {
	CreateSessions(payloads payload.SessionsPayload, auth string) (response.SessionsResponse, error)
	GetAllSessions() ([]response.SessionsResponse, error)
	GetAllSessionsByAdmin(auth string) ([]response.SessionsResponse, error)
	GetSessionsById(id string) (response.SessionsResponse, error)
	GetSessionActive() (response.IsCloseFalse, error)
	GetAllFinishedSessionCount() (response.SessionFinished, error)
	UpdateSession(payloads payload.SessionsUpdate, id string) (response.SessionsUpdate, error)
	IsCloseSession(payloads payload.SessionsIsClose, id string) (response.SessionsResponse, error)
	DeleteSession(id string) error
}

type sessionService struct {
	SessionsRepo mysqls.SessionsRepository
	VaccineRepo  mysqlv.VaccinesRepository
	BookingRepo  mysqlb.BookingRepository
	UserRepo     mysqlu.UserRepository
}

func NewSessionsService(sessionRepo mysqls.SessionsRepository, vaccineRepo mysqlv.VaccinesRepository, bookingRepo mysqlb.BookingRepository, userRepo mysqlu.UserRepository) *sessionService {
	return &sessionService{
		SessionsRepo: sessionRepo,
		VaccineRepo:  vaccineRepo,
		BookingRepo:  bookingRepo,
		UserRepo:     userRepo,
	}
}

func (s *sessionService) CreateSessions(payloads payload.SessionsPayload, auth string) (response.SessionsResponse, error) {
	var sessionModel model.Sessions
	var sessionResponse response.SessionsResponse

	defaultStatus := false

	id := uuid.NewString()

	dateSession, err := time.Parse("2006-01-02", payloads.Date)
	if err != nil {
		return sessionResponse, err
	}

	getDoseFromVaccine, err := s.VaccineRepo.GetVaccinesById(payloads.IdVaccine)
	if err != nil {
		return sessionResponse, err
	}

	sessionModel = model.Sessions{
		ID:           id,
		IdVaccine:    payloads.IdVaccine,
		SessionName:  payloads.SessionName,
		Capacity:     payloads.Capacity,
		Dose:         getDoseFromVaccine.Dose,
		Date:         dateSession,
		IsClose:      defaultStatus,
		StartSession: payloads.StartSession,
		EndSession:   payloads.EndSession,
	}

	_, err = s.SessionsRepo.CreateSession(sessionModel)
	if err != nil {
		return sessionResponse, err
	}

	totalCap, err := s.SessionsRepo.GetSumOfCapacity(getDoseFromVaccine.ID)
	if err != nil {
		return sessionResponse, err
	}

	if totalCap.TotalCapacity > getDoseFromVaccine.Stock {
		err := s.SessionsRepo.DeleteSession(id)
		if err != nil {
			return sessionResponse, err
		}
		return sessionResponse, errors.New("total capacity all sessions exceed of stock vaccine")
	}

	getBooking, err := s.BookingRepo.GetAllBookingBySession(sessionModel.ID)
	if err != nil {
		return sessionResponse, err
	}
	newData, err := s.SessionsRepo.GetSessionById(id)
	if err != nil {
		return sessionResponse, err
	}

	getbackCap := newData.Capacity - len(getBooking)

	sessionResponse = response.SessionsResponse{
		ID:           newData.ID,
		IdVaccine:    newData.IdVaccine,
		SessionName:  newData.SessionName,
		Capacity:     newData.Capacity,
		CapacityLeft: getbackCap,
		IsClose:      newData.IsClose,
		Dose:         newData.Dose,
		Date:         newData.Date,
		StartSession: newData.StartSession,
		EndSession:   newData.EndSession,
		CreatedAt:    newData.CreatedAt,
		UpdatedAt:    newData.UpdatedAt,
		Vaccine:      newData.Vaccine,
	}

	return sessionResponse, nil
}

func (s *sessionService) GetAllSessions() ([]response.SessionsResponse, error) {
	var sessionsResponse []response.SessionsResponse

	getSessions, err := s.SessionsRepo.GetAllSessions()

	if err != nil {
		return sessionsResponse, err
	}

	sessionsResponse = make([]response.SessionsResponse, len(getSessions))

	for i, v := range getSessions {

		getBooking, err := s.BookingRepo.GetAllBookingBySession(v.ID)
		if err != nil {
			return sessionsResponse, err
		}
		fmt.Print(len(getBooking))

		dataBooking := make([]response.BookingInSession, len(getBooking))

		for idx, value := range getBooking {

			dataBooking[idx] = response.BookingInSession{
				ID:        value.ID,
				IdSession: value.IdSession,
				Queue:     value.Queue,
				Status:    value.Status,
				CreatedAt: value.CreatedAt,
				UpdatedAt: value.UpdatedAt,
			}
		}

		getSessionById, err := s.SessionsRepo.GetSessionById(v.ID)
		if err != nil {
			return sessionsResponse, err
		}
		getbackCap := getSessionById.Capacity - len(getBooking)

		sessionsResponse[i] = response.SessionsResponse{
			ID:           v.ID,
			IdVaccine:    v.IdVaccine,
			SessionName:  v.SessionName,
			Capacity:     v.Capacity,
			CapacityLeft: getbackCap,
			Dose:         v.Dose,
			Date:         v.Date,
			IsClose:      v.IsClose,
			StartSession: v.StartSession,
			EndSession:   v.EndSession,
			CreatedAt:    v.CreatedAt,
			UpdatedAt:    v.UpdatedAt,
			Vaccine:      v.Vaccine,
			Booking:      dataBooking,
		}
	}

	return sessionsResponse, nil
}

func (s *sessionService) GetAllSessionsByAdmin(auth string) ([]response.SessionsResponse, error) {
	var sessionsResponse []response.SessionsResponse

	getData, err := s.SessionsRepo.GetSessionById(id)
	if err != nil {
		return responseSession, err
	}

	countBooking, err := s.BookingRepo.GetAllBookingBySession(id)
	if err != nil {
		return sessionsResponse, err
	}

	getSessions, err := s.SessionsRepo.GetSessionsByAdmin(getIdHealthFacilities)
	if err != nil {
		return sessionsResponse, err
	}

	sessionsResponse = make([]response.SessionsResponse, len(getSessions))

	for i, v := range getSessions {

		getBooking, err := s.BookingRepo.GetAllBookingBySession(v.ID)
		if err != nil {
			return sessionsResponse, err
		}

		dataBooking := make([]response.BookingInSession, len(getBooking))

		for idx, value := range getBooking {
			// getUserData, err := s.UserRepo.GetDataByIdBooking(value.ID)
			// if err != nil {
			// 	return sessionsResponse, err
			// }

			// fmt.Println(&getUserData)

			// getDataUser, err := s.UserRepo.GetUserDataByNikNoAddress(getUserData.NIK)
			// if err != nil {
			// 	return sessionsResponse, err
			// }
			// fmt.Println(&getUserData)

			// ageUser, err := s.UserRepo.GetAgeUser(getDataUser)
			// if err != nil {
			// 	return sessionsResponse, err
			// }

			// fmt.Println(ageUser.Age)

			// // getHistoryUser

			// userData := response.SessionUserCustom{
			// 	NIK:          getDataUser.NIK,
			// 	Email:        getDataUser.Email,
			// 	Fullname:     getDataUser.Fullname,
			// 	PhoneNum:     getDataUser.PhoneNum,
			// 	Gender:       getDataUser.Gender,
			// 	VaccineCount: getDataUser.VaccineCount,
			// 	BirthDate:    getDataUser.BirthDate,
			// 	Age:          ageUser.Age,
			// 	Address:      getDataUser.Address,
			// }

			dataBooking[idx] = response.BookingInSession{
				ID:        value.ID,
				IdSession: value.IdSession,
				Queue:     value.Queue,
				Status:    value.Status,
				CreatedAt: value.CreatedAt,
				UpdatedAt: value.UpdatedAt,
				User:      &value.User,
			}

		}

		getSessionById, err := s.SessionsRepo.GetSessionById(v.ID)
		if err != nil {
			return sessionsResponse, err
		}
		fmt.Println(getSessionById.Capacity)

		getbackCap := getSessionById.Capacity - len(getBooking)

		sessionsResponse[i] = response.SessionsResponse{
			ID:           v.ID,
			IdVaccine:    v.IdVaccine,
			SessionName:  v.SessionName,
			Capacity:     v.Capacity,
			CapacityLeft: getbackCap,
			Dose:         v.Dose,
			Date:         v.Date,
			IsClose:      v.IsClose,
			StartSession: v.StartSession,
			EndSession:   v.EndSession,
			CreatedAt:    v.CreatedAt,
			UpdatedAt:    v.UpdatedAt,
			Vaccine:      v.Vaccine,
			Booking:      dataBooking,
		}
	}

	return sessionsResponse, nil
}

func (s *sessionService) GetSessionsById(id string) (response.SessionsResponse, error) {
	var responseSession response.SessionsResponse

	getData, err := s.SessionsRepo.GetSessionById(id)
	if err != nil {
		return responseSession, err
	}

	countBooking, err := s.BookingRepo.GetAllBookingBySession(id)
	if err != nil {
		return responseSession, err
	}
	getbackCap := getSessionById.Capacity - len(countBooking)

	dataBooking := make([]response.BookingInSession, len(countBooking))

	for i, val := range countBooking {
		dataBooking[i] = response.BookingInSession{
			ID:        val.ID,
			IdSession: val.IdSession,
			Queue:     val.Queue,
			Status:    val.Status,
			CreatedAt: val.CreatedAt,
			UpdatedAt: val.UpdatedAt,
		}
	}

	getSessionById, err := s.SessionsRepo.GetSessionById(id)
	if err != nil {
		return responseSession, err
	}
	getbackCap := getSessionById.Capacity - len(countBooking)

	dataBooking := make([]response.BookingInSession, len(countBooking))

	for i, val := range countBooking {
		dataBooking[i] = response.BookingInSession{
			ID:        val.ID,
			IdSession: val.IdSession,
			Queue:     val.Queue,
			Status:    val.Status,
			CreatedAt: val.CreatedAt,
			UpdatedAt: val.UpdatedAt,
		}
	}

	responseSession = response.SessionsResponse{
		ID:           getData.ID,
		IdVaccine:    getData.IdVaccine,
		SessionName:  getData.SessionName,
		Capacity:     getData.Capacity,
		CapacityLeft: getbackCap,
		Dose:         getData.Dose,
		Date:         getData.Date,
		IsClose:      getData.IsClose,
		StartSession: getData.StartSession,
		EndSession:   getData.EndSession,
		CreatedAt:    getData.CreatedAt,
		UpdatedAt:    getData.UpdatedAt,
		Vaccine:      getData.Vaccine,
		Booking:      dataBooking,
	}

	return responseSession, nil
}

func (s *sessionService) UpdateSession(payloads payload.SessionsUpdate, id string) (response.SessionsUpdate, error) {
	var respData response.SessionsUpdate

	dateSession, err := time.Parse("2006-01-02", payloads.Date)
	if err != nil {
		return respData, err
	}

	getDoseFromVaccine, err := s.VaccineRepo.GetVaccinesById(payloads.IdVaccine)
	if err != nil {
		return respData, err
	}

	sessionData := model.Sessions{
		SessionName:  payloads.SessionName,
		Capacity:     payloads.Capacity,
		Dose:         getDoseFromVaccine.Dose,
		Date:         dateSession,
		StartSession: payloads.StartSession,
		EndSession:   payloads.EndSession,
	}

	if err := s.SessionsRepo.UpdateSession(sessionData, id); err != nil {
		return respData, err
	}

	respData = response.SessionsUpdate{
		ID:           id,
		IdVaccine:    payloads.IdVaccine,
		SessionName:  payloads.SessionName,
		Capacity:     payloads.Capacity,
		Dose:         getDoseFromVaccine.Dose,
		Date:         dateSession,
		StartSession: payloads.StartSession,
		EndSession:   payloads.EndSession,
	}

	return respData, nil
}

func (s *sessionService) IsCloseSession(payloads payload.SessionsIsClose, id string) (response.SessionsResponse, error) {
	var respData response.SessionsResponse

	dataUpdate := model.Sessions{
		IsClose: payloads.IsClose,
	}

	err := s.SessionsRepo.UpdateSession(dataUpdate, id)
	if err != nil {
		return respData, err
	}

	data, err := s.GetSessionsById(id)
	if err != nil {
		return respData, err
	}

	return data, nil
}

func (s *sessionService) DeleteSession(id string) error {
	if err := s.SessionsRepo.DeleteSession(id); err != nil {
		return err
	}

	return nil
}

func (s *sessionService) GetSessionActive() (response.IsCloseFalse, error) {
	var sessionResponse response.IsCloseFalse

	getSession, err := s.SessionsRepo.IsCloseFalse()

	if err != nil {
		return sessionResponse, err
	}

	sessionResponse = response.IsCloseFalse{Active: getSession.Active}

	return sessionResponse, nil
}

func (s *sessionService) GetAllFinishedSessionCount() (response.SessionFinished, error) {
	var sessionResponse response.SessionFinished

	data, err := s.SessionsRepo.GetAllFinishedSessionCount()
	if err != nil {
		return sessionResponse, err
	}

	sessionResponse = data

	return sessionResponse, err
}
