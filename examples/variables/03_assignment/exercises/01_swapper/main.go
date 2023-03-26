package main

import "fmt"

func main() {
	good, bad := "green", "red"
	good, bad = bad, good
	fmt.Println(good, bad)
}
