package main

import (
	ipc "github.com/james-barrow/golang-ipc"
	"log"
)

func main() {
	ipcClient, err := ipc.StartClient("api-server", nil)
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

			log.Println(" Message type: ", data.MsgType)
			log.Println("Client recieved: " + string(data.Data))
		}
	}
}
