package main

import (
	"fmt"
	"os"
	"strings"
)

const (
	tip = `
[command] [argument]

Available commands: lower, upper, title`
	unk = "Unknown command: %s\n"
	err = "More arguments: expect 1 but find %d\n"
)

func main() {
	args := os.Args
	if len(args) < 3 {
		fmt.Println(tip)
		return
	} else if len(args) > 3 {
		fmt.Printf(err, len(args)-2)
		return
	}
	var msg string
	switch strings.TrimSpace(args[1]) {
	case "lower":
		msg = strings.ToLower(args[2])
	case "upper":
		msg = strings.ToUpper(args[2])
	case "title":
		msg = strings.Title(args[2])
	default:
		msg = fmt.Sprintf(unk, args[1])
	}
	fmt.Println(msg)
}
