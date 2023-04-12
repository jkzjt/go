package main

import . "fmt"

func main() {
	{
		if true {
			Println(true)
		}
		if false {
			Println(false)
		}
	}
	{
		const factor = .75
		var (
			capacity = 32.
			size     = 31.
		)
		if capacity >= size+1 {
			size++
			Println("Add..")
		}
		if capacity < size+1 {
			capacity += capacity * factor
			Println("Resize...")
			size++
			Println("Add..")
		}
		Println("cap:", capacity, "size:", size)
	}
}
