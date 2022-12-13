package vaccines

import (
	"os"
	"vaksin-id-be/controllers"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

func VaccinesUnauthenticated(routes *echo.Group, api *controllers.VaccinesController) {
	{
		routes.GET("/vaccines", api.GetAllVaccines)
		routes.GET("/vaccines/dashboard", api.GetVaccineDashboard)
		routes.GET("/vaccines/all", api.GetAllVaccinesCount)
	}
}

func VaccinesAuthenticated(routes *echo.Group, api *controllers.VaccinesController) {
	authAdmin := routes.Group("/admin")
	authAdmin.Use(middleware.JWT([]byte(os.Getenv("SECRET_JWT_KEY_ADMIN"))))
	{
		authAdmin.POST("/vaccines", api.CreateVaccine)
		authAdmin.GET("/vaccines", api.GetVaccineByAdmin)
		authAdmin.PUT("/vaccines/:id", api.UpdateVaccine)
		authAdmin.DELETE("/vaccines/:id", api.DeleteVacccine)
	}
}
