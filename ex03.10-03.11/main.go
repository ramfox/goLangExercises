// Copyright © 2016 Alan A. A. Donovan & Brian W. Kernighan.
// License: https://creativecommons.org/licenses/by-nc-sa/4.0/

// See page 73.

// Comma prints its argument numbers with a comma at each power of 1000.
//
// Example:
// 	$ go build gopl.io/ch3/comma
//	$ ./comma 1 12 123 1234 1234567890
// 	1
// 	12
// 	123
// 	1,234
// 	1,234,567,890
//
package main

import (
	"bytes"
	"fmt"
	"os"
	"strings"
)

func main() {
	for i := 1; i < len(os.Args); i++ {
		fmt.Printf("  %s\n", comma(os.Args[i]))
	}
}

//!+
// comma inserts commas in a non-negative decimal integer string.
func comma(s string) string {
	n := len(s)
	if n <= 3 {
		return s
	}
	var buf bytes.Buffer
	var decimals string
	if s[0] == byte('-') || s[0] == byte('+') {
		buf.WriteString(s[0:1])
		s = s[1:]
	}
	if d := strings.Index(s, "."); d != -1 {
		decimals = s[d:]
		s = s[:d]
	}
	m := len(s) % 3
	if m == 0 {
		m = 3
	}
	buf.WriteString(s[:m])
	s = s[m:]
	for i := 0; i < len(s); i += 3 {
		buf.WriteString("," + s[i:i+3])
	}
	buf.WriteString(decimals)
	return buf.String()
}

//!-
