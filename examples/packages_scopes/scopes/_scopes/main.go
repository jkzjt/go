package main

// file scope
import "fmt"

// package scope
const name = "World"

// package scope
func main() { // block scope starts
	var hello = "Hello"

	// hello and name are visible here
	fmt.Println(hello, name)
} // block scope ends
