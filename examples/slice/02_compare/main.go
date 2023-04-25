package main

import . "fmt"

func main() {
	i  := []int8{1, 2, 3}
	i1 := []int8{1, 2, 3}
	
	// invalid operation: i == i1 (slice can only be compared to nil)
	// Println("i == i1?", i == i1)
	
	// how to compare slices
	l, l1, eq := len(i), len(i1), true
	if l == l1 {
		for k := 0; k < l; k++ {
			if i[k] != i1[k] {
				eq = false
				break
			}
		}
	} else {
		eq = false
	}
	Println("i == i1?", eq)
	
	/*a := []int8{10, 20, 30}
	b := []uint8{10, 20, 30}
	l, l1 = len(a), len(b)
	eq = l == l1
	if eq {
		for k := 0; k < l; k++ {
			// work ???
			// invalid operation: a[k] != b[k] (mismatched types int8 and uint8)
			if a[k] != b[k] {
				eq = false
				break
			}
		}
	}
	Println("a == b?", eq)*/
}