package main

import (
	"fmt"
	"time"
)

func main() {
	tick1 := time.Tick(500 * time.Millisecond)
	tick2 := time.Tick(400 * time.Millisecond)
	closedChannel := make(chan struct{})
	close(closedChannel)
	select {
	case <-tick1:
		fmt.Println("tick1!")
	case <-tick2:
		fmt.Println("tick2")
	case <-closedChannel:
		fmt.Println("closedChannel")
	}
}
