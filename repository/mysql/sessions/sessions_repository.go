package mysql

import (
	"vaksin-id-be/model"

	"gorm.io/gorm"
	"gorm.io/gorm/clause"
)

type SessionsRepository interface {
	CreateSession(data model.Sessions) error
	GetAllSessions() ([]model.Sessions, error)
	GetSessionsByAdmin(auth string) ([]model.Sessions, error)
	GetSessionAdminById(auth, id string) (model.Sessions, error)
	UpdateSession(data model.Sessions, id string) error
	CloseSession(data model.Sessions, id string) error
	DeleteSession(id string) error
}

type sessionsRepository struct {
	db *gorm.DB
}

func NewSessionsRepository(db *gorm.DB) *sessionsRepository {
	return &sessionsRepository{db: db}
}

func (s *sessionsRepository) CreateSession(data model.Sessions) error {
	if err := s.db.Create(&data).Error; err != nil {
		return err
	}

	return nil
}

func (s *sessionsRepository) GetAllSessions() ([]model.Sessions, error) {
	var session []model.Sessions
	if err := s.db.Preload(clause.Associations).Preload("Booking." + clause.Associations).Model(&model.Sessions{}).Find(&session).Error; err != nil {
		return session, err
	}

	return session, nil
}

func (s *sessionsRepository) GetSessionAdminById(auth, id string) (model.Sessions, error) {
	var session model.Sessions
	if err := s.db.Preload(clause.Associations).Preload("Booking."+clause.Associations).Where("id_health_facilities = ? AND id = ?", auth, id).First(&session).Error; err != nil {
		return session, err
	}
	return session, nil
}

func (s *sessionsRepository) GetSessionsByAdmin(auth string) ([]model.Sessions, error) {
	var session []model.Sessions
	if err := s.db.Preload(clause.Associations).Preload("Booking."+clause.Associations).Where("id_health_facilities = ?", auth).Find(&session).Error; err != nil {
		return session, err
	}
	return session, nil
}

func (s *sessionsRepository) UpdateSession(data model.Sessions, id string) error {
	if err := s.db.Model(&model.Sessions{}).Where("id = ?", id).Updates(&data).Error; err != nil {
		return err
	}
	return nil
}

func (s *sessionsRepository) CloseSession(data model.Sessions, id string) error {
	if err := s.db.Model(&model.Sessions{}).Where("id = ?", id).Updates(&data).Error; err != nil {
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
