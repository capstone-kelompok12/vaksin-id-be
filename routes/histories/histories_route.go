package histories

import (
	"os"
	"vaksin-id-be/controllers"

	"github.com/labstack/echo"
	"github.com/tiny-go/middleware"
)

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
