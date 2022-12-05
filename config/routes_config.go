package config

import (
	"vaksin-id-be/controllers"
	mysql_address "vaksin-id-be/repository/mysql/addresses"
	mysql_admin "vaksin-id-be/repository/mysql/admins"
	mysql_health "vaksin-id-be/repository/mysql/health_facilities"
	mysql_user "vaksin-id-be/repository/mysql/users"
	mysql_vaccines "vaksin-id-be/repository/mysql/vaccines"
	services_addresses "vaksin-id-be/services/addresses"
	services_admin "vaksin-id-be/services/admins"
	services_health "vaksin-id-be/services/health_facilities"
	services_user "vaksin-id-be/services/users"
	services_vaccines "vaksin-id-be/services/vaccines"

	"gorm.io/gorm"
)

func InitUserAPI(db *gorm.DB) *controllers.UserController {
	userRepo := mysql_user.NewUserRepository(db)
	addressRepo := mysql_address.NewAddressesRepository(db)
	healthRepo := mysql_health.NewHealthFacilitiesRepository(db)
	userServ := services_user.NewUserService(userRepo, addressRepo, healthRepo)
	addressServ := services_addresses.NewAddressesService(addressRepo)
	userAPI := controllers.NewUserController(userServ, addressServ)
	return userAPI
}

func InitHealthFacilitiesAPI(db *gorm.DB) *controllers.HealthFacilitiesController {
	healthRepo := mysql_health.NewHealthFacilitiesRepository(db)
	addressRepo := mysql_address.NewAddressesRepository(db)
	adminRepo := mysql_admin.NewAdminsRepository(db)
	healthServ := services_health.NewHealthFacilitiesService(healthRepo, addressRepo, adminRepo)
	addressServ := services_addresses.NewAddressesService(addressRepo)
	healthAPI := controllers.NewHealthFacilitiesController(healthServ, addressServ)
	return healthAPI
}

func InitAdminAPI(db *gorm.DB) *controllers.AdminController {
	adminRepo := mysql_admin.NewAdminsRepository(db)
	adminServ := services_admin.NewAdminService(adminRepo)
	adminAPI := controllers.NewAdminController(adminServ)
	return adminAPI
}

func InitVaccinesAPI(db *gorm.DB) *controllers.VaccinesController {
	vaccinesRepo := mysql_vaccines.NewVaccinesRepository(db)
	vaccinesServ := services_vaccines.NewVaccinesService(vaccinesRepo)
	vaccinesAPI := controllers.NewVaccinesController(vaccinesServ)
	return vaccinesAPI
}
