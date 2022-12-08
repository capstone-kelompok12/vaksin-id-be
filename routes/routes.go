package routes

import (
	"vaksin-id-be/config"
	m "vaksin-id-be/middleware"
	adm "vaksin-id-be/routes/admins"
	hf "vaksin-id-be/routes/health_facilities"
	sessions "vaksin-id-be/routes/sessions"
	users "vaksin-id-be/routes/users"
	vac "vaksin-id-be/routes/vaccines"

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
	vaccineApi := config.InitVaccinesAPI(dbConfig)
	sessionApi := config.InitSessionsAPI(dbConfig)

	routes := echo.New()

	// middleware
	m.EchoCors(routes)
	m.RemoveSlash(routes)
	m.LogMiddleware(routes)
	m.RecoverEcho(routes)

	// v1
	// unauthenticated
	v1 := routes.Group("/api/v1")

	//swagger
	routes.GET("/swagger/*", echoSwagger.WrapHandler)

	// users
	users.UserUnauthenticated(v1, userApi)
	users.UserAuthenticated(v1, userApi)

	// health facilities
	hf.HealthFacilitiesUnauthenticated(v1, healthFacilitiesApi)
	hf.HealthFacilitiesAuthenticated(v1, healthFacilitiesApi)

	// admins
	adm.AdminUnauthenticated(v1, adminApi)
	adm.AdminAuthenticated(v1, adminApi)

	// vaccines
	vac.VaccinesUnauthenticated(v1, vaccineApi)
	vac.VaccinesAuthenticated(v1, vaccineApi)

	// sessions
	sessions.SessionsUnauthenticated(v1, sessionApi)
	sessions.SessionsAuthenticated(v1, sessionApi)

	return routes
}
