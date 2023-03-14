package main

import "fmt"

func main(){
	fmt.Println("Hello, World")
	fmt.Println("Hello, Gopher")
	fmt.Println("Hello, World", "Hello, Myself")
	// fmt.Println(Hello, World, Hello, Myself) // undefined
}
/* expected 'package', found 'func'

package main

import "fmt"

*/