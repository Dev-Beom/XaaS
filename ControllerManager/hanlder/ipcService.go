package hanlder

import (
	"github.com/dev-beom/xaas/controllermanager/models"
	"log"
)

func ipcError(err error) {
	log.Println("Error: " + err.Error())
}

func ipcStatusChange(status string) {
	log.Println("Status: " + status)
}

func ipcCreate(node *models.Node) {

}

func ipcUpdate(node *models.Node) {

}
