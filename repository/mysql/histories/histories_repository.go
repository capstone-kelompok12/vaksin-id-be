package mysql

import (
	"vaksin-id-be/model"

	"gorm.io/gorm"
)

type HistoriesRepository interface {
	CreateHistory(data model.VaccineHistories) (model.VaccineHistories, error)
	GetAllHistory() ([]model.VaccineHistories, error)
	GetHistoryById(id string) (model.VaccineHistories, error)
	UpdateHistory(data model.VaccineHistories, id string) (model.VaccineHistories, error)
}

type historiesRepository struct {
	db *gorm.DB
}

func NewHistoryRepository(db *gorm.DB) *historiesRepository {
	return &historiesRepository{
		db: db,
	}
}

func (h *historiesRepository) CreateHistory(data model.VaccineHistories) (model.VaccineHistories, error) {
	var datas model.VaccineHistories
	if err := h.db.Create(&data).Error; err != nil {
		return datas, err
	}
	return data, nil
}
func (h *historiesRepository) GetAllHistory() ([]model.VaccineHistories, error) {
	var history []model.VaccineHistories
	if err := h.db.Model(&model.VaccineHistories{}).Find(&history).Error; err != nil {
		return history, err
	}
	return history, nil
}

func (h *historiesRepository) GetHistoryById(id string) (model.VaccineHistories, error) {
	var history model.VaccineHistories
	if err := h.db.Where("id = ?", id).First(&history).Error; err != nil {
		return history, err
	}
	return history, nil
}

func (h *historiesRepository) UpdateHistory(data model.VaccineHistories, id string) (model.VaccineHistories, error) {
	var history model.VaccineHistories
	if err := h.db.Model(&model.VaccineHistories{}).Where("id = ?", id).Updates(&data).Error; err != nil {
		return history, err
	}
	return history, nil
}
