package main

import "fmt"

func main(){
	n := 10
	
	// base 10
	// 10
	fmt.Printf("%d\n", n)
	
	// base 2
	// 1010
	fmt.Printf("%b\n", n)
	
	// base 8
	// 12
	fmt.Printf("%o\n", n)
	
	// base 16 with upper-case `A-F`
	// A
	fmt.Printf("%X\n", n)
	
	// base 16 with lower-case `a-f`
	// a
	fmt.Printf("%x\n", n)
	
	// the character repres­ented by the corresponding Unicode code point
	// \n
	fmt.Printf("%c\n", n)
	
	// a single­-quoted character literal safely escaped with Go syntax
	// '\n'
	fmt.Printf("%q\n", n)
	
	// Unicode format
	// U+000A
	fmt.Printf("%U\n", n)
	
}