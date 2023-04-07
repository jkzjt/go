package main

import "fmt"

func main() {
	// alias types are the same types
	// different names for readability and refactor
	
	// type byte = uint8
	// type rune = int32
	// type any  = interface{}
	
	var (
		b byte = 10
		n uint8
		m int32 = 1024
		r rune
	)
	
	// allowed
	n = b
	r = m
	fmt.Printf(
		"b: %v, n: %v, m: %v, r: %v\n", 
		b, n, m, r,
	)
	
}