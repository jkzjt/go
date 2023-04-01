package main

import "fmt"

// raw and nonraw string are strings

func main() {
	fmt.Println("Hello World!")
	fmt.Println(`Hello World!`)
	
	fmt.Println("<html>\n\t<body></body>\n</html>")
	fmt.Println(`<html>
	<body></body>
</html>`)

	fmt.Println("C:\\dev\\be\\lang\\rust")
	fmt.Println(`C:\dev\be\lang\rust`)
}