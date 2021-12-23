package main

import (
	"fmt"
	"github.com/dev-beom/faas/single"
)

func main() {
	hello, err := single.Hello("모듈 환경 테스트")
	if err != nil {
		return
	}
	fmt.Println(hello)
}
