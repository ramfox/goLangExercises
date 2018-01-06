// Go Programming Language exercises
// 04.2
//
// Write a program that prints the SHA256 hash of its standard input by default
// but supports a command-line flag to print the SHA384 or SHA512 hash instead

package main

import (
	"crypto/sha256"
	"crypto/sha512"
	"flag"
	"fmt"
	"os"
)

var sha512Flag = flag.Bool("sha512", false, "use the sha512 hash")
var sha384Flag = flag.Bool("sha384", false, "use the sha384 hash")

// add flags --sha384 --sha512
func main() {
	flag.Parse()
	if len(os.Args) < 2 {
		fmt.Println("Provide an input to hash")
		return
	}
	switch {
	case *sha512Flag:
		fmt.Printf("%x\n", sha512.Sum512([]byte(os.Args[1])))
	case *sha384Flag:
		fmt.Printf("%x\n", sha512.Sum384([]byte(os.Args[1])))
	default:
		fmt.Printf("%x\n", sha256.Sum256([]byte(os.Args[1])))
	}

}
