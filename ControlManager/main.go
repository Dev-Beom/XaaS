package main

import (
	"github.com/dev-beom/xaas/controlmanager/commands"
	"github.com/dev-beom/xaas/controlmanager/constants"
	"github.com/dev-beom/xaas/controlmanager/hanlder"
	"github.com/dev-beom/xaas/controlmanager/utils"
	"github.com/dev-beom/xaas/controlmanager/validations"
	ipc "github.com/james-barrow/golang-ipc"
)

func main() {
	err := validations.CheckDockerInstalled()
	if err != nil {
		utils.Logging(err.Error())
		return
	}
	err = validations.CheckDockerImage()
	if err != nil {
		utils.Logging(err.Error())
		commands.BuildNodeDockerImage()
	}
	ipcClient, err := ipc.StartClient(constants.IPCName, nil)
	if err != nil {
		return
	}
	hanlder.Handler(ipcClient)
}
