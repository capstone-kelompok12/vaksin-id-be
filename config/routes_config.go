package config

import (
	"vaksin-id-be/controllers"
	mysql_user "vaksin-id-be/repository/mysql/users"
	services_user "vaksin-id-be/services/users"

	"gorm.io/gorm"
)

func InitUserAPI(db *gorm.DB) *controllers.UserController {
	userRepo := mysql_user.NewUserRepository(db)
	userServ := services_user.NewUserService(userRepo)
	userAPI := controllers.NewUserController(userServ)
	return userAPI
}
