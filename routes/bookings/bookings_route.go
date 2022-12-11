package bookings

import (
	"os"
	"vaksin-id-be/controllers"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

func BookingsUnauthenticated(routes *echo.Group, api *controllers.BookingsController) {
	{
		routes.GET("/bookings", api.GetAllBookings)
		routes.GET("/bookings/:id", api.GetBooking)
		routes.POST("/bookings", api.CreateBooking)
		routes.PUT("/bookings/:id", api.UpdateBooking)
		routes.DELETE("/bookings/:id", api.DeleteBooking)
	}
}

func BookingsAuthenticated(routes *echo.Group, api *controllers.BookingsController) {
	authAdmin := routes.Group("/admin")
	authAdmin.Use(middleware.JWT([]byte(os.Getenv("SECRET_JWT_KEY_ADMIN"))))
	{
		// authAdmin.PUT("/bookings", api.UpdateHealthFacilities)
		// authAdmin.DELETE("/bookings/:id", api.DeleteHealthFacilities)
	}
}
