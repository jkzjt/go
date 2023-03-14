package main

// file scope
import "fmt"

// package scope
func main() { // block scope starts
	fmt.Print("Hello")
	
	bye()
} // block scope ends