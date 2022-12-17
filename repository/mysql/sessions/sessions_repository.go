package mysql

import (
	"vaksin-id-be/dto/response"
	"vaksin-id-be/model"

	"gorm.io/gorm"
	"gorm.io/gorm/clause"
)

type SessionsRepository interface {
	CreateSession(data model.Sessions) (model.Sessions, error)
	GetAllSessions() ([]model.Sessions, error)
	GetSumOfCapacity(id string) (response.SessionSumCap, error)
	GetSessionById(id string) (model.Sessions, error)
	GetSessionsByAdmin(auth string) ([]model.Sessions, error)
	GetAllFinishedSessionCount() (response.SessionFinished, error)
	UpdateSession(data model.Sessions, id string) error
	CloseSession(data model.Sessions, id string) error
	DeleteSession(id string) error
	IsCloseFalse() (response.IsCloseFalse, error)
}

type sessionsRepository struct {
	db *gorm.DB
}

func NewSessionsRepository(db *gorm.DB) *sessionsRepository {
	return &sessionsRepository{db: db}
}

func (s *sessionsRepository) CreateSession(data model.Sessions) (model.Sessions, error) {
	var session model.Sessions
	if err := s.db.Create(&data).Error; err != nil {
		return session, err
	}

	return data, nil
}

func (s *sessionsRepository) GetAllSessions() ([]model.Sessions, error) {
	var session []model.Sessions
	if err := s.db.Preload(clause.Associations).Preload("Booking." + clause.Associations).Preload("Vaccine").Model(&model.Sessions{}).Find(&session).Error; err != nil {
		return session, err
	}

	return session, nil
}

func (s *sessionsRepository) GetSumOfCapacity(id string) (response.SessionSumCap, error) {
	var session response.SessionSumCap
	if err := s.db.Preload("Vaccine").Raw("SELECT id, SUM(capacity) AS total_capacity FROM sessions WHERE id_vaccine = ?", id).Scan(&session).Error; err != nil {
		return session, err
	}
	return session, nil
}

func (s *sessionsRepository) GetSessionById(id string) (model.Sessions, error) {
	var session model.Sessions
	if err := s.db.Preload("Vaccine").Preload("Booking.User").Where("id = ?", id).First(&session).Error; err != nil {
		return session, err
	}
	return session, nil
}

func (s *sessionsRepository) GetAllFinishedSessionCount() (response.SessionFinished, error) {
	var session response.SessionFinished
	if err := s.db.Raw("SELECT id, COUNT(is_close) AS amount FROM sessions WHERE is_close = ?", true).Scan(&session).Error; err != nil {
		return session, err
	}
	return session, nil
}

func (s *sessionsRepository) GetSessionsByAdmin(auth string) ([]model.Sessions, error) {
	var session []model.Sessions
	if err := s.db.Preload(clause.Associations).Preload("Booking."+clause.Associations).Preload("Vaccine").Joins("Vaccine").Where("Vaccine.id_health_facilities = ?", auth).Find(&session).Error; err != nil {
		return session, err
	}
	return session, nil
}

func (s *sessionsRepository) UpdateSession(data model.Sessions, id string) error {
	if err := s.db.Preload("Vaccine").Model(&model.Sessions{}).Where("id = ?", id).Updates(&data).Error; err != nil {
		return err
	}
	return nil
}

func (s *sessionsRepository) CloseSession(data model.Sessions, id string) error {
	if err := s.db.Preload("Vaccine").Model(&model.Sessions{}).Where("id = ?", id).Updates(&data).Error; err != nil {
		return err
	}
	return nil
}

func (s *sessionsRepository) DeleteSession(id string) error {
	var session model.Sessions
	if err := s.db.Where("id = ?", id).Delete(&session).Error; err != nil {
		return err
	}
	return nil
}

func (s *sessionsRepository) IsCloseFalse() (response.IsCloseFalse, error) {
	var active response.IsCloseFalse

	var count int64

	s.db.Model(&model.Sessions{}).Where("is_close = ?", false).Count(&count)

	active.Active = int(count)
	return active, nil
}
