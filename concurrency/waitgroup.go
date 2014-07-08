package main

import (
	"fmt"
	"sync"
)

// START OMIT
func main() {
	var wg sync.WaitGroup
	wg.Add(2) // HL
	go func() {
		fmt.Println("I am goroutine #1")
		wg.Done() // HL
	}()
	go func() {
		fmt.Println("I am goroutine #2")
		wg.Done() // HL
	}()
	fmt.Println("Waiting...")
	wg.Wait() // HL
	fmt.Println("All done")
}

// END OMIT
