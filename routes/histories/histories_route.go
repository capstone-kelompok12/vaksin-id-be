package histories

import (
	"os"
	"vaksin-id-be/controllers"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

func HistoriesUnauthenticated(routes *echo.Group, api *controllers.HistoriesController) {
	{
		routes.GET("/dashboard/history", api.GetTotalUserVaccinated)
	}
}

func HistoriesAuthenticated(routes *echo.Group, api *controllers.HistoriesController) {
	authAdmin := routes.Group("/admin")
	authUser := routes.Group("/users")
	authAdmin.Use(middleware.JWT([]byte(os.Getenv("SECRET_JWT_KEY_ADMIN"))))
	authUser.Use(middleware.JWT([]byte(os.Getenv("SECRET_JWT_KEY"))))
	{
		authAdmin.POST("/histories", api.CreateHistory)
		authAdmin.GET("/histories", api.GetAllHistory)
		authAdmin.GET("/histories/:id", api.GetHistoryById)
		authAdmin.PUT("/histories/:id", api.UpdateHistory)

	}
}
