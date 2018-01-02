// Copyright © 2016 Alan A. A. Donovan & Brian W. Kernighan.
// License: https://creativecommons.org/licenses/by-nc-sa/4.0/

// See page 10.
//!+

// Dup2 prints the count and text of lines that appear more than once
// in the input.  It reads from stdin or from a list of named files.
package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

type tally struct {
	count     int
	fileNames []string
}

func (t *tally) String() string {
	return strings.Join(t.fileNames, ", ")
}

func (t *tally) Increment() int {
	t.count++
	return t.count
}

func (t *tally) AddFileName(f string) []string {
	var appears bool
	for i := 0; i < len(t.fileNames); i++ {
		if t.fileNames[i] == f {
			appears = true
			break
		}
	}
	if !appears {
		t.fileNames = append(t.fileNames, f)
	}
	return t.fileNames
}

func main() {

	counts := make(map[string]*tally)
	files := os.Args[1:]
	if len(files) == 0 {
		countLines(os.Stdin, counts)
	} else {
		for _, arg := range files {
			f, err := os.Open(arg)
			if err != nil {
				fmt.Fprintf(os.Stderr, "dup2: %v\n", err)
				continue
			}
			countLines(f, counts)
			f.Close()
		}
	}
	for line, n := range counts {
		if n.count > 1 {
			fmt.Printf("%d\t%s\t%s\n", n.count, line, n)
		}
	}
}

func countLines(f *os.File, counts map[string]*tally) {
	input := bufio.NewScanner(f)
	for input.Scan() {
		_, ok := counts[input.Text()]
		if !ok {
			counts[input.Text()] = &tally{}
		}
		counts[input.Text()].Increment()
		counts[input.Text()].AddFileName(f.Name())
	}
	// NOTE: ignoring potential errors from input.Err()
}

//!-
