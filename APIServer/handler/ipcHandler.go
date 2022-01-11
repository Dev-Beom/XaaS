package handler

import (
	"encoding/json"
	"github.com/dev-beom/xaas/apiserver/constants/IPCMessage"
	"github.com/dev-beom/xaas/apiserver/domains/node"
	"github.com/dev-beom/xaas/apiserver/models"
	ipc "github.com/james-barrow/golang-ipc"
	"log"
)

func ipcError(err error) {
	log.Println("Error: ", err.Error())
}

func IpcHandler(pipe *ipc.Server, repository node.Repository) {
	for {
		data, err := pipe.Read()
		switch data.MsgType {
		case IPCMessage.ERROR:
			ipcError(err)
		case IPCMessage.STATUS_CHANGE:
			log.Println("Status: ", data.Status)
		case IPCMessage.UPDATE:
			nodeModel := new(models.Node)
			err := json.Unmarshal(data.Data, &nodeModel)
			if err != nil {
				ipcError(err)
				continue
			}
			_, err = repository.Update(nodeModel.Id, *nodeModel)
			if err != nil {
				ipcError(err)
			}
		}
	}
}
