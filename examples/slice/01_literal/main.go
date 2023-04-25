package main

import (
	. "fmt"
	"unsafe"
)

func main() {
	// literal
	
	// var identifier []type = []type{...}
	var a []byte = []byte{'H', 'i'}
	
	// var identifier = []type{...}
	var b = []bool{true, false}
	
	// identifier := []type{...}
	i := []int8{1, 2, 3, 5, 8}
	
	// %#v -- print type and elements' value inside
	Printf("a: %#v\n", a)
	Printf("b: %#v\n", b)
	Printf("i: %#v\n", i)
	
	// size
	// 24 ???
	Printf("a's size: %d\n", unsafe.Sizeof(a))
	Printf("b's size: %d\n", unsafe.Sizeof(b))
	Printf("i's size: %d\n", unsafe.Sizeof(i))
	
	//  ptr
	Printf("a's ptr: %p\n", &a)
	Printf("b's ptr: %p\n", &b)
	Printf("i's ptr: %p\n", &i)
	
	// GET AND SET
	Println("a[0]:", a[0])
	a[0] = 'h'
	Println("a[0]:", a[0])
	
	// contiguous
	Printf("%p %p %p\n", &a, &a[0], &a[1])
	
}