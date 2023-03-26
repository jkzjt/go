package main

import "fmt"

func main() {
	i, j := 0, 9

	i, j = j, i

	fmt.Println(i, j)
}
