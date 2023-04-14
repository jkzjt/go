package main

import . "fmt"

func main() {
	var city = "Berlin"
	if city == "Berlin" {
		Println("Germany")
	}
	if city == "Berlin" {
		Println("Germany")
	} else if city == "Tokyo" {
		Println("Japan")
	}
	if city == "Berlin" {
		Println("Germany")
	} else if city == "Tokyo" {
		Println("Japan")
	} else {
		Println("Unknow")
	}
}