package main

import . "fmt"

func main() {
	/*switch i := 10; true {
	case i > 0:
		Println("Positive")
	case i < 0:
		Println("Negative")
	default:
		Println("Zero")
	}*/

	switch i := 10; {
	case i > 0:
		Println("Positive")
	case i < 0:
		Println("Negative")
	default:
		Println("Zero")
	}
}
