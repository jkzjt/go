package main

import . "fmt"

func main() {
	// -----------------------------------------------
	// func len(v Type) int
	// see https://pkg.go.dev/builtin@latest#len
	//
	// func cap(v Type) int
	// see https://pkg.go.dev/builtin@latest#cap
	// -----------------------------------------------

	n := []int{}

	Println("len    cap")
	Printf("%d    %d\n", len(n), cap(n))
	for i := 0; i < 10; i++ {
		n = append(n, i)
		Printf("%d    %d\n", len(n), cap(n))
	}
}
