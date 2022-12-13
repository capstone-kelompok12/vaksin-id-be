package users

import (
	"os"
	"vaksin-id-be/controllers"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

func UserUnauthenticated(routes *echo.Group, api *controllers.UserController) {
	{
		routes.POST("/signup", api.RegisterUser)
		routes.POST("/login", api.LoginUser)
	}
}

func UserAuthenticated(routes *echo.Group, api *controllers.UserController) {
	authUser := routes.Group("/profile")
	authUser.Use(middleware.JWT([]byte(os.Getenv("SECRET_JWT_KEY"))))
	{
		authUser.GET("", api.GetUserDataByNik)
		authUser.GET("/address", api.GetUserAddress)
		authUser.GET("/check/:nik", api.GetUserDataByNikCheck)
		authUser.PUT("", api.UpdateUser)
		authUser.DELETE("", api.DeleteUser)
		authUser.PUT("/address", api.UpdateUserAddress)
		authUser.POST("/nearby", api.UserNearbyHealth)
	}
}
