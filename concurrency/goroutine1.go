package main

import (
	"fmt"
	"math/rand"
	"sync"
	"time"
)

// START OMIT
func tick(msg string, count int, wg *sync.WaitGroup) {
	defer wg.Done()
	for i := 0; i < count; i++ {
		time.Sleep(time.Duration(100+rand.Intn(200)) * time.Millisecond)
		fmt.Printf("%s: %d\n", msg, i)
	}
}

func main() {
	var wg sync.WaitGroup
	wg.Add(2)
	go tick("foo", 5, &wg) // HL
	go tick("bar", 5, &wg) // HL
	fmt.Println("main routine waiting...")
	wg.Wait()
	fmt.Println("everything is done")
}

// END OMIT
