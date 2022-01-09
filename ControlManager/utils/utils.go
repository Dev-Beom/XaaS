package utils

import (
	"io"
	"log"
	"os/exec"
	"time"
)

func Logging(message string) {
	// todo file logging
	log.Println(time.Now(), ":", message)
}

func RunCommand(name string, arg ...string) string {
	cmd := exec.Command(name, arg...)
	stdin, err := cmd.StdinPipe()
	if err != nil {
		log.Println(err)
	}

	go func() {
		defer func(stdin io.WriteCloser) {
			err := stdin.Close()
			if err != nil {
				return
			}
		}(stdin)
	}()

	out, err := cmd.CombinedOutput()
	if err != nil {
		log.Println(err)
	}
	return string(out)
}
