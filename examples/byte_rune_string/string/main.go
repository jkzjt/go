package main

import (
	"fmt"
	"unsafe"
	"unicode/utf8"
)

func main() {
	// https://go.dev/ref/spec#String_types
	
	// A string type represents the set of string values. 
	// A string value is a (possibly empty) sequence of bytes. 
	// The number of bytes is called the length of the string and is never negative. 
	// Strings are immutable: once created, it is impossible to change the contents of a string. 
	// The predeclared string type is `string`; it is a defined type. 
	
	// The length of a string `s` can be discovered using the built-in function len. 
	// The length is a compile-time constant if the string is a constant. 
	// A string's bytes can be accessed by integer indices 0 through len(s)-1. 
	// It is illegal to take the address of such an element; if s[i] is the i'th byte of a string, &s[i] is invalid. 
	
	// var emp string
	emp := ""
	s   := "Hello, World!"
	fmt.Println("len(emp):", len(emp))
	fmt.Println("len(s)  :", len(s))
	
	fmt.Printf("size(emp): %d bytes\n", unsafe.Sizeof(emp))
	fmt.Printf("size(s)  : %d bytes\n", unsafe.Sizeof(s))
	
	// IMMUTABLE
	// cannot assign
	// emp[0] = 'h'
	// s[0]   = 'h'
	
	// index access
	fmt.Println("s[0]:", s[0])
	// illegal: invalid operation: cannot take address of s[0] (value of type byte)
	// fmt.Println("&s[0]:", &s[0])
	
	// https://go.dev/ref/spec#String_literals
	
	// A string literal represents a string constant obtained from concatenating a sequence of characters. 
	// There are two forms: raw string literals and interpreted string literals. 
	
	// Raw string literals are character sequences between back quotes, as in `foo`. 
	// Within the quotes, any character may appear except back quote. 
	// The value of a raw string literal is the string composed of the 
	// uninterpreted (implicitly UTF-8-encoded) characters between the quotes; 
	// in particular, backslashes have no special meaning and the string may contain newlines. 
	// Carriage return characters ('\r') inside raw string literals are discarded from the raw string value. 
	
	// syntax error: unexpected literal `
	// raws := ```
	
	// newline ???
	raws := `\n` // not work
	fmt.Println("raws:", raws)
	raws = `U+000A` // not work
	fmt.Println("raws:", raws)
	raws = `
a` // work
	fmt.Println("raws:", raws)
	
	//  `\r`
	raws = `\r`
	fmt.Println("raws:", raws)
	
	// ------------------------------
	// backslashes have no special meaning in raw string literal
	// -------------------------------------
	
	// Interpreted string literals are character sequences between double quotes, as in "bar". 
	// Within the quotes, any character may appear except newline and unescaped double quote. 
	// The text between the quotes forms the value of the literal, with backslash escapes interpreted 
	// as they are in rune literals (except that \' is illegal and \" is legal), with the same restrictions. 
	// The three-digit octal (\nnn) and two-digit hexadecimal (\xnn) escapes represent individual bytes of the resulting string; 
	// all other escapes represent the (possibly multi-byte) UTF-8 encoding of individual characters. 
	// Thus inside a string literal \377 and \xFF represent a single byte of value 0xFF=255, 
	// while ÿ, \u00FF, \U000000FF and \xc3\xbf represent the two bytes 0xc3 0xbf of the UTF-8 encoding of character U+00FF. 
	
	// build fail: newline in string
	/*inters := "
"*/
	// syntax error: unexpected literal "
	// inters := """
	
	// unknown escape "\'" "\`"
	// inters := "\'"
	
	inters := "\""
	fmt.Println("inters:", inters)
	
	// "777" -- octal escape value 511 > 255
	inters = "\111"
	fmt.Println("inters:", inters)
	
	inters = "\x49"
	fmt.Println("inters:", inters)
	
	fmt.Println(`------------------------------------
Octal    hexadecimal    unicode
------------------------------------------------`)
	
	fmt.Printf("%-5q    %-5q    %-5q\n", "\110", "\x48", "\U00000048") // H
	fmt.Printf("%-5q    %-5q    %-5q\n", "\145", "\x65", "\U00000065") // e 
	fmt.Printf("%-5q    %-5q    %-5q\n", "\154", "\x6c", "\U0000006c") // l 
	fmt.Printf("%-5q    %-5q    %-5q\n", "\154", "\x6c", "\U0000006c") // l
	fmt.Printf("%-5q    %-5q    %-5q\n", "\157", "\x6f", "\U0000006f") // o 
	fmt.Printf("%-5q    %-5q    %-5q\n", "\054", "\x2c", "\U0000002c") // ,
	fmt.Printf("%-5q    %-5q    %-5q\n", "\040", "\x20", "\U00000020") // " "
	fmt.Printf("%-5q   %-5q   %-5q\n", "\344\270\226", "\xe4\xb8\x96", "\U00004e16") // 世
	fmt.Printf("%-5q   %-5q   %-5q\n", "\347\225\214", "\xe7\x95\x8c", "\U0000754c") // 界
	
	fmt.Println("\n")
	fmt.Println("\110\145\154\154\157\054\040\344\270\226\347\225\214")
	fmt.Println("\x48\x65\x6c\x6c\x6f\x2c\x20\xe4\xb8\x96\xe7\x95\x8c")
	fmt.Println("\U00000048\U00000065\U0000006c\U0000006c\U0000006f\U0000002c\U00000020\U00004e16\U0000754c")
	
	// How to count runes in a string
	// https://pkg.go.dev/unicode/utf8@latest#RuneCountInString
	str := "Здравствуй, мир!"
	fmt.Println("len(str)  :", len(str))
	fmt.Println("count(str):", utf8.RuneCountInString(str))
	
}