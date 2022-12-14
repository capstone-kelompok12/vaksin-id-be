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
		routes.GET("/dashboard/bookings", api.GetBookingDashboard)
		routes.DELETE("/bookings/:id", api.DeleteBooking)
	}
}

func BookingsAuthenticated(routes *echo.Group, api *controllers.BookingsController) {
	authAdmin := routes.Group("/admin")
	authUser := routes.Group("/users")
	authAdmin.Use(middleware.JWT([]byte(os.Getenv("SECRET_JWT_KEY_ADMIN"))))
	authUser.Use(middleware.JWT([]byte(os.Getenv("SECRET_JWT_KEY"))))
	{
		authAdmin.PUT("/bookings/acc", api.UpdateBooking)
		authAdmin.PUT("/bookings/history/acc", api.UpdateAccAttendend)
		authUser.POST("/bookings", api.CreateBooking)
		authUser.PUT("/bookings/:nik/cancel", api.UpdateCanceledBooking)
	}
}
