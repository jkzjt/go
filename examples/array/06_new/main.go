package main

import . "fmt"

func main() {
	// new()
	// https://pkg.go.dev/builtin@latest#new
	
	// func new(Type) *Type
	// The new built-in function allocates memory. The first argument is a type, 
	// not a value, and the value returned is a pointer to a newly allocated zero 
	// value of that type. 
	
	p := new([3]int)
	Printf("  p's type: %T\n", p) // *[3]int
	Printf(" p's value: %v\n", *p) // [0 0 0]
	Printf("p's length: %v\n", len(p)) // 3
	Printf("         p: %p\n", p)
	Printf("      p[i]: %p %p %p\n", &p[0], &p[1], &p[2])
}