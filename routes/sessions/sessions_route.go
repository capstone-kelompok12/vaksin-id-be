package sessions

import (
	"os"
	"vaksin-id-be/controllers"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

func SessionsUnauthenticated(routes *echo.Group, api *controllers.SessionsController) {
	{
		routes.GET("/dashboard/sessions", api.GetSessionActive)
		routes.GET("/dashboard/sessions/amount", api.GetAllFinishedSessionCount)
	}
}

func SessionsAuthenticated(routes *echo.Group, api *controllers.SessionsController) {
	authAdmin := routes.Group("/admin")
	authUser := routes.Group("/users")
	authUser.Use(middleware.JWT([]byte(os.Getenv("SECRET_JWT_KEY"))))
	authAdmin.Use(middleware.JWT([]byte(os.Getenv("SECRET_JWT_KEY_ADMIN"))))
	{
		// authAdmin.GET("/sessions", api.GetSessionByAdmin)
		// admin
		authAdmin.GET("/sessions", api.GetAllSessions)
		authAdmin.GET("/sessions/:id", api.GetSessionsById)
		authAdmin.POST("/sessions", api.CreateSession)
		authAdmin.PUT("/sessions/:id", api.UpdateSession)
		authAdmin.PUT("/sessions/:id/close", api.IsCloseSession)
		authAdmin.DELETE("/sessions/:id", api.DeleteSession)
		// users
		authUser.GET("/sessions", api.GetAllSessions)
		authUser.GET("/sessions/:id", api.GetSessionsById)
	}
}
