package mysql

import (
	"vaksin-id-be/model"

	"gorm.io/gorm"
)

type AddressesRepository interface {
	CreateAddress(data model.Addresses) error
	UpdateAddressUser(data model.Addresses, nik string) error
	UpdateAddressHealthFacilities(data model.Addresses, id string) error
	GetAddressUser(nik string) (model.Addresses, error)
	GetAddressHealthFacilities(id string) (model.Addresses, error)
	DeleteAddressUser(nik string) error
	DeleteAddressHealthFacilities(id string) error
}

type addressesRepository struct {
	db *gorm.DB
}

func NewAddressesRepository(db *gorm.DB) *addressesRepository {
	return &addressesRepository{db: db}
}

func (a *addressesRepository) CreateAddress(data model.Addresses) error {
	if err := a.db.Save(&data).Error; err != nil {
		return err
	}
	return nil
}

func (a *addressesRepository) UpdateAddressUser(data model.Addresses, nik string) error {
	var address model.Addresses

	if err := a.db.Model(&address).Where("nik_user = ?", nik).Updates(&data).Error; err != nil {
		return err
	}
	return nil
}

func (a *addressesRepository) UpdateAddressHealthFacilities(data model.Addresses, id string) error {
	var address model.Addresses

	if err := a.db.Model(&address).Where("id_health_facilities = ?", id).Updates(&data).Error; err != nil {
		return err
	}
	return nil
}

func (a *addressesRepository) GetAddressUser(nik string) (model.Addresses, error) {
	var address model.Addresses

	if err := a.db.Where("nik_user = ?", nik).First(&address).Error; err != nil {
		return address, err
	}
	return address, nil
}

func (a *addressesRepository) GetAddressHealthFacilities(id string) (model.Addresses, error) {
	var address model.Addresses

	if err := a.db.Where("id_health_facilities = ?", id).First(&address).Error; err != nil {
		return address, err
	}
	return address, nil
}

func (a *addressesRepository) DeleteAddressUser(nik string) error {
	var address model.Addresses

	if err := a.db.Where("nik_user = ?", nik).Find(&address).Unscoped().Delete(&address).Error; err != nil {
		return err
	}
	return nil
}

func (a *addressesRepository) DeleteAddressHealthFacilities(id string) error {
	var address model.Addresses

	if err := a.db.Where("id_health_facilities = ?", id).Find(&address).Delete(&address).Error; err != nil {
		return err
	}
	return nil
}
