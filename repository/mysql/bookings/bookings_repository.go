package mysql

import (
	"vaksin-id-be/model"

	"gorm.io/gorm"
)

type BookingRepository interface {
	CreateBooking(data model.BookingSessions) error
	UpdateBooking(data model.BookingSessions, id string) error
	GetAllBooking() ([]model.BookingSessions, error)
	GetBooking(id string) (model.BookingSessions, error)
	DeleteBooking(id string) error
}

type bookingRepository struct {
	db *gorm.DB
}

func NewBookingRepository(db *gorm.DB) *bookingRepository {
	return &bookingRepository{db: db}
}

func (b *bookingRepository) CreateBooking(data model.BookingSessions) error {
	if err := b.db.Save(&data).Error; err != nil {
		return err
	}
	return nil
}

func (b *bookingRepository) UpdateBooking(data model.BookingSessions, id string) error {
	if err := b.db.Model(&model.BookingSessions{}).Where("id = ?", id).Updates(&data).Error; err != nil {
		return err
	}
	return nil
}

func (b *bookingRepository) GetAllBooking() ([]model.BookingSessions, error) {
	var booking []model.BookingSessions
	if err := b.db.Model(&model.BookingSessions{}).Find(&booking).Error; err != nil {
		return booking, err
	}
	return booking, nil
}

func (b *bookingRepository) GetBooking(id string) (model.BookingSessions, error) {
	var booking model.BookingSessions
	if err := b.db.Where("id = ?", id).First(&booking).Error; err != nil {
		return booking, err
	}
	return booking, nil
}

func (b *bookingRepository) DeleteBooking(id string) error {
	var booking model.BookingSessions
	if err := b.db.Where("id = ?", id).Delete(&booking).Error; err != nil {
		return err
	}
	return nil
}
