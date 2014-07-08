package main

import (
	"fmt"
)

var (
	a chan string = make(chan string)
	b chan string = make(chan string)
)

// START OMIT
func main() {
	go func() {
		a <- "foo"
	}()
	go func() {
		b <- (<-a)
	}()
	fmt.Println(<-b)
}

// END OMIT
