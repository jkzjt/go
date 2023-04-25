package main

import . "fmt"

// -------------------------------------------
// func copy(dst, src []Type) int
//  The copy built-in function copies elements from a source slice into a
//  destination slice. (As a special case, it also will copy bytes from a
//  string to a slice of bytes.) The source and destination may overlap. Copy
//  returns the number of elements copied, which will be the minimum of
//  len(src) and len(dst).
// 
// ----------------------------------------------------------------

func main() {
	s := []int{1,1,2,3,5,8,13,21}
	d := make([]int, 4, 8)
	
	copy(d, s)
	Println("d:", d)
	Println("s:", s)
	Println()
	
	a := make([]byte, 16, 32)
	
	// invalid operation: invalid use of ... with built-in copy
	// copy(a, "Hello World!!!"...)
	
	copy(a, []byte("Hello World!!!"))
	Println("a:", string(a))
	Println()
	
	b := s[4:]
	copy(b, s)
	Println("b:", b)
	Println("s:", s)
}