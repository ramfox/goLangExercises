// Copyright © 2016 Alan A. A. Donovan & Brian W. Kernighan.
// License: https://creativecommons.org/licenses/by-nc-sa/4.0/

// See page 43.
//!+

// Cf converts its numeric argument to Celsius and Fahrenheit.
package main

import (
	"fmt"
	"os"
	"strconv"

	"gopl.io/goLangExercises/ex02.01/tempconv"
)

func main() {
	for _, arg := range os.Args[1:] {
		t, err := strconv.ParseFloat(arg, 64)
		if err != nil {
			fmt.Fprintf(os.Stderr, "cf: %v\n", err)
			os.Exit(1)
		}
		f := tempconv.Fahrenheit(t)
		c := tempconv.Celsius(t)
		k := tempconv.Kelvin(t)
		fmt.Printf("%s = %s\t\t%s = %s\n%s = %s\t\t%s = %s\n%s = %s\t\t%s = %s\n",
			f, tempconv.FToC(f), f, tempconv.FToK(f), c, tempconv.CToF(c), c, tempconv.CToK(c), k, tempconv.KToC(k), k, tempconv.KToF(k))
	}
}

//!-
