package admins

import (
	"os"
	"vaksin-id-be/controllers"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

func UserUnauthenticated(routes *echo.Group, api *controllers.AdminsController) {
	{
		routes.POST("/admin/signup", api.RegisterAdmin)
		routes.POST("/admin/login", api.LoginAdmin)
	}
}

func UserAuthenticated(routes *echo.Group, api *controllers.AdminsController) {
	authAdmin := routes.Group("/admin")
	authAdmin.Use(middleware.JWT([]byte(os.Getenv("SECRET_JWT_KEY_ADMIN"))))
	{
		authAdmin.GET("/all", api.GetAllAdmin)
		authAdmin.GET("/:id", api.GetAdmin)
		authAdmin.PUT("/:id", api.UpdateAdmin)
		authAdmin.DELETE("/:id", api.DeleteAdmin)
	}
}
