package services

import (
	"vaksin-id-be/dto/payload"
	"vaksin-id-be/dto/response"
	m "vaksin-id-be/middleware"
	"vaksin-id-be/model"
	mysqls "vaksin-id-be/repository/mysql/sessions"

	"github.com/google/uuid"
)

type SessionsService interface {
	CreateSessions(payloads payload.SessionsPayload, auth string) (model.Sessions, error)
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
}

func NewSessionsService(sessionRepo mysqls.SessionsRepository) *sessionService {
	return &sessionService{
		SessionsRepo: sessionRepo,
	}
}

func (s *sessionService) CreateSessions(payloads payload.SessionsPayload, auth string) (model.Sessions, error) {
	var sessionModel model.Sessions

	getIdHealthFacilities, err := m.GetIdHealthFacilities(auth)
	if err != nil {
		return sessionModel, err
	}

	defaultStatus := false

	id := uuid.NewString()

	sessionModel = model.Sessions{
		ID:                 id,
		IdHealthFacilities: getIdHealthFacilities,
		SessionName:        payloads.SessionName,
		Capacity:           payloads.Capacity,
		Dose:               payloads.Dose,
		IsClose:            defaultStatus,
		StartSession:       payloads.StartSession,
		EndSession:         payloads.EndSession,
	}

	err = s.SessionsRepo.CreateSession(sessionModel)
	if err != nil {
		return sessionModel, err
	}

	return sessionModel, nil
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
			SessionName:  v.SessionName,
			Capacity:     v.Capacity,
			Dose:         v.Dose,
			IsClose:      v.IsClose,
			StartSession: v.StartSession,
			EndSession:   v.EndSession,
			CreatedAt:    v.CreatedAt,
			UpdatedAt:    v.UpdatedAt,
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
		SessionName:  getData.SessionName,
		Capacity:     getData.Capacity,
		Dose:         getData.Dose,
		IsClose:      getData.IsClose,
		StartSession: getData.StartSession,
		EndSession:   getData.EndSession,
		CreatedAt:    getData.CreatedAt,
		UpdatedAt:    getData.UpdatedAt,
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
			SessionName:  val.SessionName,
			Capacity:     val.Capacity,
			Dose:         val.Dose,
			IsClose:      val.IsClose,
			StartSession: val.StartSession,
			EndSession:   val.EndSession,
			CreatedAt:    val.CreatedAt,
			UpdatedAt:    val.UpdatedAt,
			Booking:      val.Booking,
		}
	}

	return responseSession, nil
}

func (s *sessionService) UpdateSession(payloads payload.SessionsUpdate, id string) (response.SessionsUpdate, error) {
	var respData response.SessionsUpdate

	sessionData := model.Sessions{
		SessionName:  payloads.SessionName,
		Capacity:     payloads.Capacity,
		StartSession: payloads.StartSession,
		EndSession:   payloads.EndSession,
	}

	if err := s.SessionsRepo.UpdateSession(sessionData, id); err != nil {
		return respData, err
	}

	respData = response.SessionsUpdate{
		ID:           id,
		SessionName:  payloads.SessionName,
		Capacity:     payloads.Capacity,
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
