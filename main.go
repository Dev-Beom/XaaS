package main

import (
	"github.com/dev-beom/xaas/domains/instance"
	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"
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

	app.GET("/api/instance/:id", instanceController.Get)
	app.GET("/api/instances", instanceController.GetAll)
	app.POST("/api/instance", instanceController.Create)
	app.DELETE("/api/instance/:id", instanceController.Delete)

	app.Logger.Fatal(app.Start(":" + strconv.Itoa(port)))
}
