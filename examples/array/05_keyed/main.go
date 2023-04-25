package main

import . "fmt"

func main() {
	// build fail: index 2 is out of bounds (>= 2)
	// a := [2]string{"Domain", "Drive", "Design"}
	
	// unkeyed
	i := [...]int8{
		10,
		9,
		8,
	}
	Println(" i:", i)
	// keyed
	i1 := [...]int8{
		0: 10,
		1: 9,
		2: 8,
	}
	Println("i1:", i1)
	// keyed order
	i2 := [...]int8{
		1: 9,
		0: 10, 
		2: 8,
	}
	Println("i2:", i2)
	// auto initialized key
	// [...]int8{0,0,8}
	i3 := [...]int8{
		2: 8,
	}
	Println("i3:", i3)
	
	// build fail: invalid argument: index 2 out of bounds [0:2]
	// i4 := [2]int8{
	//	2: 8,
	// }
	// Println("i4:", i4)
	
	i5 := [4]int8{
		2: 8,
	}
	Println("i5:", i5)
	
	// index 4 is out of bounds (>= 4)
	// i6 := [4]int8{
	//	2: 8,
	//	9,
	//	10,
	// }
	// Println("i6:", i6)
	
	i7 := [4]int8{
		9, 
		2: 8,
	}
	Println("i7:", i7)
	
	// duplicate index 2 in array or slice literal
	// i8 := [4]int8{
	//	3,
	//	9, 
	//	6,
	//	2: 8,
	// }
	// Println("i8:", i8)
	
	// duplicate index 2 in array or slice literal
	// i9 := [4]int8{
	//	9, 
	//	2: 8,
	//	2: 10,
	// }
	// Println("i9:", i9)
	
	// keyed and unkeyed
	f := [...]float32{
		1, 
		2, 
		5: 10, 
		15, 
		20, 
		25: 100,
	}
	Println("f:", f)
	
}
