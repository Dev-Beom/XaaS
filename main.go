package main

import (
	APIServer "github.com/dev-beom/xaas/apiserver"
	ControlManager "github.com/dev-beom/xaas/controlmanager"
	"sync"
)

func main() {
	runner := new(sync.WaitGroup)
	runner.Add(2)
	go ControlManager.Run(runner)
	go APIServer.Run(runner)
	runner.Wait()
}
