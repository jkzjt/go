package main

import . "fmt"

const (
	pass = iota*10 + 60
	well
	excellent
	best
)

const (
	PASS      = "PASS"
	WELL      = "WELL"
	EXCELLENT = "EXCELLENT"
	BEST      = "BEST"
)

func main() {
	var (
		mark   = 88
		remark string
	)

	if mark >= pass {
		remark = PASS
	} else if mark >= well {
		remark = WELL
	} else if mark >= excellent {
		remark = EXCELLENT
	} else if mark >= best {
		remark = BEST
	}

	// expect `88 -- excellent` but `88 -- PASS`
	Println(mark, "--", remark)
	
	// --------------------------------------------
	// an if statement has more than one branches(elseif, else)
	// only one of them can be executed provided meet
	// the rest are ignored in spite of they are satisfied
	// --------------------------------------------
}
