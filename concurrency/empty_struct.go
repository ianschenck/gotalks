package main

import (
	"fmt"
	"unsafe"
)

func main() {
	a := struct{}{}
	b := struct{}{}
	fmt.Printf("a and b have addresses %p and %p\n", &a, &b)
	fmt.Printf("sizeof(a) and sizeof(b) is %d\n", unsafe.Sizeof(a))
}
