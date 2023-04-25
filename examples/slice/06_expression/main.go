package main

import . "fmt"

func main() {
	// `[n:m]` is a half-open interval, `n` included and `m` excluded
	// `n` and `m`' range: 0~capacity. if n > m, it will build fail(invalid slice indices: m < n)
	// provided both are provided explicitly or it will panic (slice bounds out of range)

	// simple expression
	a := []int{1, 2, 3, 4, 5, 6}
	Printf("%p %[1]v\n", a)

	Println()
	Printf("%p %[1]v\n", a[0:1])
	Printf("%p %[1]v\n", a[1:2])
	Printf("%p %[1]v\n", a[2:3])
	Printf("%p %[1]v\n", a[3:4])
	Printf("%p %[1]v\n", a[4:5])
	Printf("%p %[1]v\n", a[5:6])

	Println()
	Printf("%p %[1]v\n", a[0:0])
	Printf("%p %[1]v\n", a[1:1])
	Printf("%p %[1]v\n", a[2:2])
	Printf("%p %[1]v\n", a[3:3])
	Printf("%p %[1]v\n", a[4:4])
	Printf("%p %[1]v\n", a[5:5])
	Printf("%p %[1]v\n", a[6:6])

	// --------------
	// default indices
	// -------------------
	Println()
	// a[:] eq [0:len](a) {1, 2, 3, 4, 5, 6}
	Printf("%p %[1]v\n", a[:])

	// a[1:] eq [1:len] {2, 3, 4, 5, 6}
	Printf("%p %[1]v\n", a[1:])

	// a[:3] eq [0:3] {1, 2, 3}
	Printf("%p %[1]v\n", a[:3])

	Println()
	Println()

	// ---------------------------
	// slicing an array, it works? YES
	// ----------------------------------
	b := [5]int{1, 2, 3}
	Printf("%v\n", b)

	Println()
	Printf("%p %[1]v\n", b[2:])

	Println()
	Println()

	// full expression [l:h:m]
	c := []int{0, 1, 2, 3, 4, 5}
	s := c[1:3:5]
	// len: 3-1 cap: 5-1
	Printf("%d %d %v\n", len(s), cap(s), s)

	// middle and final index required in 3-index slice
	// s = c[::]
	s = c[:3:5]
	// len: 3-0 cap: 5-0
	Printf("%d %d %v\n", len(s), cap(s), s)
}
