package main

import (
	"fmt"
	"reflect"
	"unsafe"
)

func main() {
	s := "Hello, 世界!"
	
	sh := (*reflect.StringHeader)(unsafe.Pointer(&s))
	
	fmt.Printf("&s      : %p\n", &s)
	fmt.Printf("&sh.Data: %p\n", unsafe.Pointer(sh.Data))
	fmt.Printf("sh      : %#v\n", sh)	
	
	
	// Why the size of a string is always 16 bytes (based on 64-bit machine)
	// In the source code file runtime/string.go, there is a struct stringStruct as follows:
	// 
	//  type stringStruct struct {
	//    str unsafe.Pointer
	//    len int
	//  }
	// 
	// And the unsafe.Pointer is under unsafe package
	// 
	//  type Pointer *ArbitraryType
	// 
	//  type ArbitraryType int
	
	// The StringHeader in the package reflect is also a struct
	//
	// type StringHeader struct {
	//   Data uintptr
	//   Len  int
    // }
	
	// The underlying type of a string is an array of byte
	// The immutability of a string makes it illegal to take address of each element in the bucket

}