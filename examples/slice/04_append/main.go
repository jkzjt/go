package main

import . "fmt"

func main() {
	// func append(slice []Type, elems ...Type) []Type
	// https://pkg.go.dev/builtin@latest#append
	
	a := []byte{'H','e','l','l', 'o'}
	Println(string(a))
	
	a = append(a, ' ', 'H', 'i')
	Println(string(a))
	
	a = append(a, " Hey"...)
	Println(string(a))
	
	s := []byte{' ', 'W', 'o', 'r', 'l', 'd'}
	a = append(a, s...)
	Println(string(a))
}