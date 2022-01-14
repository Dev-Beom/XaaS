package controlmanager

import (
	"github.com/dev-beom/xaas/controlmanager/constants"
	"github.com/dev-beom/xaas/controlmanager/hanlder"
	"github.com/dev-beom/xaas/controlmanager/validations"
	ipc "github.com/james-barrow/golang-ipc"
	"sync"
)

func Run(runner *sync.WaitGroup) {
	defer runner.Done()

	validations.Run()
	ipcClient, err := ipc.StartClient(constants.IPCName, nil)
	if err != nil {
		return
	}
	hanlder.Handler(ipcClient)
}
