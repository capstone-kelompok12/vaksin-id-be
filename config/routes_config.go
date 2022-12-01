package config

import (
	"vaksin-id-be/controllers"
	mysql_address "vaksin-id-be/repository/mysql/addresses"
	mysql_health "vaksin-id-be/repository/mysql/health_facilities"
	mysql_user "vaksin-id-be/repository/mysql/users"
	services_addresses "vaksin-id-be/services/addresses"
	services_health "vaksin-id-be/services/health_facilities"
	services_user "vaksin-id-be/services/users"

	"gorm.io/gorm"
)

func InitUserAPI(db *gorm.DB) *controllers.UserController {
	userRepo := mysql_user.NewUserRepository(db)
	addressRepo := mysql_address.NewAddressesRepository(db)
	userServ := services_user.NewUserService(userRepo, addressRepo)
	addressServ := services_addresses.NewAddressesService(addressRepo)
	userAPI := controllers.NewUserController(userServ, addressServ)
	return userAPI
}

func InitHealthFacilitiesAPI(db *gorm.DB) *controllers.HealthFacilitiesController {
	healthRepo := mysql_health.NewHealthFacilitiesRepository(db)
	addressRepo := mysql_address.NewAddressesRepository(db)
	healthServ := services_health.NewHealthFacilitiesService(healthRepo, addressRepo)
	addressServ := services_addresses.NewAddressesService(addressRepo)
	healthAPI := controllers.NewHealthFacilitiesController(healthServ, addressServ)
	return healthAPI
}
