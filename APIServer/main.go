package main

import (
	"github.com/dev-beom/xaas/apiserver/constants"
	"github.com/dev-beom/xaas/apiserver/domains/node"
	ipc "github.com/james-barrow/golang-ipc"
	"github.com/joho/godotenv"
	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"
	"os"
	"strconv"
)

func main() {
	err := godotenv.Load(".env")
	port := 5000
	if err == nil {
		port, _ = strconv.Atoi(os.Getenv("PORT"))
	}
	ipcServer, _ := ipc.StartServer(constants.IPCName, nil)
	app := echo.New()
	nodeRepository := node.NewRepository(ipcServer)
	nodeService := node.NewService(nodeRepository)
	nodeController := node.NewController(nodeService)

	app.Use(middleware.Logger())
	app.Use(middleware.Recover())

	app.GET("/api/node/:id", nodeController.Get)
	app.GET("/api/nodes", nodeController.GetAll)
	app.POST("/api/node", nodeController.Create)
	app.DELETE("/api/node/:id", nodeController.Delete)
	// todo file upload 기능
	app.Logger.Fatal(app.Start(":" + strconv.Itoa(port)))
}
