package main

import (
	. "fmt"
	"unsafe"
)

func main() {
	// An array is a piece of memory
	// ---------------------------------------------
	// var identifier [len]type
	// ---------------------------------------------
	var (
		i [2]int
		f [4]float64
		a [5]string 
		b [0]bool
	)
	// fixed-size
	Printf("i's len: %d, size(byte): %d\n", len(i), unsafe.Sizeof(i))
	Printf("f's len: %d, size(byte): %d\n", len(f), unsafe.Sizeof(f))
	Printf("a's len: %d, size(byte): %d\n", len(a), unsafe.Sizeof(a))
	Printf("b's len: %d, size(byte): %d\n", len(b), unsafe.Sizeof(b))
	
	// build fail: invalid array length -1 (untyped int constant)
	// var e [-1]byte
	
	// contiguous
	// An array's address is the address of first element of the array if any
	Printf("%p %p %p\n", &i, &i[0], &i[1])
	Printf("%p %p %p\n", &f, &f[0], &f[1])
	Printf("%p %p %p\n", &a, &a[0], &a[1])
	Printf("%p\n", &b)
	
	// consisted of length and type in it
	Printf("%T %T %T %T\n", i, f, a, b)
	
	// filled with zero-value
	Printf("i: %#v\n", i)
	Printf("f: %#v\n", f)
	Printf("a: %#v\n", a)
	Printf("b: %#v\n", b)
	
	// `%#v` output the array's type and element values in it
	
	// ---------------
	// GET AND SET
	// ------------------
	// index starts with 0 and up to it's length-1 if length is positive [0, len(a))
	Printf("i: %v %v\n", i[0], i[1])
	Printf("f: %v %v\n", f[0], f[1])
	Printf("a: %q %q\n", a[0], a[1])
	
	// build fail: invalid argument: index -1 (constant of type int) must not be negative
	// i[-1]
	
	// build fail: invalid argument: index 2 out of bounds [0:2]
	// i[2]
	
	i[0], i[1] = 4, 17
	Printf("i: %v %v\n", i[0], i[1])
}
