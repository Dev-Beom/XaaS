package validations

import (
	"errors"
	"github.com/dev-beom/xaas/controlmanager/constants"
	"github.com/dev-beom/xaas/controlmanager/utils"
	"strings"
)

func checkDockerInstalled() error {
	output := utils.RunCommand("docker", "version")
	ok := strings.Contains(output, "Version:")
	if ok == false {
		return errors.New("도커엔진이 설치되어있지 않습니다")
	}
	return nil
}

func checkDockerRunning() error {
	output := utils.RunCommand("docker")
	ok := strings.Contains(output, "Docker")
	if ok == false {
		return errors.New("도커엔진이 실행되고있지 않습니다")
	}
	return nil
}

func checkDockerImage() error {
	output := utils.RunCommand("docker", "images")
	ok := strings.Contains(output, constants.NodeDockerImageName)
	if ok == false {
		return errors.New("이미지가 존재하지 않습니다")
	}
	return nil
}
