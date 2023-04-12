package main

import . "fmt"

func main() {
	var n int

	if n := 10; true {
		Println(n) // 10
	}
	Println(n) // 0

	if n, m := 10, 20; true {
		Println(n) // 10
		_ = m
	}
	Println(n) // 0

	if n = 10; true {
		Println(n) // 10
	}
	Println(n) // 10
}
