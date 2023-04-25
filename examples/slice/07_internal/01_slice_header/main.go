package main

import (
	"fmt"
	"reflect"
	"unsafe"
)

func main() {
	s := []int{2, 3, 5, 7, 11}
	
	// See https://pkg.go.dev/reflect@latest#SliceHeader
	// See https://pkg.go.dev/unsafe@latest#Pointer
	sh := (*reflect.SliceHeader)(unsafe.Pointer(&s))
	
	fmt.Printf("sh's Data: %v\n", sh.Data)
	fmt.Printf("sh's Len : %v\n", sh.Len) // 5
	fmt.Printf("sh's Cap : %v\n", sh.Cap) // 5
	
	fmt.Printf("s's array: %p\n", s)
}