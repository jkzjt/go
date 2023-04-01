package main

import "fmt"

func main(){
	n := 100
	
	// +=
	n += 100
	fmt.Println(n)
	
	// -=
	n -= 99
	fmt.Println(n)
	
	// *=
	n *= 2
	fmt.Println(n)
	
	// /=
	n /= 101
	fmt.Println(n)
	
	// %=
	n %= 2
	fmt.Println(n)
}