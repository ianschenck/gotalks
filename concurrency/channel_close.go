package main

import (
	"fmt"
	"sync"
	"time"
)

var wg sync.WaitGroup

// START OMIT
func worker(i int, exit chan struct{}) {
	for {
		select {
		case <-exit:
			fmt.Printf("worker%d - exit\n", i)
			wg.Done()
			return
		case <-time.Tick(80 * time.Millisecond):
			fmt.Printf("worker%d - tick\n", i)
		}
	}
}

func main() {
	exit := make(chan struct{})
	for i := 0; i < 10; i++ {
		wg.Add(1)
		go worker(i, exit)
	}
	<-time.After(1 * time.Second)
	close(exit)
	wg.Wait()
}

// END OMIT
