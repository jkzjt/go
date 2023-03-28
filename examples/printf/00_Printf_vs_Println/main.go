package main

import "fmt"

func main() {
	ops, success, fail := 2350, 543, 433

	fmt.Println(
		"ops:", ops, "- success", success, "/", fail,
	)

	fmt.Printf(
		"ops: %d - success %d / %d\n",
		ops, success, fail,
	)

	brand := "Apple"

	// print `brand` in quoted-form like ""
	fmt.Println("\"" + brand + "\"")

	fmt.Printf("%q\n", brand)

	// if `brand` is not a string type
	// how to achieve it
}
