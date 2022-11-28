package config

import (
	"vaksin-id-be/model"

	"gorm.io/gorm"
)

func MigrateDB(db *gorm.DB) error {
	return db.AutoMigrate(
		model.Users{},
		model.HealthFacilities{},
		model.Admins{},
		model.BookingSessions{},
		model.Vaccines{},
		model.Addresses{},
		model.Sessions{},
		model.VaccineHistories{},
	)
}
