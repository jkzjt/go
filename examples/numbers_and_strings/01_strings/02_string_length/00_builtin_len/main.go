package main

import "fmt"

func main() {
	// string is made up of bytes, can be 1~4 bytes
	// the builtin function `len` counts
	// the bytes in string not characters
	fmt.Println(len("Garfield"))
	// `世界` is made up of 6 bytes
	fmt.Println(len("Hello, 世界"))
}
