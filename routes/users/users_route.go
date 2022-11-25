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
	routes.Use(middleware.JWT([]byte(os.Getenv("SECRET_JWT_KEY"))))
	{
		routes.GET("/profile", api.GetUserDataByNik)
		routes.PUT("/profile", api.UpdateUser)
		routes.DELETE("/profile", api.DeleteUser)
	}
}
