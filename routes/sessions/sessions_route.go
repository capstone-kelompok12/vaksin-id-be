package sessions

import (
	"os"
	"vaksin-id-be/controllers"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

func SessionsUnauthenticated(routes *echo.Group, api *controllers.SessionsController) {
	{
		routes.GET("/sessions", api.GetAllSessions)
	}
}

func SessionsAuthenticated(routes *echo.Group, api *controllers.SessionsController) {
	authAdmin := routes.Group("/admin")
	authAdmin.Use(middleware.JWT([]byte(os.Getenv("SECRET_JWT_KEY_ADMIN"))))
	{
		authAdmin.POST("/sessions", api.CreateSession)
		authAdmin.GET("/sessions", api.GetSessionByAdmin)
		authAdmin.GET("/sessions/:id", api.GetSessionsAdminById)
		authAdmin.PUT("/sessions/:id", api.UpdateSession)
		authAdmin.PUT("/sessions/status/:id", api.IsCloseSession)
		authAdmin.DELETE("/sessions/:id", api.DeleteSession)
	}
}
