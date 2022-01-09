package main

import (
	"fmt"
	"os"
	"time"
)

func main() {
	// Todo Node functions -> The file that's currently mounted.
	nodeName := os.Getenv("NODE_NAME")
	for _ = range time.Tick(time.Second * 10) {
		fmt.Println(nodeName, ":", time.Now())
	}
}
