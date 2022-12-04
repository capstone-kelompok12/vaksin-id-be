package mysql

import (
	"vaksin-id-be/model"

	"gorm.io/gorm"
)

type HealthFacilitiesRepository interface {
	CreateHealthFacilities(data model.HealthFacilities) error
	GetAllHealthFacilities() ([]model.HealthFacilities, error)
	GetAllHealthFacilitiesByCity(city string) ([]model.HealthFacilities, error)
	GetHealthFacilities(name string) (model.HealthFacilities, error)
	UpdateHealthFacilities(data model.HealthFacilities, id string) error
	DeleteHealthFacilities(id string) error
}

type healthFacilitiesRepository struct {
	db *gorm.DB
}

func NewHealthFacilitiesRepository(db *gorm.DB) *healthFacilitiesRepository {
	return &healthFacilitiesRepository{db: db}
}

func (h *healthFacilitiesRepository) CreateHealthFacilities(data model.HealthFacilities) error {
	if err := h.db.Create(&data).Error; err != nil {
		return err
	}
	return nil
}

func (h *healthFacilitiesRepository) GetAllHealthFacilities() ([]model.HealthFacilities, error) {
	var healthFacils []model.HealthFacilities
	if err := h.db.Preload("Address").Model(&model.HealthFacilities{}).Find(&healthFacils).Error; err != nil {
		return healthFacils, err
	}
	return healthFacils, nil
}
func (h *healthFacilitiesRepository) GetAllHealthFacilitiesByCity(city string) ([]model.HealthFacilities, error) {
	var healthFacils []model.HealthFacilities
	likeCity := "%" + city + "%"
	if err := h.db.Preload("Address").Joins("Address").Where("Address.city LIKE ?", likeCity).Find(&healthFacils).Error; err != nil {
		return healthFacils, err
	}
	return healthFacils, nil
}

func (h *healthFacilitiesRepository) GetHealthFacilities(name string) (model.HealthFacilities, error) {
	var healthFacil model.HealthFacilities
	likeName := "%" + name + "%"
	if err := h.db.Where("name LIKE ?", likeName).First(&healthFacil).Error; err != nil {
		return healthFacil, err
	}
	return healthFacil, nil
}

func (h *healthFacilitiesRepository) UpdateHealthFacilities(data model.HealthFacilities, id string) error {
	if err := h.db.Model(&model.HealthFacilities{}).Where("id = ?", id).Updates(&data).Error; err != nil {
		return err
	}
	return nil
}

func (h *healthFacilitiesRepository) DeleteHealthFacilities(id string) error {
	var healthFacil model.HealthFacilities
	if err := h.db.Where("id = ?", id).Delete(&healthFacil).Error; err != nil {
		return err
	}
	return nil
}
