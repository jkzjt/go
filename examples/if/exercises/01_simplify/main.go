package main

import . "fmt"

func main() {
	// #1
	/*isSphere, radius := true, 200

	var big bool

	if radius >= 50 {
		if radius >= 100 {
			if radius >= 200 {
				big = true
			}
		}
	}

	if big != true {
		fmt.Println("I don't know.")
	} else if !(isSphere == false) {
		fmt.Println("It's a big sphere.")
	} else {
		fmt.Println("I don't know.")
	}*/
	
	// simplify the above code
	isSphere, radius := true, 200
	if isSphere && radius >= 200 {
		Println("It's a big sphere.")
	} else {
		Println("I don't know.")
	}
}