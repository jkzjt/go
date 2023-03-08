package main

import (
	"embed"
)

//go:embed folder/single_file.txt
var fileStr string

//go:embed folder/single_file.txt
var fileBt []byte

//go:embed folder/single_file.txt
//go:embed folder/*.hash
var fd embed.FS

func main() {
	print(fileStr)
	print(string(fileBt))

	content1, _ := fd.ReadFile("folder/f1.hash")
	print(string(content1))

	content2, _ := fd.ReadFile("folder/f2.hash")
	print(string(content2))
}
