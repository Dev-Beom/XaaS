package main

import (
	"github.com/dev-beom/xaas/controllermanager/constants"
	"github.com/dev-beom/xaas/controllermanager/hanlder"
	ipc "github.com/james-barrow/golang-ipc"
)

func main() {
	ipcClient, err := ipc.StartClient(constants.IPCName, nil)
	if err != nil {
		return
	}
	hanlder.Handler(ipcClient)
}
