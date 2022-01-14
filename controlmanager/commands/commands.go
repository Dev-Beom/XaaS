package commands

import (
	"github.com/dev-beom/xaas/controlmanager/constants"
	"github.com/dev-beom/xaas/controlmanager/utils"
)

func BuildNodeDockerImage() {
	utils.Logging("이미지 생성 작업이 시작되었습니다.")
	utils.RunCommand("docker", "build", "-t", constants.NodeDockerImageName, constants.NodeDockerfilePath)
	utils.Logging("이미지 생성 작업이 완료되었습니다.")
}

// RunNodeDockerImage todo -v option 을 통한 mount 처리
func RunNodeDockerImage(nodeName string) {
	env := "NODE_NAME=" + nodeName
	utils.Logging(nodeName + " 컨테이너 실행 작업이 시작되었습니다.")
	utils.RunCommand("docker", "run", "-d", "--name", nodeName, "-e", env, constants.NodeDockerImageName)
	utils.Logging(nodeName + " 컨테이너 실행 작업이 완료되었습니다.")
}

// DeleteNodeContainer todo 다양한 docker rm option 을 활용한 제거, volume 관련
func DeleteNodeContainer(nodeName string) {
	utils.Logging(nodeName + "컨테이너 제거 작업이 시작되었습니다.")
	utils.RunCommand("docker", "rm", "-f", nodeName)
	utils.Logging(nodeName + "컨테이너 제거 작업이 완료되었습니다.")
}
