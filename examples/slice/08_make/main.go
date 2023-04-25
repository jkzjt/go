package main

import . "fmt"

// --------------------------------------------------
// func make(t Type, size ...IntegerType) Type
//  The make built-in function allocates and initializes an object of type slice
//  the first argument is a type
//  the return result is the same as the type
//
// Slice: The size specifies the length. The capacity of the slice is
//	equal to its length. A second integer argument may be provided to
//	specify a different capacity; it must be no smaller than the
//	length. For example, make([]int, 0, 10) allocates an underlying array
//	of size 10 and returns a slice of length 0 and capacity 10 that is
//	backed by this underlying array.
// --------------------------------------------------------------

func main() {
	s := make([]int8, 10)
	Printf("len: %d    cap: %d\n", len(s), cap(s))
	
	s = make([]int8, 5, 10)
	Printf("len: %d    cap: %d\n", len(s), cap(s))
	
	// invalid operation: make([]int8) expects 2 or 3 arguments; found 1
	// s = make([]int8)
	
	// invalid argument: length and capacity swapped
	// s = make([]int8, 5, 4)
}
