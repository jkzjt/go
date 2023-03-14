package main

// package scope
func bye() { // block scope starts
	// ERROR: undefined: fmt
	// This file cannot see main.go's imported names ("fmt").
	// Because the imported names belong to file scope.
	// fmt.Println(" World")
} // block scope ends