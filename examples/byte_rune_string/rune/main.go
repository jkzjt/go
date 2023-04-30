package main

import (
	. "fmt"
	"unsafe"
	"unicode/utf8"
)

// rune is an alias for int32 and is equivalent to int32 in all ways. It is
// used, by convention, to distinguish character values from integer values.

// type rune = int32

// range -2^31 ~ 2^31-1

func main() {
	// var r rune = '中'
	r := '中'
	Printf("%q %[1]T %[1]v %d-bit\n", r, unsafe.Sizeof(r)*8)
	
	// [72 101 108 108 111 44 19990 30028]
	rs := []rune{'H', 'e', 'l', 'l', 'o', ',', ' ', '世', '界'}
	Println(rs, string(rs))
	
	s := "Hello, 世界"
	Println([]rune(s), s)
	
	// invalid argument: r
	// Println("len(r):", len(r))
	
	Println("len(r):", len(string(r)))
	
	// any ways to do it -- unicode/utf8#RuneLen
	Println("len(r):", utf8.RuneLen(r))
	
	// cannot convert r (variable of type rune) to type []byte
	// Println("len(r):", len([]byte(r)))
}