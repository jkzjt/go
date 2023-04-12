package main

import (
	"fmt"
	"os"
	"strings"
	"time"
)

const (
	nleap = iota + 28
	leap
	small
	large
)

const (
	jan = "january"
	feb = "february"
	mar = "march"
	apr = "april"
	may = "may"
	jun = "june"
	jul = "july"
	aug = "august"
	sep = "september"
	oct = "october"
	nov = "november"
	dec = "december"
)

const (
	tip    = "Give me a month"
	msg    = "%q has %d days\n"
	errMsg = "%q is not a month\n"
)

const m4, m100, m400 = 4, 100, 400

func main() {
	if len(os.Args) == 1 {
		fmt.Println(tip)
		return
	}
	s := strings.TrimSpace(os.Args[1])
	m := strings.ToLower(s)
	if m == jan || m == mar || m == may || m == jul || m == aug || m == oct || m == dec {
		fmt.Printf(msg, s, large)
	} else if m == apr || m == jun || m == sep || m == nov {
		fmt.Printf(msg, s, small)
	} else if m == feb {
		y := time.Now().Year()
		if y%m4 == 0 && (y%m100 != 0 || y%m400 == 0) {
			fmt.Printf(msg, s, leap)
		} else {
			fmt.Printf(msg, s, nleap)
		}
	} else {
		fmt.Printf(errMsg, s)
	}
}
