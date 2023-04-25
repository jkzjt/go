package main

import . "fmt"

func main() {
	a := [5]int{1, 2, 3, 4, 5}
	b := a[:3]
	c := a[2:]
	d := a[1:4]

	Printf("len(b): %d cap(b): %d\n", len(b), cap(b))
	Printf("len(c): %d cap(c): %d\n", len(c), cap(c))
	Printf("len(d): %d cap(d): %d\n", len(d), cap(d))

	// b, c, d share the same backing store array
	// even if the same element has different indices
	Printf("a[0]==b[0]: %t\n", &a[0] == &b[0])
	Printf("a[2]==c[0]: %t\n", &a[2] == &c[0])
	Printf("a[3]==d[2]: %t\n", &a[3] == &d[2])
	Printf(
		"a[2]==b[2]==c[0]==d[1]: %t\n", 
		&a[2] == &b[2] && 
		&b[2] == &c[0] &&
		&c[0] == &d[1],
	)
}
