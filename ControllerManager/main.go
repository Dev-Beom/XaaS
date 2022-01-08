package main

import (
	"encoding/json"
	"github.com/dev-beom/xaas/controllermanager/constants"
	"github.com/dev-beom/xaas/controllermanager/constants/IPCMessage"
	"github.com/dev-beom/xaas/controllermanager/models"
	ipc "github.com/james-barrow/golang-ipc"
	"log"
)

func main() {
	ipcClient, err := ipc.StartClient(constants.IPCName, nil)
	if err != nil {
		return
	}

	for {
		data, err := ipcClient.Read()

		if err != nil {
			break
		}

		if data.MsgType == IPCMessage.STATUS_CHANGE {
			log.Println("Status: " + data.Status)
		}

		if data.MsgType == IPCMessage.ERROR {
			log.Println("Error: " + err.Error())
		}

		if data.MsgType > 0 { // all message types above 0 have been recieved over the connection
			var node models.Node
			json.Unmarshal(data.Data, &node)
			log.Println(" Message type: ", data.MsgType)
			log.Println("Client received: ", node.Id)
		}
	}
}
