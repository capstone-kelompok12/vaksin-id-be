package mysql

import (
	"vaksin-id-be/dto/response"
	"vaksin-id-be/model"

	"gorm.io/gorm"
	"gorm.io/gorm/clause"
)

type UserRepository interface {
	RegisterUser(data model.Users) error
	CheckExistNik(nik string) (model.Users, error)
	ReactivatedUser(nik string) error
	ReactivatedUpdateUser(data model.Users, nik string) error
	ReactivatedAddress(nik string) error
	LoginUser(data model.Users) (model.Users, error)
	GetUserDataByNik(nik string) (model.Users, error)
	GetUserDataByNikNoAddress(nik string) (model.Users, error)
	GetUserHistoryByNik(nik string) (model.Users, error)
	UpdateUserProfile(data model.Users) error
	GetAgeUser(data model.Users) (response.AgeUser, error)
	DeleteUser(nik string) error
	NearbyHealthFacilities(city string) ([]model.Addresses, error)
	// UpdateVaccineHistory()
}

type userRepository struct {
	db *gorm.DB
}

func NewUserRepository(db *gorm.DB) *userRepository {
	return &userRepository{db: db}
}

func (u *userRepository) RegisterUser(data model.Users) error {
	if err := u.db.Create(&data).Error; err != nil {
		return err
	}
	return nil
}

func (u *userRepository) CheckExistNik(nik string) (model.Users, error) {
	var user model.Users
	if err := u.db.Unscoped().Where("nik = ?", nik).First(&user).Error; err != nil {
		return user, err
	}
	return user, nil
}

func (u *userRepository) ReactivatedUser(nik string) error {
	var user model.Users
	if err := u.db.Unscoped().Model(&user).Where("nik = ?", nik).Update("deleted_at", nil).Error; err != nil {
		return err
	}
	return nil
}

func (u *userRepository) ReactivatedUpdateUser(data model.Users, nik string) error {
	var user model.Users
	if err := u.db.Unscoped().Model(&user).Where("nik = ?", nik).Updates(&data).Error; err != nil {
		return err
	}
	return nil
}

func (u *userRepository) ReactivatedAddress(nik string) error {
	var address model.Addresses
	if err := u.db.Unscoped().Model(&address).Where("nik_user = ?", nik).Update("deleted_at", nil).Error; err != nil {
		return err
	}
	return nil
}

func (u *userRepository) LoginUser(data model.Users) (model.Users, error) {
	var user model.Users

	if err := u.db.Where("email = ?", data.Email).First(&user).Error; err != nil {
		return user, err
	}
	return user, nil
}

func (u *userRepository) GetUserDataByNik(nik string) (model.Users, error) {
	var user model.Users

	if err := u.db.Preload("Address").Where("nik = ?", nik).First(&user).Error; err != nil {
		return user, err
	}
	return user, nil
}

func (u *userRepository) GetUserHistoryByNik(nik string) (model.Users, error) {
	var user model.Users

	if err := u.db.Preload(clause.Associations).Preload("History."+clause.Associations).Preload("History.Booking.Session.Vaccine").Preload("Address").Where("nik = ?", nik).First(&user).Error; err != nil {
		return user, err
	}
	return user, nil
}

func (u *userRepository) GetUserDataByNikNoAddress(nik string) (model.Users, error) {
	var user model.Users

	if err := u.db.Where("nik = ?", nik).First(&user).Error; err != nil {
		return user, err
	}
	return user, nil
}

func (u *userRepository) UpdateUserProfile(data model.Users) error {
	var user model.Users

	if err := u.db.Model(&user).Where("nik = ?", data.NIK).Updates(&data).Error; err != nil {
		return err
	}
	return nil
}

func (u *userRepository) GetAgeUser(data model.Users) (response.AgeUser, error) {
	var age response.AgeUser

	if err := u.db.Raw("SELECT birth_date, DATE_FORMAT(FROM_DAYS(DATEDIFF(NOW(), birth_date)), '%Y') + 0 AS age FROM users").Scan(&age).Error; err != nil {
		return age, err
	}
	return age, nil
}

func (u *userRepository) DeleteUser(nik string) error {
	var user model.Users
	if err := u.db.Where("nik = ?", nik).Find(&user).Delete(&user).Error; err != nil {
		return err
	}
	return nil
}

func (u *userRepository) NearbyHealthFacilities(city string) ([]model.Addresses, error) {
	var address []model.Addresses
	if err := u.db.Where("city = ? AND nik_user = ?", city, nil).Find(&address).Error; err != nil {
		return address, err
	}
	return address, nil
}
