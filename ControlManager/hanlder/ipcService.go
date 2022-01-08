package hanlder

import (
	"github.com/dev-beom/xaas/controlmanager/models"
	"log"
)

func ipcError(err error) {
	log.Println("Error: " + err.Error())
}

func ipcStatusChange(status string) {
	log.Println("Status: " + status)
}

func ipcNodeCreate(node *models.Node) {
	// Todo node create logic
}

func ipcNodeUpdate(node *models.Node) {
	// Todo node update logic
}
