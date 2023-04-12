package main

import (
	"fmt"
	"os"
	"strings"
)

const (
	tip     = `Usage: [username] [password]`
	u       = "sam"
	p       = "abc123"
	denied  = "Access denied for %q\n"
	wrong   = "Wrong password for %q\n"
	granted = "Access granted to %q\n"
)

func main() {
	if len(os.Args) != 3 {
		fmt.Println(tip)
		return
	}

	username := strings.TrimSpace(os.Args[1])
	password := strings.TrimSpace(os.Args[2])

	var msg string
	if username != u {
		msg = denied
	} else if password != p {
		msg = wrong
	} else {
		msg = granted
	}
	fmt.Printf(msg, username)
}
