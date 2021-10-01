package route

import (
	constants "project/constant"
	"project/controller"
	m "project/middlewares"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

func New() *echo.Echo {
	e := echo.New()
	jwtUser := middleware.JWT([]byte(constants.JWT_SECRET_USER))
	jwtAdmin := middleware.JWT([]byte(constants.JWT_SECRET_ADMIN))

	e.POST("/user/register", controller.RegisterUserController)
	e.GET("/user/login", controller.LoginUserController)
	e.POST("/admin/register", controller.RegisterAdminController)
	e.GET("/admin/login", controller.LoginAdminController)
	e.POST("/addBooks", controller.AddBookController, jwtAdmin)
	e.GET("/search", controller.SearchBookByTitle, jwtUser)
	// e.GET("/admin/reservation", controller.GetReservation, jwtAdmin)
	e.POST("/admin/status", controller.AddStatusReservationCode)
	e.PUT("/admin/reservation", controller.ReservationProcces, jwtAdmin)
	m.LogMiddleware(e)

	e.POST("/reservation", controller.LoanBook, jwtUser)
	return e
}
