package hanlder

import (
	"encoding/json"
	"github.com/dev-beom/xaas/controllermanager/constants/IPCMessage"
	"github.com/dev-beom/xaas/controllermanager/models"
	ipc "github.com/james-barrow/golang-ipc"
)

func Handler(pipe *ipc.Client) {
	for {
		data, err := pipe.Read()

		switch data.MsgType {
		case IPCMessage.ERROR:
			ipcError(err)
		case IPCMessage.STATUS_CHANGE:
			ipcStatusChange(data.Status)
		case IPCMessage.CREATE:
			node := new(models.Node)
			err := json.Unmarshal(data.Data, &node)
			if err != nil {
				ipcError(err)
				continue
			}
			ipcCreate(node)
		case IPCMessage.UPDATE:
			node := new(models.Node)
			err := json.Unmarshal(data.Data, &node)
			if err != nil {
				ipcError(err)
				continue
			}
			ipcUpdate(node)
		}
	}
}
