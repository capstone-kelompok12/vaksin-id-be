package mysql

import (
	"vaksin-id-be/model"

	"gorm.io/gorm"
)

type AdminRepository interface {
	RegisterAdmin(data model.Admins) error
	LoginAdmin(data model.Admins) (model.Admins, error)
	GetAllAdmin() ([]model.Admins, error)
	GetAdmin(id string) (model.Admins, error)
	UpdateAdmin(data model.Admins, id string) error
	DeleteAdmin(id string) error
}

type adminRepository struct {
	db *gorm.DB
}

func NewAdminRepository(db *gorm.DB) *adminRepository {
	return &adminRepository{db: db}
}

func (a *adminRepository) RegisterAdmin(data model.Admins) error {
	if err := a.db.Create(&data).Error; err != nil {
		return err
	}
	return nil
}

func (a *adminRepository) LoginAdmin(data model.Admins) (model.Admins, error) {
	var admin model.Admins

	if err := a.db.Where("email = ?", data.Email).First(&admin).Error; err != nil {
		return admin, err
	}
	return admin, nil
}

func (a *adminRepository) GetAllAdmin() ([]model.Admins, error) {
	var admins []model.Admins
	if err := a.db.Find(&admins).Error; err != nil {
		return admins, err
	}
	return admins, nil
}

func (a *adminRepository) GetAdmin(id string) (model.Admins, error) {
	var admin model.Admins
	if err := a.db.Where("id = ?", id).First(&admin).Error; err != nil {
		return admin, err
	}
	return admin, nil
}

func (a *adminRepository) UpdateAdmin(data model.Admins, id string) error {
	var admin model.Admins

	if err := a.db.Model(&admin).Where("id = ?", id).Updates(&data).Error; err != nil {
		return err
	}
	return nil
}

func (a *adminRepository) DeleteAdmin(id string) error {
	var admin model.Admins
	if err := a.db.Where("id = ?", id).Delete(&admin).Error; err != nil {
		return err
	}
	return nil
}
