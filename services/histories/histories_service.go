package services

import (
	"vaksin-id-be/dto/payload"
	"vaksin-id-be/dto/response"
	"vaksin-id-be/model"
	mysqlhs "vaksin-id-be/repository/mysql/histories"

	"github.com/google/uuid"
)

type HistoriesRepository interface {
	CreateHistory(payloads payload.HistoriesPayload) error
	GetAllHistory() ([]response.HistoryResponse, error)
	GetHistoryById(id string) (response.HistoryResponse, error)
	UpdateHistory(payloads payload.UpdateAccHistory, id string) (response.HistoryResponse, error)
}

type historiesService struct {
	HistoriesRepo mysqlhs.HistoriesRepository
}

func NewHistoriesService(historiesRepo mysqlhs.HistoriesRepository) *historiesService {
	return &historiesService{
		HistoriesRepo: historiesRepo,
	}
}

func (h *historiesService) CreateHistory(payloads payload.HistoriesPayload) error {
	var historyModel model.VaccineHistories

	id := uuid.NewString()

	historyModel = model.VaccineHistories{
		ID:         id,
		IdBooking:  payloads.IdBooking,
		NikUser:    payloads.NikUser,
		IdSameBook: payloads.IdSameBook,
		Status:     &payloads.Status,
	}

	err := h.HistoriesRepo.CreateHistory(historyModel)

	if err != nil {
		return err
	}

	return nil
}

func (h *historiesService) GetAllHistory() ([]response.HistoryResponse, error) {
	var historyResponse []response.HistoryResponse

	getHistory, err := h.HistoriesRepo.GetAllHistory()

	if err != nil {
		return historyResponse, err
	}

	historyResponse = make([]response.HistoryResponse, len(getHistory))

	for i, v := range getHistory {
		historyResponse[i] = response.HistoryResponse{
			ID:         v.ID,
			IdBooking:  v.IdBooking,
			NikUser:    v.NikUser,
			IdSameBook: v.IdSameBook,
			Status:     v.Status,
			CreatedAt:  v.CreatedAt,
			UpdatedAt:  v.UpdatedAt,
			User:       *v.User,
		}
	}

	return historyResponse, nil
}

func (h *historiesService) GetHistoryById(id string) (response.HistoryResponse, error) {
	var responseHistory response.HistoryResponse

	getData, err := h.HistoriesRepo.GetHistoryById(id)
	if err != nil {
		return responseHistory, err
	}

	responseHistory = response.HistoryResponse{
		ID:         getData.ID,
		IdBooking:  getData.IdBooking,
		NikUser:    getData.NikUser,
		IdSameBook: getData.IdSameBook,
		Status:     getData.Status,
		CreatedAt:  getData.CreatedAt,
		UpdatedAt:  getData.UpdatedAt,
		User:       *getData.User,
	}

	return responseHistory, nil
}

func (h *historiesService) UpdateHistory(id string, payloads payload.UpdateAccHistory) (response.HistoryResponse, error) {
	var responseHistory response.HistoryResponse

	historyData := model.VaccineHistories{
		ID:      payloads.ID,
		NikUser: payloads.NikUser,
		Status:  &payloads.Status,
	}

	if _, err := h.HistoriesRepo.UpdateHistory(historyData, id); err != nil {
		return responseHistory, err
	}
	return responseHistory, nil
}
