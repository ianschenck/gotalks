package main

import (
	"fmt"
	"runtime"
	"sync"
)

var (
	foo     string
	bar     string
	fooLock sync.Mutex
	barLock sync.Mutex
)

// START OMIT
func main() {
	go func() {
		fooLock.Lock()
		defer fooLock.Unlock()
		foo = "foo"
	}()
	go func() {
		fooLock.Lock()
		defer fooLock.Unlock()
		barLock.Lock()
		defer barLock.Unlock()
		bar = foo
	}()
	runtime.Gosched()
	fmt.Println(bar)
}

// END OMIT
