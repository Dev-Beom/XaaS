package main

import (
	"fmt"
	"time"
)

func main() {
	// Todo Node functions -> The file that's currently mounted.
	for _ = range time.Tick(time.Second * 10) {
		fmt.Println(time.Now())
	}
}
