package main

import (
	"fmt"
	"sync"
)

// START OMIT
var initer sync.Once

func someInitFunction() {
	initer.Do(func() {
		fmt.Println("initialized")
	})
}

func main() {
	var wg sync.WaitGroup
	wg.Add(1)
	go func() {
		someInitFunction()
		wg.Done()
	}()
	someInitFunction()
	someInitFunction()
	wg.Wait()
}

// END OMIT
