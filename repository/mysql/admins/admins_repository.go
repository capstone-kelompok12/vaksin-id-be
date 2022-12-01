package mysql

import (
	"vaksin-id-be/model"

	"gorm.io/gorm"
)

type AdminsRepository interface {
	RegisterAdmins(data model.Admins) error
	LoginAdmins(data model.Admins) (model.Admins, error)
	GetAllAdmins() ([]model.Admins, error)
	GetAdmins(id string) (model.Admins, error)
	UpdateAdmins(data model.Admins, id string) error
	DeleteAdmins(id string) error
	DeleteAdminsByHealth(id string) error
}

type adminsRepository struct {
	db *gorm.DB
}

func NewAdminsRepository(db *gorm.DB) *adminsRepository {
	return &adminsRepository{db: db}
}

func (a *adminsRepository) RegisterAdmins(data model.Admins) error {
	if err := a.db.Create(&data).Error; err != nil {
		return err
	}
	return nil
}

func (a *adminsRepository) LoginAdmins(data model.Admins) (model.Admins, error) {
	var admins model.Admins

	if err := a.db.Where("email = ?", data.Email).First(&admins).Error; err != nil {
		return admins, err
	}
	return admins, nil
}

func (a *adminsRepository) GetAllAdmins() ([]model.Admins, error) {
	var admins []model.Admins
	if err := a.db.Model(&model.Admins{}).Find(&admins).Error; err != nil {
		return admins, err
	}
	return admins, nil
}

func (a *adminsRepository) GetAdmins(id string) (model.Admins, error) {
	var admins model.Admins
	if err := a.db.Where("id = ?", id).First(&admins).Error; err != nil {
		return admins, err
	}
	return admins, nil
}

func (a *adminsRepository) UpdateAdmins(data model.Admins, id string) error {
	if err := a.db.Model(&model.Admins{}).Where("id = ?", id).Updates(&data).Error; err != nil {
		return err
	}
	return nil
}

func (a *adminsRepository) DeleteAdmins(id string) error {
	var admins model.Admins
	if err := a.db.Where("id = ?", id).Delete(&admins).Error; err != nil {
		return err
	}
	return nil
}

func (a *adminsRepository) DeleteAdminsByHealth(id string) error {
	var admins model.Admins
	if err := a.db.Where("id_health_facilities = ?", id).Delete(&admins).Error; err != nil {
		return err
	}
	return nil
}
