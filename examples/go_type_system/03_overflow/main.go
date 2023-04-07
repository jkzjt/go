package main

import (
	"fmt"
	. "math"
)

func main() {
	// Go Compiler can check overflow at compile-time
	// BUILD FAIL overflows
	// fmt.Println(int8(127 + 1))

	// but not at runtime
	var (
		i8   int8   = 127
		ui8  uint8  = 255
		ui16 uint16 = 65535
	)

	i8++
	ui8++
	fmt.Println(i8)  // -128
	fmt.Println(ui8) // 0

	i8--
	ui8--
	fmt.Println(i8)  // 127
	fmt.Println(ui8) // 255

	// for integers, it will wrap around when overflows

	var f32 float32 = MaxFloat32
	fmt.Println(f32 * 1.1) // +Inf

	// for float-point numbers, it will be `+Inf` when overflows

	// when converted, it will be truncated to max value if overflows
	fmt.Println(uint8(ui16)) // 255
}
