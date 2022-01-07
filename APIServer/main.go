package main

import (
	"github.com/dev-beom/xaas/apiserver/domains/instance"
	ipc "github.com/james-barrow/golang-ipc"
	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"
	"strconv"
	"time"
)

func main() {
	err := godotenv.Load(".env")
	port := 5000
	if err == nil {
		port, _ = strconv.Atoi(os.Getenv("PORT"))
	}
	for {

	instanceRepository := instance.NewRepository()
	instanceService := instance.NewService(instanceRepository)
	instanceController := instance.NewController(instanceService)
	app := echo.New()

	app.Use(middleware.Logger())
	app.Use(middleware.Recover())

	app.GET("/api/instance/:id", instanceController.Get)
	app.GET("/api/instances", instanceController.GetAll)
	app.POST("/api/instance", instanceController.Create)
	app.DELETE("/api/instance/:id", instanceController.Delete)

	app.Logger.Fatal(app.Start(":" + strconv.Itoa(port)))
	_ = ipcServer.Write(1, []byte("Message from server"))
	_ = ipcServer.Write(1, []byte("Message from server"))
	_ = ipcServer.Write(1, []byte("Message from server"))
}
