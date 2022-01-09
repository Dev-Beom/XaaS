package main

import (
	"fmt"
	"time"
)

func main() {
	// Todo Node functions
	/**
	1. The file that's currently mounted.
	2. Log of commands.
	*/
	for _ = range time.Tick(time.Second * 10) {
		fmt.Println(time.Now())
	}
}
