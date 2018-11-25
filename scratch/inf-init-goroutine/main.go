package main

import (
	"fmt"
	"time"
)

var t time.Time

func init() {
	go func() {
		for t.IsZero() {
			fmt.Println("t is still not set")
			time.Sleep(1 * time.Second)
		}

		fmt.Println("exiting init func")
	}()
}

func main() {
	fmt.Println("main has started")
	time.Sleep(1 * time.Second)
	t = time.Now()
	time.Sleep(2 * time.Second)
	fmt.Println("exiting main func")
}
