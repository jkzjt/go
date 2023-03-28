package main

import "fmt"

func main() {
	var (
		n = -1
		f = 5.5
		z = 0
	)
	
	// ---------------------------------------------
	// `+` -- always print a sign for numeric values
	// guarantee ASCII-only output for %q (%+q).
	// ---------------------------------------------
	// -1
	fmt.Printf("%+d\n", n)
	// +5.5
	fmt.Printf("%+g\n", f)
	// +0
	fmt.Printf("%+d\n", z)
	
	// 'Ā'
	fmt.Printf("%q\n", 256)
	// '\u0100'
	fmt.Printf("%+q\n", 256)
	
	// ----------------------------------------
	// `-` pad with spaces on the right rather than the left (left-­justify the field).
	// ???
	// ----------------------------------------
	
	// ----------------------------------------
	// `#` alternate format
	// ----------------------------------------
	// add leading 0 for octal (%#o)
	// 120
	fmt.Printf("%o\n", 0120)
	// 0120
	fmt.Printf("%#o\n", 0120)
	
	// add 0x for hex (%#x)
	// 1af
	fmt.Printf("%x\n", 0x1af)
	// 0x1af
	fmt.Printf("%#x\n", 0x1af)
	
	// add 0X for hex (%#X)
	// 226
	fmt.Printf("%X\n", 0x226)
	// 0X226
	fmt.Printf("%#X\n", 0x226)
	
	// for %q, print a raw (back-quoted) string if strconv.CanBackquote returns true
	// "HelloWorld"
	fmt.Printf("%q\n", "HelloWorld")
	// `HelloWorld`
	fmt.Printf("%#q\n", "HelloWorld")
	
	// suppress 0x for %p (%#p)
	// TODO
	
	// ----------------------------------------
	// `''` (space)
	// ----------------------------------------
	// leave a space for elided sign in numbers (% d)
	// `-1`
	fmt.Printf("% d\n", n)
	// ` 5.5`
	fmt.Printf("% g\n", f)
	// ` 0`
	fmt.Printf("% d\n", z)
	
	// put spaces between bytes printing strings or slices in hex (% x, % X)
	// 41 6d 65 6e
	fmt.Printf("% x\n", "Amen")
	// 41 6D 65 6E
	fmt.Printf("% X\n", "Amen")
	
	// --------------------------------------------
	// `0` pad with leading zeros rather than spaces
	// for numbers, this moves the padding after the sign
	// ???
	// --------------------------------------------
	
}