package main

import (
	"fmt"
)

func main() {
	c := make(chan string, 1)
	c <- "foo"
	close(c)
	v, ok := <-c
	fmt.Printf("value='%s', ok=%v\n", v, ok)
	v, ok = <-c
	fmt.Printf("value='%s', ok=%v\n", v, ok)
}
