package main

import . "fmt"

type (
	// integer int

	n [3]int
	m [3]int
)

func main() {
	i := [3]int{1, 4, 5}
	i1 := [3]int{1, 4, 5}
	Println("i == i1?", i == i1) // true

	f := [3]float64{10., 40.}
	f1 := [3]float64{10., 40., 50.}
	Println("f == f1?", f == f1) // false f[2] != f1[2]

	a := [...]string{"a", "e", "i"}
	a1 := [...]string{"i", "e", "a"}
	Println("a == a1?", a == a1) // false elements inside are identical but orders not

	// named and unnamed
	p := n{1, 2, 3}
	p1 := m{1, 2, 3}
	p2 := [...]int{1, 2, 3}
	Println("p == p1?", m(p) == p1)
	Println("p == p2?", p == p2)
}
