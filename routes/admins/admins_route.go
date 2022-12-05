package admins

import (
	"os"
	"vaksin-id-be/controllers"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

func AdminUnauthenticated(routes *echo.Group, api *controllers.AdminController) {
	{
		routes.POST("/admin/login", api.LoginAdmin)
	}
}

func AdminAuthenticated(routes *echo.Group, api *controllers.AdminController) {
	authAdmin := routes.Group("/admin")
	authAdmin.Use(middleware.JWT([]byte(os.Getenv("SECRET_JWT_KEY_ADMIN"))))
	{
		authAdmin.GET("/profile", api.GetAdmins)
		authAdmin.GET("/all", api.GetAllAdmins)
		authAdmin.PUT("/profile", api.UpdateAdmins)
		authAdmin.DELETE("/profile/:id", api.DeleteAdmins)
	}
}
