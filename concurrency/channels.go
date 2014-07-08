package main

import (
	"fmt"
)

func main() {
	buffered := make(chan struct{}, 1) // HL
	unbuffered := make(chan struct{})  // HL
	go func() {
		unbuffered <- (<-buffered)
		fmt.Println("buffered had a value, moved it into unbuffered") // HL
	}()
	fmt.Println("shoving something into buffered") // HL
	buffered <- struct{}{}
	fmt.Println("shoved something into buffered") // HL
	<-unbuffered
	fmt.Println("got something out of unbuffered") // HL
}
