package mysqldb

import (
	"vaksin-id-be/model"

	"gorm.io/gorm"
)

type UserRepository interface {
	RegisterUser(data model.Users) error
	LoginUser(data model.Users) (model.Users, error)
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

func (u *userRepository) LoginUser(data model.Users) (model.Users, error) {
	var user model.Users

	createData := u.db.Where("email = ?", data.Email).First(&user)
	if err := createData.Error; err != nil {
		return user, err
	}
	return user, nil
}
