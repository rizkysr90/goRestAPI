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
	jwt := middleware.JWT([]byte(constants.JWT_SECRET))
	e.POST("user/register", controller.RegisterController)
	e.POST("user/login", controller.LoginController)
	e.GET("/addBooks", controller.AddBookController)
	e.GET("/search", controller.SearchBookByTitle, jwt)
	m.LogMiddleware(e)

	e.POST("/:user_id/books/:id", controller.LoanBook)
	return e
}
