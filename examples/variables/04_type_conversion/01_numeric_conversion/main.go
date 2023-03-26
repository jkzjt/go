package main

import "fmt"

func main() {
	var (
		i  int
		i8 int8
		a int
		s string
	)

	// BUILD FAIL
	// `int` and `int8` are different types
	// i = i8

	// `int` and `int8` are allowed to convert each other
	// because both are numeric types
	i, i8 = int(i8), int8(i)
	
	i, i8 = 200, 100
	
	// up and down conversion are allowed
	// but latter may lead to unpredictable consequences
	// below `i`(int 200) is converted to int8 which
	// beyond it's range(-128~127)
	// overflow makes it negative 
	// i = 100, i8 = -56
	i, i8 = int(i8), int8(i)
	fmt.Println(i, i8)
	
	// we can convert an int to a string,but only int
	// 97 is a
	a = 97
	s = string(a)
	fmt.Println(s)
	
	// also works for byte slices
	// underlying type of byte is uint8, also an int
	fmt.Println(string([]byte{104, 105}))
	
}
