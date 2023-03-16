package main

import "fmt"

func main() {
	// You will see that Gofmt tool will format your code automatically.
	// https://golang.org/cmd/gofmt/
	//
	// This is because, for Go, it doesn't matter whether
	// you use semicolons between the statements or not.
	//
	// It adds them automatically after all.

	fmt.Println("Hello");fmt.Println("World");fmt.Println("Gopher");
	
	/*
		after run `go fmt` command, it will be
		
		fmt.Println("Hello")
		fmt.Println("World")
		fmt.Println("Gopher")
		
	*/
	
}
