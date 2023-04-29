package main

import (
	. "fmt"
	"unsafe"
	"unicode/utf8"
)

// byte is an alias for uint8 and is equivalent to uint8 in all ways. It is
// used, by convention, to distinguish byte values from 8-bit unsigned
// integer values.

// type byte = uint8

// uint8 range [0, 256]
// a byte is 8-bit

func main() {
	var b byte = 'H'
	Printf("%q %[1]T %[1]v %08[1]b %d-bit\n", b, unsafe.Sizeof(b)*8)
	
	s  := "Hello"
	bs := []byte{72, 101, 108, 108, 111}
	Println(" `s`  as bytes:", []byte(s))
	Println("`bs` as string:", string(bs))
	
	// beginning with go 1.13 you can type: 0b0110_1000 instead
	Printf("%q as binary: %08[1]b\n", 0b0110_1000)
	
	// How to count runes in a byte slice?
	// https://pkg.go.dev/unicode/utf8@latest#RuneCount
	
	// Hello, 世界
	Println("len  :", len([]byte{0x48, 0x65, 0x6c, 0x6c, 0x6f, 0x20, 0x2c, 0xe4, 0xb8, 0x96, 0xe7, 0x95, 0x8c}))
	Println("count:", utf8.RuneCount([]byte{0x48, 0x65, 0x6c, 0x6c, 0x6f, 0x20, 0x2c, 0xe4, 0xb8, 0x96, 0xe7, 0x95, 0x8c}))
}