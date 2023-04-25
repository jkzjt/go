package main

import (
	. "fmt"
	"unsafe"
)

func main() {
	// ------------------------------
	// Two-dimensional array(2D array)
	// -------------------------------
	// var identifier [len][len]type
	var a [2][3]string
	// type
	Printf("a's type: %T\n", a)
	// size
	Printf("a's size: %d\n", unsafe.Sizeof(a))
	// contiguous
	Printf("    a[i]: %p %p %p\n", &a, &a[0], &a[1])
	Printf(" a[i][j]: %p %p %p %p %p %p\n",
		&a[0][0], &a[0][1], &a[0][2],
		&a[1][0], &a[1][1], &a[1][2],
	)
	// zero values
	Printf("       a: %#v\n", a)

	// -------------------------
	// GET AND SET
	// ------------------
	Printf(" a[1][2]: %q\n", a[1][2])

	a[1][2] = `END`
	Printf(" a[1][2]: %q\n", a[1][2])

	// ---------------------------
	// Literal
	// --------------
	// var identifier [len][len]type = [len][len]type{...}
	var b [3][3]bool = [3][3]bool{
		[3]bool{true, false, true},
		[3]bool{false, true, false},
		[3]bool{true, false, true},
	}

	// var identifier = [len][len]type{...}
	var f = [2][2]float32{
		[2]float32{1, 2},
		[2]float32{10, 20},
	}

	// identifier := [len][len]type{...}
	i := [3][3]byte{
		{0, 2, 3},
		{4, 0, 6},
		{7, 8, 0},
	}

	// ellipsis `...`

	// build fail: invalid use of [...] array (outside a composite literal)
	// u := [3][...]uint8{{0, 1, 3},{5, 7},{9}}

	u := [...][3]uint8{
		{0, 1, 3}, 
		{5, 7}, 
		{9}, 
	}
	Printf("       u: %#v\n", u)

	_, _, _ = b, f, i
	
	// -------------------
	// COMPARE AND ASSIGNMENT
	// ------------------
	n := [...][2]int8{
		{3, 5}, 
		{4, 8},
	}
	
	n1 := [...][2]int8{
		{3, 5}, 
		{4, 8},
	}
	
	n2 := [...][2]int8{
		{5, 4}, 
		{3, 8},
	}
	
	Println(" n == n1?", n == n1)
	Println(" n == n2?", n == n2)
	
	
	var (
		s = [...][2]string{
			{"a", "b"},
			{"i", "j"},
		}
		s1 = s
	)
	Println(" s == s1?", s == s1)
	Printf("       s: %p\n", &s)
	Printf("      s1: %p\n", &s1)
	
	// `s1` is a new 2d array
}
