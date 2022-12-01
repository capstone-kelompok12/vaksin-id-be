package routes

import (
	"vaksin-id-be/config"
	m "vaksin-id-be/middleware"
	adm "vaksin-id-be/routes/admins"
	hf "vaksin-id-be/routes/health_facilities"
	users "vaksin-id-be/routes/users"

	_ "vaksin-id-be/docs"

	"github.com/labstack/echo/v4"
	echoSwagger "github.com/swaggo/echo-swagger"
)

func Init() *echo.Echo {
	dbConfig := config.InitGorm()

	// route config
	userApi := config.InitUserAPI(dbConfig)
	healthFacilitiesApi := config.InitHealthFacilitiesAPI(dbConfig)
	adminApi := config.InitAdminAPI(dbConfig)

	routes := echo.New()

	// middleware
	m.RemoveSlash(routes)
	m.LogMiddleware(routes)

	// v1
	// unauthenticated
	v1 := routes.Group("/api/v1")

	//swagger
	routes.GET("/swagger/*", echoSwagger.WrapHandler)

	// users
	users.UserUnauthenticated(v1, userApi)
	users.UserAuthenticated(v1, userApi)

	// health facilities
	hf.HealthFacilitiesAuthenticated(v1, healthFacilitiesApi)

	// admins
	adm.AdminUnauthenticated(v1, adminApi)
	adm.AdminAuthenticated(v1, adminApi)

	return routes
}
