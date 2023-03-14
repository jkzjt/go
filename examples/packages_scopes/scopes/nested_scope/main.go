package main

// file scope
import "fmt"

// package scope
var declareMeAgain = 10

// package scope
func nested() { // block scope starts
	// declares the same variable
	// they both can exist together
	// this one only belongs to this scope
	// package's variable is still intact

	var declareMeAgain = 5
	fmt.Println("inside nested: ", declareMeAgain)
} // block scope ends

// package scope
func main() { // block scope starts
	
	fmt.Println("inside main: ", declareMeAgain)
	
	nested()
	
	// package-level declareMeAgain isn't effected
	// from the change inside the nested func
	fmt.Println("inside main: ", declareMeAgain)
	
} // block scope ends