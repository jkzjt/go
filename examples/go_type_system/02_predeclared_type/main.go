package main

import (
	"fmt"
	"math"
	"unsafe"
)

func main() {
	// numeric types
	var (
		// ------------
		// int
		// ------------
		
		// signed integer
		// at least 32-bit, depending on the machine
		// where your programs run
		i   int
		
		// 8-bit, range: (-1<<7 ~ 1<<7 - 1)
		i8  int8
		// 16-bit, range: (-1<<15 ~ 1<<15 - 1)
		i16 int16
		// 32-bit, range: (-1<<31 ~ 1<<31 - 1)
		i32 int32
		// 64-bit, range: (-1<<63 ~ 1<<63 - 1)
		i64 int64
		
		// unsigned integer
		// at least 32-bit, depending on the machine
		// where your programs run
		ui   uint
		
		// 8-bit, range: (0 ~ 1<<8 - 1)
		ui8  uint8
		// 16-bit, range: (0 ~ 1<<16 - 1)
		ui16 uint16
		// 32-bit, range: (0 ~ 1<<32 - 1)
		ui32 uint32
		// 64-bit, range: (0 ~ 1<<64 - 1)
		ui64 uint64
		
		// floating-point numbers
		// the set of IEEE-754 32-bit
		f32 float32
		// the set of IEEE-754 64-bit
		f64 float64
	)
	
	// the zero value of all the numeric types above is `0`
	fmt.Println(
		i, i8, i16, i32, i64, ui, ui8, ui16, ui32, ui64, f32, f64,
	)
	
	// ----------
	// min and max
	// ----------
	fmt.Printf("min and max of int:     %d ~ %d\n", math.MinInt, math.MaxInt)
	fmt.Printf("min and max of int8:    %d ~ %d\n", math.MinInt8, math.MaxInt8)
	fmt.Printf("min and max of int16:   %d ~ %d\n", math.MinInt16, math.MaxInt16)
	fmt.Printf("min and max of int32:   %d ~ %d\n", math.MinInt32, math.MaxInt32)
	fmt.Printf("min and max of int64:   %d ~ %d\n", math.MinInt64, math.MaxInt64)
	
	// printing directly will overflow, only used for constant expression
	fmt.Printf("min and max of uint:     %d ~ %d\n", 0, uint64(math.MaxUint))
	
	fmt.Printf("min and max of uint8:    %d ~ %d\n", 0, math.MaxUint8)
	fmt.Printf("min and max of uint16:   %d ~ %d\n", 0, math.MaxUint16)
	fmt.Printf("min and max of uint32:   %d ~ %d\n", 0, math.MaxUint32)
	
	// printing directly will overflow, only used for constant expression
	fmt.Printf("min and max of uint64:   %d ~ %d\n", 0, uint64(math.MaxUint64))
	
	// e means zeros after the number(scientific notation)
	// 2e1 is 20, 2e2 is 200 ...
	fmt.Printf("min and max of float32:  %g ~ %g\n", math.SmallestNonzeroFloat32, math.MaxFloat32)
	fmt.Printf("min and max of float64:  %g ~ %g\n", math.SmallestNonzeroFloat64, math.MaxFloat64)
	
	// bool
	// the set of boolean values
	var b bool
	// the zero value is `false`
	fmt.Println(b)
	
	// string
	// the set of all strings of 8-bit bytes
	// string type is immutable
	var s string
	// the zero value is `""` -- empty string
	fmt.Printf("%q\n", s)
	
	// ------------------------------
	// TODO
	// complex64 complex128 uintptr
	// -----------------------------
	
	// memory cost
	fmt.Printf("int      : %v bytes\n", unsafe.Sizeof(i))
	fmt.Printf("int8     : %v bytes\n", unsafe.Sizeof(i8))
	fmt.Printf("int16    : %v bytes\n", unsafe.Sizeof(i16))
	fmt.Printf("int32    : %v bytes\n", unsafe.Sizeof(i32))
	fmt.Printf("int64    : %v bytes\n", unsafe.Sizeof(i64))
	
	fmt.Printf("uint     : %v bytes\n", unsafe.Sizeof(ui))
	fmt.Printf("uint8    : %v bytes\n", unsafe.Sizeof(ui8))
	fmt.Printf("uint16   : %v bytes\n", unsafe.Sizeof(ui16))
	fmt.Printf("uint32   : %v bytes\n", unsafe.Sizeof(ui32))
	fmt.Printf("uint64   : %v bytes\n", unsafe.Sizeof(ui64))
	
	fmt.Printf("float32  : %v bytes\n", unsafe.Sizeof(f32))
	fmt.Printf("float64  : %v bytes\n", unsafe.Sizeof(f64))
	
	fmt.Printf("bool     : %v bytes\n", unsafe.Sizeof(b))
	
	fmt.Printf("string   : %v bytes\n", unsafe.Sizeof(s))
}