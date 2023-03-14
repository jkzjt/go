package main

import "fmt"

func nope() { // block scope starts
	// hello and name are only visible here
	const name = "World"
	var hello = "Hello"

	_ = hello
} // block scope ends

func main() { // block scope starts
	// hello and name are not visible here

	// ERROR:
	// fmt.Println(hello, name)
} // block scope ends
