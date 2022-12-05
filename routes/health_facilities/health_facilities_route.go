package healthfacilities

import (
	"os"
	"vaksin-id-be/controllers"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

func HealthFacilitiesUnauthenticated(routes *echo.Group, api *controllers.HealthFacilitiesController) {
	{
		routes.GET("/healthfacilities", api.GetAllHealthFacilities)
		routes.GET("/healthfacilities/:name", api.GetHealthFacilities)
		routes.POST("/healthfacilities", api.CreateHealthFacilities)
	}
}

func HealthFacilitiesAuthenticated(routes *echo.Group, api *controllers.HealthFacilitiesController) {
	authAdmin := routes.Group("/admin")
	authAdmin.Use(middleware.JWT([]byte(os.Getenv("SECRET_JWT_KEY_ADMIN"))))
	{
		authAdmin.PUT("/healthfacilities/:id", api.UpdateHealthFacilities)
		authAdmin.DELETE("/healthfacilities/:id", api.DeleteHealthFacilities)
	}
}
