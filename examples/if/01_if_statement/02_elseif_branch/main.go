package main

import . "fmt"

func main() {
	// cannot exist alone
	// else if{
	// 	Println("else if")
	// }
	if true {
		Println("true if")
	} else if true {
		Println("true elseif")
	}
	if false {
		Println("false if")
	} else if true {
		Println("true elseif")
	}
}
