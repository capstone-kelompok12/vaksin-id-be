package healthfacilities

import (
	"vaksin-id-be/controllers"

	"github.com/labstack/echo/v4"
)

func HealthFacilitiesAuthenticated(routes *echo.Group, api *controllers.HealthFacilitiesController) {
	authUser := routes.Group("/admin")
	// authUser.Use(middleware.JWT([]byte(os.Getenv("SECRET_JWT_KEY_ADMIN"))))
	{
		authUser.POST("/healthfacilities", api.CreateHealthFacilities)
		authUser.PUT("/healthfacilities/:id", api.UpdateHealthFacilities)
		authUser.DELETE("/healthfacilities/:id", api.DeleteHealthFacilities)
	}
}
