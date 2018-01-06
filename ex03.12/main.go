// anagram tells you if two strings are anagrams of each other
//
package main

import (
	"fmt"
	"os"
	"strings"
)

func main() {
	if len(os.Args) < 3 {
		fmt.Println("Need two strings to compare")
		return
	}
	if anagrams(os.Args[1], os.Args[2]) {
		fmt.Printf("Yes, the strings \"%s\" and \"%s\" ARE anagrams\n", os.Args[1], os.Args[2])
	} else {
		fmt.Printf("No, the strings \"%s\" and \"%s\" ARE NOT anagrams\n", os.Args[1], os.Args[2])
	}
}

//!+
// anagrams checks if two strings are anagrams
// thoughts:
// probs there is a quicker way to do this.
// replace creates a new string each time, and maybe don't need
// to do a contains before each. It just feels like having
// the contains there means we can bail early, and that
// is worth it I think
func anagrams(a, b string) bool {
	if len(a) != len(b) {
		return false
	}
	for i := 0; i < len(a); i++ {
		if !strings.Contains(b, a[i:i+1]) {
			return false
		}
		b = strings.Replace(b, a[i:i+1], "", 1)
	}
	return b == ""
}

//!-
