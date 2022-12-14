package config

import (
	"vaksin-id-be/controllers"
	mysql_address "vaksin-id-be/repository/mysql/addresses"
	mysql_admin "vaksin-id-be/repository/mysql/admins"
	mysql_bookings "vaksin-id-be/repository/mysql/bookings"
	mysql_health "vaksin-id-be/repository/mysql/health_facilities"
	mysql_histories "vaksin-id-be/repository/mysql/histories"
	mysql_sessions "vaksin-id-be/repository/mysql/sessions"
	mysql_user "vaksin-id-be/repository/mysql/users"
	mysql_vaccines "vaksin-id-be/repository/mysql/vaccines"
	services_addresses "vaksin-id-be/services/addresses"
	services_admin "vaksin-id-be/services/admins"
	services_bookings "vaksin-id-be/services/bookings"
	services_health "vaksin-id-be/services/health_facilities"
	services_histories "vaksin-id-be/services/histories"
	services_sessions "vaksin-id-be/services/sessions"
	services_user "vaksin-id-be/services/users"
	services_vaccines "vaksin-id-be/services/vaccines"

	"gorm.io/gorm"
)

func InitUserAPI(db *gorm.DB) *controllers.UserController {
	userRepo := mysql_user.NewUserRepository(db)
	vaccinesRepo := mysql_vaccines.NewVaccinesRepository(db)
	historyRepo := mysql_histories.NewHistoryRepository(db)
	addressRepo := mysql_address.NewAddressesRepository(db)
	bookingsRepo := mysql_bookings.NewBookingRepository(db)
	healthRepo := mysql_health.NewHealthFacilitiesRepository(db)
	sessionsRepo := mysql_sessions.NewSessionsRepository(db)
	userServ := services_user.NewUserService(userRepo, addressRepo, healthRepo, historyRepo, bookingsRepo, sessionsRepo, vaccinesRepo)
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

func InitSessionsAPI(db *gorm.DB) *controllers.SessionsController {
	sessionsRepo := mysql_sessions.NewSessionsRepository(db)
	vaccinesRepo := mysql_vaccines.NewVaccinesRepository(db)
	bookingsRepo := mysql_bookings.NewBookingRepository(db)
	historyRepo := mysql_histories.NewHistoryRepository(db)
	userRepo := mysql_user.NewUserRepository(db)
	sessionsServ := services_sessions.NewSessionsService(sessionsRepo, vaccinesRepo, bookingsRepo, userRepo, historyRepo)
	sessionsAPI := controllers.NewSessionsController(sessionsServ)
	return sessionsAPI
}

func InitBookingsAPI(db *gorm.DB) *controllers.BookingsController {
	bookingsRepo := mysql_bookings.NewBookingRepository(db)
	userRepo := mysql_user.NewUserRepository(db)
	historyRepo := mysql_histories.NewHistoryRepository(db)
	sessionsRepo := mysql_sessions.NewSessionsRepository(db)
	bookingsServ := services_bookings.NewBookingService(bookingsRepo, historyRepo, sessionsRepo, userRepo)
	bookingsAPI := controllers.NewBookingController(bookingsServ)
	return bookingsAPI
}

func InitHistoryAPI(db *gorm.DB) *controllers.HistoriesController {
	historyRepo := mysql_histories.NewHistoryRepository(db)
	historyServ := services_histories.NewHistoriesService(historyRepo)
	historyAPI := controllers.NewHistoriesController(historyServ)
	return historyAPI
}
