package mysql

import (
	"vaksin-id-be/model"

	"gorm.io/gorm"
)

type VaccinesRepository interface {
	CreateVaccine(data model.Vaccines) error
	GetAllVaccines() ([]model.Vaccines, error)
	GetVaccinesById(id string) (model.Vaccines, error)
	GetVaccinesByIdAdmin(idhealthfacil string) ([]model.Vaccines, error)
	GetVaccineByName() ([]model.Vaccines, error)
	UpdateVaccine(data model.Vaccines, id string) error
	DeleteVacccine(id string) error
	CheckNameExist(idhealthfacil, name string) error
}

type vaccinesRepository struct {
	db *gorm.DB
}

func NewVaccinesRepository(db *gorm.DB) *vaccinesRepository {
	return &vaccinesRepository{db: db}
}

func (v *vaccinesRepository) CreateVaccine(data model.Vaccines) error {
	if err := v.db.Create(&data).Error; err != nil {
		return err
	}

	return nil
}

func (v *vaccinesRepository) GetAllVaccines() ([]model.Vaccines, error) {
	var vaccines []model.Vaccines
	if err := v.db.Model(&model.Vaccines{}).Order("name").Find(&vaccines).Error; err != nil {
		return vaccines, err
	}
	return vaccines, nil
}

func (v *vaccinesRepository) GetVaccinesById(id string) (model.Vaccines, error) {
	var vaccines model.Vaccines
	if err := v.db.Model(&model.Vaccines{}).Where("id = ?", id).Find(&vaccines).Error; err != nil {
		return vaccines, err
	}
	return vaccines, nil
}

func (v *vaccinesRepository) GetVaccinesByIdAdmin(idhealthfacil string) ([]model.Vaccines, error) {
	var vaccines []model.Vaccines
	if err := v.db.Model(&model.Vaccines{}).Where("id_health_facilities = ?", idhealthfacil).Find(&vaccines).Error; err != nil {
		return vaccines, err
	}
	return vaccines, nil
}

func (v *vaccinesRepository) GetVaccineByName() ([]model.Vaccines, error) {
	var vaccines []model.Vaccines
	if err := v.db.Raw("SELECT name, SUM(stock) AS stock FROM vaccines GROUP BY name").Scan(&vaccines).Error; err != nil {
		return vaccines, err
	}
	return vaccines, nil
}

func (v *vaccinesRepository) UpdateVaccine(data model.Vaccines, id string) error {
	if err := v.db.Model(&model.Vaccines{}).Where("id = ?", id).Updates(&data).Error; err != nil {
		return err
	}
	return nil
}

func (v *vaccinesRepository) DeleteVacccine(id string) error {
	var vaccines model.Vaccines
	if err := v.db.Where("id = ?", id).Delete(&vaccines).Error; err != nil {
		return err
	}
	return nil
}

func (v *vaccinesRepository) CheckNameExist(idhealthfacil, name string) error {
	var vaccines model.Vaccines
	if err := v.db.Model(&model.Vaccines{}).Where("id_health_facilities = ? AND name = ?", idhealthfacil, name).First(&vaccines).Error; err != nil {
		return err
	}
	return nil
}
