package main

import . "fmt"

func main() {
	var mark = 98
	switch {
	case mark >= 90:
		Println("Best")
	case mark >= 80:
		Println("Excellent")
	case mark >= 70:
		Println("Well")
	case mark >= 60:
		Println("Pass")
	default:
		Println("Fail")
	}
	
	switch {
	case mark >= 90:
		Print("Best ")
		fallthrough
	case mark >= 80:
		Print("Excellent ")
		fallthrough
	case mark >= 70:
		Print("Well ")
		fallthrough
	case mark >= 60:
		Print("Pass ")
		fallthrough
	default:
		Println("Fail")
		// BUILD FAIL: cannot fallthrough final case in switch
		// fallthrough
	}
	// Best Excellent Well Pass Fail
}
