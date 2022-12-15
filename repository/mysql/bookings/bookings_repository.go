package mysql

import (
	"vaksin-id-be/model"

	"gorm.io/gorm"
	"gorm.io/gorm/clause"
)

type BookingRepository interface {
	CreateBooking(data model.BookingSessions) error
	UpdateBooking(data model.BookingSessions) error
	UpdateBookingAcc(data model.BookingSessions) (model.BookingSessions, error)
	GetAllBooking() ([]model.BookingSessions, error)
	GetBooking(id string) (model.BookingSessions, error)
	GetBookingBySession(id string) ([]model.BookingSessions, error)
	GetBookingBySessionDen(id string) ([]model.BookingSessions, error)
	FindMaxQueue(is_session string) (model.BookingSessions, error)
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

func (b *bookingRepository) UpdateBooking(data model.BookingSessions) error {
	if err := b.db.Preload("Session.Vaccine").Preload("History").Model(&model.BookingSessions{}).Where("id_session = ? AND id = ?", data.IdSession, data.ID).Updates(&data).Error; err != nil {
		return err
	}
	return nil
}

func (b *bookingRepository) UpdateBookingAcc(data model.BookingSessions) (model.BookingSessions, error) {
	var booking model.BookingSessions
	if err := b.db.Preload("Session.Vaccine").Preload("History").Model(&booking).Where("id = ?", data.ID).Updates(&data).Error; err != nil {
		return booking, err
	}
	return booking, nil
}

func (b *bookingRepository) GetAllBooking() ([]model.BookingSessions, error) {
	var booking []model.BookingSessions
	if err := b.db.Preload(clause.Associations).Preload("Session.Vaccine").Preload("History." + clause.Associations).Model(&model.BookingSessions{}).Find(&booking).Error; err != nil {
		return booking, err
	}
	return booking, nil
}

func (b *bookingRepository) GetBooking(id string) (model.BookingSessions, error) {
	var booking model.BookingSessions
	if err := b.db.Preload("Session.Vaccine").Where("id = ?", id).First(&booking).Error; err != nil {
		return booking, err
	}
	return booking, nil
}

func (b *bookingRepository) GetBookingBySession(id string) ([]model.BookingSessions, error) {
	var booking []model.BookingSessions
	if err := b.db.Preload("Session.Vaccine").Where("id_session = ? AND NOT status = ?", id, "Rejected").Find(&booking).Error; err != nil {
		return booking, err
	}
	return booking, nil
}

func (b *bookingRepository) GetBookingBySessionDen(id string) ([]model.BookingSessions, error) {
	var booking []model.BookingSessions
	if err := b.db.Preload("Session.Vaccine").Where("id_session = ? AND NOT status = ?", id, "Rejected").Find(&booking).Error; err != nil {
		return booking, err
	}
	return booking, nil
}

func (b *bookingRepository) FindMaxQueue(id_session string) (model.BookingSessions, error) {
	var booking model.BookingSessions
	if err := b.db.Model(&booking).Where("id_session = ?", id_session).Order("queue desc").First(&booking).Error; err != nil {
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
