package main

import (
	"encoding/json"
	ipc "github.com/james-barrow/golang-ipc"
	"log"
	"time"
)

type Node struct {
	Id          string    `json:"id"`
	Description string    `json:"description"`
	CreateAt    time.Time `json:"create_at"`
	UpdateAt    time.Time `json:"update_at"`
	State       string    `json:"state"`
}

func main() {
	ipcClient, err := ipc.StartClient("XaaS", nil)
	if err != nil {
		return
	}
	for {
		data, err := ipcClient.Read()

		if err != nil {
			break
		}

		if data.MsgType == -1 { // message type -1 is status change
			log.Println("Status: " + data.Status)
		}

		if data.MsgType == -2 { // message type -2 is an error, these won't automatically cause the recieve channel to close.
			log.Println("Error: " + err.Error())
		}

		if data.MsgType > 0 { // all message types above 0 have been recieved over the connection
			var node Node
			json.Unmarshal(data.Data, &node)
			log.Println(" Message type: ", data.MsgType)
			log.Println("Client received: ", node.Id)
		}
	}
}
