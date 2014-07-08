package main

import (
	"fmt"
)

func main() {
	buffered := make(chan struct{}, 42)
	buffered <- struct{}{}
	buffered <- struct{}{}
	fmt.Printf("cap(buffered) = %d and len(buffered) = %d\n", cap(buffered), len(buffered))
}
