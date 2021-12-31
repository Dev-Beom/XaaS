package main

import (
	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"
	"net/http"
	"strconv"
)

func main() {
	app := echo.New()
	port := 5000

	instanceRepository := instance.NewRepository()
	instanceService := instance.NewService(instanceRepository)
	instanceController := instance.NewController(instanceService)

	app.Use(middleware.Logger())
	app.Use(middleware.Recover())

	app.GET("/", func(c echo.Context) error {
		return c.String(http.StatusOK, "Hello, World!\n")
	})

	app.Logger.Fatal(app.Start(":" + strconv.Itoa(port)))
}
