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
	app := echo.New()
	port := 5000

	ipcServer, err := ipc.StartServer("api-server", nil)
	if err != nil {
		return
	}
	for {

		ipcServer.Write(5, []byte("Hello Client 1"))
		ipcServer.Write(7, []byte("Hello Client 2"))
		ipcServer.Write(9, []byte("Hello Client 3"))

		time.Sleep(time.Second)
	}

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
	_ = ipcServer.Write(1, []byte("Message from server"))
	_ = ipcServer.Write(1, []byte("Message from server"))
	_ = ipcServer.Write(1, []byte("Message from server"))
}
