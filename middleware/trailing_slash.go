package middleware

import (
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

func RemoveSlash(routes *echo.Echo) {
	routes.Pre(middleware.RemoveTrailingSlash())
}
