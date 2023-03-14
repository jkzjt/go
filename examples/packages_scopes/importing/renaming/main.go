package main

// file scope
import "fmt"
import f "fmt"

// package scope
func main() { // block scope starts 
	fmt.Println("I'm fmt")
	f.Println("I'm an alias of fmt")
} // block scope ends