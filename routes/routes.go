package routes

import (
	"vaksin-id-be/config"
	m "vaksin-id-be/middleware"
	users "vaksin-id-be/routes/users"

	"github.com/labstack/echo/v4"
)

func Init() *echo.Echo {
	dbConfig := config.InitGorm()

	// route config
	userApi := config.InitUserAPI(dbConfig)

	routes := echo.New()

	// middleware
	m.RemoveSlash(routes)
	m.LogMiddleware(routes)

	// v1
	// unauthenticated
	v1 := routes.Group("/api/v1")

	// users
	users.UserUnauthenticated(v1, userApi)
	users.UserAuthenticated(v1, userApi)

	return routes
}
