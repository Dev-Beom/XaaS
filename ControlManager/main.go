package main

import (
	"github.com/dev-beom/xaas/controlmanager/constants"
	"github.com/dev-beom/xaas/controlmanager/hanlder"
	ipc "github.com/james-barrow/golang-ipc"
)

func main() {
	// Todo ipc name -> ENV
	ipcClient, err := ipc.StartClient(constants.IPCName, nil)
	if err != nil {
		return
	}
	hanlder.Handler(ipcClient)
}
