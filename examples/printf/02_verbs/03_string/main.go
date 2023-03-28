package main

import "fmt"

func main() {
	str := "Hello"
	
	// the uninterpreted bytes of the string or slice
	// Hello
	fmt.Printf("%s\n", str)
	
	// a double-quoted string safely escaped with Go syntax
	// "Hello"
	fmt.Printf("%q\n", str)
	
	// base 16, lower­case, two characters per byte
	// 48656c6c6f
	fmt.Printf("%x\n", str)
	
	// base 16, upper­case, two characters per byte
	// 48656C6C6F
	fmt.Printf("%X\n", str)
}