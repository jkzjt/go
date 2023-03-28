package main

import "fmt"

func main(){
	// argument indexes start from 1
	// why?
	// 0 is used by the `formatter` of func Printf
	fmt.Printf(
		"%s is a developer. %[1]s is %d. He is a freshman of GO? %t\n",
		"Tan", 24, true,
	)
	
	// `%v` is powerful, why use other verbs than?
	// for type-safety
	// fmt.Printf(
	//	"%d is a developer. %[1]d is %d. He is a freshman of GO? %s\n",
	//	"Tan", 24, true,
	// )
}