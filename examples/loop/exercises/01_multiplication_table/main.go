package main

import . "fmt"

// -----------------------------------
// Make a multiplication table like below
//
//	     X     0     1     2     3     4
//       
//       0     0     0     0     0     0
//
//       1     0     1     2     3     4
//       
//       2     0     2     4     6     8
//
//       3     0     3     6     9    12
//  
//       4     0     4     8    12    16
//
// ---------------------------------------------------------------

const n = 6

func main() {
	// make the table header
	Printf("%5s", "X")
	for i := 0; i <= n; i++ {
		Printf("%5d", i)
	}
	Println() // a new line makes it clear
	for i := 0; i <= n; i++ {
		Printf("%5d", i)
		for j := 0; j <= n; j++ {
			Printf("%5d", i*j)
		}
		Println()
	}
}