package main

import . "fmt"

const (
	young = iota*10 + 10
	adulthood
	wiser
	older = 60
)

func main() {
	age := 25
	if age > older {
		Println("Getting older")
	} else if age > wiser {
		Println("Getting wiser")
	} else if age > adulthood {
		Println("Adulthood")
	} else if age > young {
		Println("Young blood")
	} else {
		Println("Booting up")
	}
}

