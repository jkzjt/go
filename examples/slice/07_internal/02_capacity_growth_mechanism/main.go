package main

import . "fmt"

func main() {
	s, oldcap := []int{1}, 1
	for cap(s) < 5e5 {
		s = append(s, 1)
		if cap(s) != oldcap {
			Printf(
				"Len: %8d    Oldcap: %8d    Newcap: %8d    R: %.2f\n",
				len(s), oldcap, cap(s), float64(cap(s))/float64(oldcap),
			)
			oldcap = cap(s)
		}
	}
}
