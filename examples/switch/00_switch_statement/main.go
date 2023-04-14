package main

import . "fmt"

func main() {
	var city = "Berlin"
	// one case if ... {...}
	switch city {
	case "Berlin":
		Println("Germany")
	}
	// more case if ... {...} else if ... {...}
	switch city {
	case "Berlin":
		Println("Germany")
	case "Tokyo":
		Println("Japan")
	}
	// default clause if ... {...} else if ... {...} else {...}
	switch city {
	case "Berlin":
		Println("Germany")
	case "Tokyo":
		Println("Japan")
	default:
		Println("Unknow")
	}
	// multiple conditions if ... || ... {...} else if ... || ... {...} else {...}
	switch city {
	case "Tokyo", "Singapore":
		Println("Asia")
	case "Berlin", "Hamburg":
		Println("Europe")
	default:
		Println("Where?")
	}
	
	var i = 1
	switch {
	case i > 0:
		Println("gt zero")
	case i < 0:
		Println("lt zero")
	default:
		Println("eq zero")
	}
}
