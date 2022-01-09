package commands

import (
	"github.com/dev-beom/xaas/controlmanager/constants"
	"github.com/dev-beom/xaas/controlmanager/utils"
)

func BuildNodeDockerImage() {
	utils.Logging("이미지를 생성합니다.")
	utils.RunCommand("docker", "build", "-t", constants.NodeDockerImageName, constants.NodeDockerfilePath)
	utils.Logging("이미지 생성이 완료되었습니다.")
}

func RunNodeDockerImage(nodeName string) {
	env := "NODE_NAME=" + nodeName
	// todo -v option 을 통한 mount 처리
	utils.Logging(nodeName + "컨테이너를 실행합니다.")
	utils.RunCommand("docker", "run", "-d", "--name", nodeName, "-e", env, constants.NodeDockerImageName)
}

func DeleteNodeContainer(nodeName string) {
	// todo 다양한 docker rm option 을 활용한 제거, volume 관련
	utils.Logging(nodeName + "컨테이너를 제거합니다.")
	utils.RunCommand("docker", "rm", "-f", nodeName)
	utils.Logging(nodeName + "컨테이너를 제거가 완료되었습니다.")
}
