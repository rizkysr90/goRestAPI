package route

import (
	"project/controller"
	m "project/middleware"

	"github.com/labstack/echo/v4"
)

func New() *echo.Echo {
	e := echo.New()

	e.POST("/register", controller.RegisterController)
	e.POST("/login", controller.LoginController)
	e.GET("/addBooks", controller.AddBookController)
	e.GET("/users/books/search", controller.SearchBookByTitle)
	m.LogMiddleware(e)
	e.POST("/:user_id/books/:id", controller.LoanBook)
	return e
}
