package main

import "fmt"

func main() {
	// `byte` is an alias for uint8
	// 1 byte = 8 bits
	
	var b byte
	fmt.Printf("%08b = %[1]d\n", b)
	b = 255
	fmt.Printf("%08b = %[1]d\n", b)
	
	// BUILD FAIL
	// reason? overflow
	// b = 256
	// fmt.Printf("%08b = %[1]d\n", b)
}