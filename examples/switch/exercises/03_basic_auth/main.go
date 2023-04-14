package main

import (
	"fmt"
	"os"
	"strings"
)

// ------------------------------------
// https://github.com/jkzjt/go/blob/master/examples/if/exercises/08_basic_auth/README.txt
// -----------------------------------------

const (
	username      = "jack"
	password      = "123456"
	usage         = "Usage: [username] [password]"
	noExisted     = "Access denied for %q\n"
	wrongPassword = "Wrong password for %q\n"
	authSuccess   = "Access granted to %q\n"
)

func main() {
	if len(os.Args) != 3 {
		fmt.Println(usage)
		return
	}
	u := strings.TrimSpace(os.Args[1])
	p := strings.TrimSpace(os.Args[2])

	switch {
	case u != username:
		fmt.Printf(noExisted, u)
	case p != password:
		fmt.Printf(wrongPassword, u)
	default:
		fmt.Printf(authSuccess, u)
	}
}
