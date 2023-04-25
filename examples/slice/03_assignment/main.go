package main

import . "fmt"

func main() {
	a := []string{"H","e","l", "l", "o"}
	s := a
	Println("a == s?", eq(a, s))
	
	m := []string{"H","i"}
	s = m
	Println("s == m?", eq(s, m))
	
	Printf("%p %p %p\n", &a, &s, &m)
}

func eq(a, b []string) bool {
	loa, lob := len(a), len(b)
	eq := loa == lob
	if eq {
		for i := 0; i < loa; i++ {
			if a[i] != b[i] {
				eq = false
				break
			}
		}
	}
	return eq
}