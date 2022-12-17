package mysql

import (
	"vaksin-id-be/dto/response"
	"vaksin-id-be/model"

	"gorm.io/gorm"
	"gorm.io/gorm/clause"
)

type HistoriesRepository interface {
	CreateHistory(data model.VaccineHistories) error
	GetAllHistory() ([]model.VaccineHistories, error)
	GetHistoryById(id string) (model.VaccineHistories, error)
	GetHistoryByIdNoUserHistory(id string) (model.VaccineHistories, error)
	GetHistoriesById(id string) ([]model.VaccineHistories, error)
	GetHistoryByIdBooking(id string) ([]model.VaccineHistories, error)
	GetHistoriesByIdBooking(id string) ([]model.VaccineHistories, error)
	GetHistoryByIdSameBook(id string) ([]model.VaccineHistories, error)
	GetHistoryByNIK(id, nik string) (model.VaccineHistories, error)
	GetHistoriesByNIK(nik string) ([]model.VaccineHistories, error)
	UpdateHistoryByNik(data model.VaccineHistories, nik, id string) (model.VaccineHistories, error)
	UpdateHistory(data model.VaccineHistories, id string) (model.VaccineHistories, error)
	CheckVaccineCount(nik string) ([]model.VaccineHistories, error)
	GetTotalUserVaccinated() (response.VaccinatedUser, error)
}

type historiesRepository struct {
	db *gorm.DB
}

func NewHistoryRepository(db *gorm.DB) *historiesRepository {
	return &historiesRepository{
		db: db,
	}
}

func (h *historiesRepository) CreateHistory(data model.VaccineHistories) error {
	if err := h.db.Create(&data).Error; err != nil {
		return err
	}
	return nil
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
	if err := h.db.Preload(clause.Associations).Preload("Booking."+clause.Associations).Preload("User."+clause.Associations).Where("id = ?", id).First(&history).Error; err != nil {
		return history, err
	}
	return history, nil
}

func (h *historiesRepository) GetHistoryByIdNoUserHistory(id string) (model.VaccineHistories, error) {
	var history model.VaccineHistories
	if err := h.db.Preload(clause.Associations).Preload("Booking."+clause.Associations).Preload("User.Address").Where("id = ?", id).First(&history).Error; err != nil {
		return history, err
	}
	return history, nil
}

func (h *historiesRepository) GetHistoriesById(id string) ([]model.VaccineHistories, error) {
	var history []model.VaccineHistories
	if err := h.db.Preload(clause.Associations).Preload("Booking."+clause.Associations).Where("id = ?", id).Find(&history).Error; err != nil {
		return history, err
	}
	return history, nil
}

func (h *historiesRepository) GetHistoryByIdBooking(id string) ([]model.VaccineHistories, error) {
	var history []model.VaccineHistories
	if err := h.db.Preload("User.Address").Where("id_booking = ?", id).Find(&history).Error; err != nil {
		return history, err
	}
	return history, nil
}

func (h *historiesRepository) GetHistoryByIdSameBook(id string) ([]model.VaccineHistories, error) {
	var history []model.VaccineHistories
	if err := h.db.Preload("User.Address").Preload("Booking.Session").Where("id_same_book = ?", id).Find(&history).Error; err != nil {
		return history, err
	}
	return history, nil
}

func (h *historiesRepository) GetHistoryByNIK(id, nik string) (model.VaccineHistories, error) {
	var history model.VaccineHistories
	if err := h.db.Preload("User.Address").Where("nik_user = ? AND id_booking = ?", nik, id).First(&history).Error; err != nil {
		return history, err
	}
	return history, nil
}

func (h *historiesRepository) GetHistoriesByNIK(nik string) ([]model.VaccineHistories, error) {
	var history []model.VaccineHistories
	if err := h.db.Preload("User.Address").Where("nik_user = ?", nik).Find(&history).Error; err != nil {
		return history, err
	}
	return history, nil
}

func (h *historiesRepository) GetHistoriesByIdBooking(id string) ([]model.VaccineHistories, error) {
	var history []model.VaccineHistories
	if err := h.db.Preload("User.Address").Where("id_booking = ?", id).Find(&history).Error; err != nil {
		return history, err
	}
	return history, nil
}

func (h *historiesRepository) GetHistoriesByNIK(id string) ([]model.VaccineHistories, error) {
	var history []model.VaccineHistories
	if err := h.db.Preload("User.Address").Where("id_booking = ?", id).Find(&history).Error; err != nil {
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

func (h *historiesRepository) UpdateHistoryByNik(data model.VaccineHistories, nik, id string) (model.VaccineHistories, error) {
	var history model.VaccineHistories
	if err := h.db.Model(&history).Where("nik_user  = ? AND id_booking = ?", nik, id).Updates(&data).Error; err != nil {
		return history, err
	}
	return history, nil
}

func (h *historiesRepository) CheckVaccineCount(nik string) ([]model.VaccineHistories, error) {
	var history []model.VaccineHistories
	if err := h.db.Model(&history).Where("nik_user = ? AND status = ?", nik, "Attended").Find(&history).Error; err != nil {
		return history, err
	}
	return history, nil
}

func (h *historiesRepository) GetTotalUserVaccinated() (response.VaccinatedUser, error) {
	var history response.VaccinatedUser
	if err := h.db.Raw("SELECT id, COUNT(DISTINCT nik_user) AS vaccinated FROM vaccine_histories WHERE status = ?", "Attended").Scan(&history).Error; err != nil {
		return history, err
	}
	return history, nil
}
