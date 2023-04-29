package main

import (
	"mini_project/controller"
	"net/http"

	"github.com/labstack/echo/v4"
)

func main() {
	e := echo.New()
	e.GET("/", func(c echo.Context) error {
		return c.String(http.StatusOK, "Hello, World!")
	})
	e.POST("/signin", controller.Login)
	e.POST("/signup", controller.CreateUser)
	e.Logger.Fatal(e.Start(":1323"))
}
