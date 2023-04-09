package main

import . "fmt"

const (
	Spring = iota*3 + 3
	Summer
	Fall
	Winter
)

func main() {
	Println(
		Spring, Summer, Fall, Winter,
	)
}
