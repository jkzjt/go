package main

import "fmt"

func main() {
	fmt.Println("Hello")
	
	// You can access functions from other files
	// which are in the same package

	// Here, `main()` can access `bye()` and `hey()`

	// It's because bye.go, hey.go and main.go
	//  are in the main package.
	
	hey()
	bye()
}

// Note that: it's case-sensitive in GO
// GOOD 
func Bye() {
	fmt.Println("Bye")
}

// Note that: declare a fn is not related to parameter list
// ERROR: Bye redeclared in this block .\main.go:22:6: other declaration of Bye
/*func Bye(name string) {
	fmt.Println("Bye", name)
}*/

// Note that: declare a fn is not related to return  list
// ERROR: Bye redeclared in this block .\main.go:22:6: other declaration of Bye
/*func Bye() error {
	fmt.Println("Bye")
	return nil
}*/

// ERROE: bye redeclared in this block .\bye.go:5:6: other declaration of bye
/*func bye() {
	fmt.Println("Bye")
}*/


// ***** EXPLANATION *****
//
// ERROR: "bye" function "redeclared"
//        in "this block"
//
// "this block" means = "main package"
//
// "redeclared" means = you're using the same name
//   in the same scope again
//
// main package's scope is:
// all source-code files which are in the same main package