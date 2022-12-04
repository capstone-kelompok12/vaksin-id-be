package services

import (
	"vaksin-id-be/dto/payload"
	"vaksin-id-be/dto/response"
	"vaksin-id-be/model"
	mysqls "vaksin-id-be/repository/mysql/sessions"

	"github.com/google/uuid"
)

type SessionsService interface {
	CreateSessions(payloads payload.SessionsPayload) error
	GetAllSessions() ([]response.SessionsResponse, error)
	GetSession(id string) (response.SessionsResponse, error)
	UpdateSession(payloads payload.SessionsUpdate, id string) error
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

func (s *sessionService) CreateSessions(payloads payload.SessionsPayload) error {
	id := uuid.NewString()

	sessionModel := model.Sessions{
		ID:                 id,
		IdHealthFacilities: payloads.IdHealthFacilities,
		SessionName:        payloads.SessionName,
		Capacity:           payloads.Capacity,
		StartSession:       payloads.StartSession,
		EndSession:         payloads.EndSession,
	}

	err := s.SessionsRepo.CreateSession(sessionModel)
	if err != nil {
		return err
	}

	return nil
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
			ID:            v.ID,
			SessionName:   v.SessionName,
			Capacity:      v.Capacity,
			SessionStatus: v.SessionStatus,
			StartSession:  v.StartSession,
			EndSession:    v.EndSession,
			CreatedAt:     v.CreatedAt,
			UpdatedAt:     v.UpdatedAt,
		}
	}

	return sessionsResponse, nil
}

func (s *sessionService) GetSession(id string) (response.SessionsResponse, error) {
	var responseSession response.SessionsResponse

	getData, err := s.SessionsRepo.GetSession(id)
	if err != nil {
		return responseSession, err
	}

	responseSession = response.SessionsResponse{
		ID:            getData.ID,
		SessionName:   getData.SessionName,
		Capacity:      getData.Capacity,
		SessionStatus: getData.SessionStatus,
		StartSession:  getData.StartSession,
		EndSession:    getData.EndSession,
		CreatedAt:     getData.CreatedAt,
		UpdatedAt:     getData.UpdatedAt,
	}

	return responseSession, nil
}

func (s *sessionService) UpdateSession(payloads payload.SessionsUpdate, id string) error {
	sessionData := model.Sessions{
		ID:            payloads.ID,
		SessionName:   payloads.SessionName,
		Capacity:      payloads.Capacity,
		SessionStatus: payloads.SessionStatus,
		StartSession:  payloads.StartSession,
		EndSession:    payloads.EndSession,
	}

	if err := s.SessionsRepo.UpdateSession(sessionData, id); err != nil {
		return err
	}
	return nil
}

func (s *sessionService) DeleteSession(id string) error {
	if err := s.SessionsRepo.DeleteSession(id); err != nil {
		return err
	}

	return nil
}
