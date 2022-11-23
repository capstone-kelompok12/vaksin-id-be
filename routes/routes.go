package routes

import (
	"net/http"
	"vaksin-id-be/config"

	"github.com/labstack/echo/v4"
)

func Init() *echo.Echo {
	config.InitGorm()

	routes := echo.New()

	v1 := routes.Group("/api/v1")

	v1.GET("/user", func(c echo.Context) error {
		return c.JSON(http.StatusOK, map[string]interface{}{
			"message": "hello user test",
		})
	})

	return routes
}
