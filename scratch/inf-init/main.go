package main

import (
	"fmt"
	"time"
)

var stop = time.Now().Add(5 * time.Second)

func init() {
	for !stop.Before(time.Now()) {
		time.Sleep(1 * time.Second)
	}

	fmt.Println("done initializing")
}

func main() {
	fmt.Println("main has started")
}
