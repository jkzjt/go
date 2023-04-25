package main

import . "fmt"

func main() {
	i := [3]int{1, 2, 3}
	a := i

	Printf("i: %p %p %p\n", &i[0], &i[1], &i[2])
	Printf("a: %p %p %p\n", &a[0], &a[1], &a[2])
	
	Println("i == a?", i == a) // true
	
	// `a` is a new array
}
