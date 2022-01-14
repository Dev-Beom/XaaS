package validations

import (
	"github.com/dev-beom/xaas/controlmanager/commands"
	"github.com/dev-beom/xaas/controlmanager/utils"
	"os"
)

func Run() {
	err := checkDockerRunning()
	if err != nil {
		utils.Logging(err.Error())
		os.Exit(0)
	}
	err = checkDockerInstalled()
	if err != nil {
		utils.Logging(err.Error())
		os.Exit(0)
	}
	err = checkDockerImage()
	if err != nil {
		utils.Logging(err.Error())
		commands.BuildNodeDockerImage()
	}
}
