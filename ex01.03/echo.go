// Copyright © 2016 Alan A. A. Donovan & Brian W. Kernighan.
// License: https://creativecommons.org/licenses/by-nc-sa/4.0/

// See page 4.
//!+

// Echo1 prints its command-line arguments.
package echo

import (
	// "fmt"
	"strings"
)

func Echo1(args []string) {
	var s, sep string
	for i := 1; i < len(args); i++ {
		s += sep + args[i]
		sep = " "
	}
	// fmt.Println(s)
}

func Echo2(args []string) {
	s, sep := "", ""
	for _, arg := range args {
		s += sep + arg
		sep = " "
	}
	// fmt.Println(s)
}

func Echo3(args []string) {
	// fmt.Println(strings.Join(args, " "))
	strings.Join(args, " ")
}

//!-
