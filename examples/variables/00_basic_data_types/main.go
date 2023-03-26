package main

import (
	"fmt"
)

func main() {
	// integet literal
	fmt.Println(
		-200, -100, 0, 100, 500,
	)
	
	// float literal
	fmt.Println(
		-0.5, -100.5, .0, 6.5, 99.9, 100.,
	)
	
	// bool constants
	fmt.Println(
		true, false,
	)
	
	// string literal -- utf8
	fmt.Println(
		"", // empty - prints just a space
		"hi",

		// unicode
		"nasılsın?",   // "how are you" in turkish
		"hi there 星!", // "hi there star!"
	)
}