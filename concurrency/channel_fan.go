package main

import (
	"fmt"
	"time"
)

// START OMIT
func fanIn(i int, c chan<- int) {
	for _ = range time.Tick(100 * time.Millisecond) {
		c <- i
	}
}

func fanOut(i int, c <-chan int) {
	for x := range c {
		fmt.Printf("fanout-%d got message from fanin-%d\n", i, x)
	}
}

func main() {
	c := make(chan int)
	for i := 0; i < 7; i++ {
		go fanIn(i, c) // HL
	}
	for i := 0; i < 5; i++ {
		go fanOut(i, c) // HL
	}
	<-time.After(1 * time.Second)
}

// END OMIT
