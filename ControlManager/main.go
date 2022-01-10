package main

import (
	"github.com/dev-beom/xaas/controlmanager/constants"
	"github.com/dev-beom/xaas/controlmanager/hanlder"
	"github.com/dev-beom/xaas/controlmanager/validations"
	ipc "github.com/james-barrow/golang-ipc"
)

func main() {
	validations.Run()
	ipcClient, err := ipc.StartClient(constants.IPCName, nil)
	if err != nil {
		return
	}
	hanlder.Handler(ipcClient)
}
