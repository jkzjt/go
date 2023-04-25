package main

import (
	. "fmt"
	"unsafe"
)

func main() {
	// slice is a dynamically-sized, flexible view into an underlying array
	
	// var identifier []type
	var (
		i []uint8
		b []bool
		a []string
	)
	
	// type
	Printf("i's type: %T\n", i) // []uint8
	Printf("b's type: %T\n", b) // []bool
	Printf("a's type: %T\n", a) // []string
	
	// zero values
	Printf("i's zero-value: %#v\n", i) // nil
	Printf("b's zero-value: %#v\n", b) // nil
	Printf("a's zero-value: %#v\n", a) // nil
	
	// ptr
	Printf("i's ptr: %p\n", &i)
	Printf("b's ptr: %p\n", &b)
	Printf("a's ptr: %p\n", &a)
	
	// size
	// 24 ???
	Printf("i's size: %d\n", unsafe.Sizeof(i))
	Printf("b's size: %d\n", unsafe.Sizeof(b))
	Printf("a's size: %d\n", unsafe.Sizeof(a))
}