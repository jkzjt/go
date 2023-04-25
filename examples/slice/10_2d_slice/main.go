package main

import (
	. "fmt"
	"unsafe"
)

func main() {
	a := [][]int{
		{1, 1, 2},
		{3, 5, 8},
		{13, 21, 34},
		{55, 89, 144},
	}
	
	
	
	
	
	
	
	Printf("len: %d cap: %d a: %v\n", len(a), cap(a), a)
	Printf("a's type: %T\n", a) // [][]int
	Printf("a's size: %d\n", unsafe.Sizeof(a)) // 24
	Println(&a[0][2], &a[1][0]) // contiguous
}