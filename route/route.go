package route

import (
	"project/controller"

	"github.com/labstack/echo/v4"
)

func New() *echo.Echo {
	e := echo.New()

	e.POST("/register", controller.RegisterController)
	e.POST("/login", controller.LoginController)
	e.GET("/addBooks", controller.AddBookController)
	e.GET("/books/search", controller.BookSearchByTitle)
	return e
}
