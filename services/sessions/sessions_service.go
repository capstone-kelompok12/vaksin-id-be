package services

import (
	"errors"
	"time"
	"vaksin-id-be/dto/payload"
	"vaksin-id-be/dto/response"
	m "vaksin-id-be/middleware"
	"vaksin-id-be/model"
	mysqls "vaksin-id-be/repository/mysql/sessions"
	mysqlv "vaksin-id-be/repository/mysql/vaccines"

	"github.com/google/uuid"
)

type SessionsService interface {
	CreateSessions(payloads payload.SessionsPayload, auth string) (response.SessionsResponse, error)
	GetAllSessions() ([]response.SessionsResponse, error)
	GetSessionsAdminById(auth, id string) (response.SessionsResponse, error)
	GetSessionByAdmin(auth string) ([]response.SessionsResponse, error)
	GetSessionActive() (response.IsCloseFalse, error)
	UpdateSession(payloads payload.SessionsUpdate, id string) (response.SessionsUpdate, error)
	IsCloseSession(payloads payload.SessionsIsClose, auth, id string) (response.SessionsResponse, error)
	DeleteSession(id string) error
}

type sessionService struct {
	SessionsRepo mysqls.SessionsRepository
	VaccineRepo  mysqlv.VaccinesRepository
}

func NewSessionsService(sessionRepo mysqls.SessionsRepository, vaccineRepo mysqlv.VaccinesRepository) *sessionService {
	return &sessionService{
		SessionsRepo: sessionRepo,
		VaccineRepo:  vaccineRepo,
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

	newData, err := s.SessionsRepo.GetSessionById(id)
	if err != nil {
		return sessionResponse, err
	}

	sessionResponse = response.SessionsResponse{
		ID:           newData.ID,
		IdVaccine:    newData.IdVaccine,
		SessionName:  newData.SessionName,
		Capacity:     newData.Capacity,
		IsClose:      newData.IsClose,
		Dose:         newData.Dose,
		Date:         newData.Date,
		StartSession: newData.StartSession,
		EndSession:   newData.EndSession,
		CreatedAt:    newData.CreatedAt,
		UpdatedAt:    newData.UpdatedAt,
		Vaccine:      newData.Vaccine,
		Booking:      newData.Booking,
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
		sessionsResponse[i] = response.SessionsResponse{
			ID:           v.ID,
			IdVaccine:    v.IdVaccine,
			SessionName:  v.SessionName,
			Capacity:     v.Capacity,
			Dose:         v.Dose,
			Date:         v.Date,
			IsClose:      v.IsClose,
			StartSession: v.StartSession,
			EndSession:   v.EndSession,
			CreatedAt:    v.CreatedAt,
			UpdatedAt:    v.UpdatedAt,
			Vaccine:      v.Vaccine,
			// Booking:      v.Booking,
		}
	}

	return sessionsResponse, nil
}

func (s *sessionService) GetSessionsAdminById(auth, id string) (response.SessionsResponse, error) {
	var responseSession response.SessionsResponse

	getIdHealthFacilities, err := m.GetIdHealthFacilities(auth)
	if err != nil {
		return responseSession, err
	}

	getData, err := s.SessionsRepo.GetSessionAdminById(getIdHealthFacilities, id)
	if err != nil {
		return responseSession, err
	}

	responseSession = response.SessionsResponse{
		ID:           getData.ID,
		IdVaccine:    getData.IdVaccine,
		SessionName:  getData.SessionName,
		Capacity:     getData.Capacity,
		Dose:         getData.Dose,
		Date:         getData.Date,
		IsClose:      getData.IsClose,
		StartSession: getData.StartSession,
		EndSession:   getData.EndSession,
		CreatedAt:    getData.CreatedAt,
		UpdatedAt:    getData.UpdatedAt,
		Vaccine:      getData.Vaccine,
		// Booking:      getData.Booking,
	}

	return responseSession, nil
}

func (s *sessionService) GetSessionByAdmin(auth string) ([]response.SessionsResponse, error) {
	var responseSession []response.SessionsResponse

	getIdHealthFacilities, err := m.GetIdHealthFacilities(auth)
	if err != nil {
		return responseSession, err
	}

	getData, err := s.SessionsRepo.GetSessionsByAdmin(getIdHealthFacilities)
	if err != nil {
		return responseSession, err
	}

	responseSession = make([]response.SessionsResponse, len(getData))

	for i, val := range getData {
		responseSession[i] = response.SessionsResponse{
			ID:           val.ID,
			IdVaccine:    val.IdVaccine,
			SessionName:  val.SessionName,
			Capacity:     val.Capacity,
			Dose:         val.Dose,
			Date:         val.Date,
			IsClose:      val.IsClose,
			StartSession: val.StartSession,
			EndSession:   val.EndSession,
			CreatedAt:    val.CreatedAt,
			UpdatedAt:    val.UpdatedAt,
			Vaccine:      val.Vaccine,
			Booking:      val.Booking,
		}
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

func (s *sessionService) IsCloseSession(payloads payload.SessionsIsClose, auth, id string) (response.SessionsResponse, error) {
	var respData response.SessionsResponse

	dataUpdate := model.Sessions{
		IsClose: payloads.IsClose,
	}

	err := s.SessionsRepo.UpdateSession(dataUpdate, id)
	if err != nil {
		return respData, err
	}

	data, err := s.GetSessionsAdminById(auth, id)
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
