package main

import (
	"fmt"
	"strconv"
)

func main() {
	// raw string
	fmt.Println(
		`Hi` + ", " + `how` + " " + `are` + " " + `you` + "?")

	fmt.Println(
		`Hi` + ", " + `how` + " " + `are` + " " + `you` + "?",
	)

	exp := "1 + 1 = "
	res := 1 + 1
	// work? NO
	// invalid op
	// mismatched types
	// fmt.Println(
	//	exp + res,
	// )

	// ------------------------------------------------
	// String concatenation only works with strings
	// ------------------------------------------------
	// Itoa means Integer to ASCII
	fmt.Println(
		exp + strconv.Itoa(res),
	)

	fmt.Println(
		strconv.FormatBool(true) + " or " + strconv.FormatBool(false) + `?`,
	)
}
