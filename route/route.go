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
	e.POST("/user/login", controller.LoginUserController)
	e.POST("/admin/register", controller.RegisterAdminController)
	e.POST("/admin/login", controller.LoginAdminController)
	e.GET("/addBooks", controller.AddBookController, jwtAdmin)
	e.GET("/search", controller.SearchBookByTitle, jwtUser)
	m.LogMiddleware(e)

	e.POST("/reservation/:user_id/:book_id", controller.LoanBook, jwtUser)
	return e
}
