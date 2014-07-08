package main

import (
	"fmt"
)

// START OMIT
func receive(c chan string) (v string, ok bool) {
	select {
	case v, ok = <-c:
	default:
	}
	return
}

func main() {
	c := make(chan string, 1)
	c <- "foo"
	v, ok := receive(c)
	fmt.Printf("value='%s', ok=%v\n", v, ok)
	v, ok = receive(c)
	fmt.Printf("value='%s', ok=%v\n", v, ok)
}

// END OMIT
