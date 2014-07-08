package main

import (
	"fmt"
	"time"
)

// START OMIT
func WaitMany(a, b chan struct{}) {
	for a != nil || b != nil {
		select {
		case <-a:
			a = nil
		case <-b:
			b = nil
		}
	}
}

func main() {
	a, b := make(chan struct{}), make(chan struct{})
	t0 := time.Now()
	go func() {
		close(a)
		close(b)
	}()
	WaitMany(a, b)
	fmt.Printf("waited %v for WaitMany\n", time.Since(t0))
}

// END OMIT
