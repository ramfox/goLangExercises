// Copyright © 2016 Alan A. A. Donovan & Brian W. Kernighan.
// License: https://creativecommons.org/licenses/by-nc-sa/4.0/

// See page 16.
//!+

// Fetch prints the content found at each specified URL.
package main

import (
	"fmt"
	"io"
	"net/http"
	"os"
	"strings"
)

func main() {
	for _, url := range os.Args[1:] {
		if !strings.HasPrefix(url, "http://") || !strings.HasPrefix(url, "https://") {
			url = "http://" + url
		}
		resp, err := http.Get(url)
		if err != nil {
			fmt.Fprintf(os.Stderr, "fetch: %v\n", err)
			os.Exit(1)
		}
		fmt.Println("Server Status: ", resp.Status)
		_, err = io.Copy(os.Stdout, resp.Body)
		resp.Body.Close()
		if err != nil {
			fmt.Fprintf(os.Stderr, "fetch: reading %s: %v\n", url, err)
			os.Exit(1)
		}
	}
}

//!-
// Interesting, io.Copy writes to os.Stdout, so you don't need to fmt.PrintF(%s, b). However, this means that you don't need both results from io.Copy (b is the number of bytes that were successfully copied to the stream). So I used a blank identifier to discard one of the outputs of io.Copy. But! I got this error when I tried to compile:
// ./main.go:26:10: no new variables on left side of :=
//  Which is super interesting! apparently you only need to use := when there is a new variable.
