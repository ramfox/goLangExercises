// Copyright © 2016 Alan A. A. Donovan & Brian W. Kernighan.
// License: https://creativecommons.org/licenses/by-nc-sa/4.0/

// Run with "web" command-line argument for web server.
// See page 13.
//!+main

// Lissajous generates GIF animations of random Lissajous figures.
package main

import (
	"fmt"
	"image"
	"image/color"
	"image/gif"
	"io"
	"math"
	"math/rand"
	"os"
	"strconv"
)

//!-main
// Packages not needed by version in book.
import (
	"log"
	"net/http"
	"net/url"
	"time"
)

//!+main

var palette = []color.Color{color.Black, color.RGBA{0x00, 0xFF, 0x00, 0xff}, color.RGBA{0xff, 0x00, 0xff, 0xff}}

const (
	whiteIndex = 0 // first color in palette
	blackIndex = 1 // next color in palette
)

func main() {
	//!-main
	// The sequence of images is deterministic unless we seed
	// the pseudo-random number generator using the current time.
	// Thanks to Randall McPherson for pointing out the omission.
	rand.Seed(time.Now().UTC().UnixNano())

	if len(os.Args) > 1 && os.Args[1] == "web" {
		//!+http

		handler := func(w http.ResponseWriter, r *http.Request) {
			var values url.Values = r.URL.Query()
			lissajous(w, values)
		}
		// paramHandler := func(w http.ResponseWriter, r *http.Request) {
		// }

		http.HandleFunc("/", handler)
		// http.HandleFunc("/?")
		//!-http
		log.Fatal(http.ListenAndServe("localhost:8000", nil))
		return
	}
	//!+main
	// lissajous(os.Stdout)
}

func lissajous(out io.Writer, values url.Values) {
	// default values
	var cycles = 5.0
	var res = 0.001
	var size = 100.0
	var nframes = 64
	var delay = 8
	// for each param given, set the value, if there is an error, stay with the default
	for key, val := range values {
		switch key {
		case "cycles":
			var c, err = strconv.ParseFloat(val[0], 64)
			if err != nil {
				fmt.Println(err)
				continue
			}
			cycles = c
		case "res":
			var r, err = strconv.ParseFloat(val[0], 64)
			if err != nil {
				fmt.Println(err)
				continue
			}
			res = r
		case "size":
			var s, err = strconv.ParseFloat(val[0], 64)
			if err != nil {
				fmt.Println(err)
				continue
			}
			size = s
		case "nframes":
			var n, err = strconv.Atoi(val[0])
			if err != nil {
				fmt.Println(err)
				continue
			}
			nframes = n
		case "delay":
			var d, err = strconv.Atoi(val[0])
			if err != nil {
				fmt.Println(err)
				continue
			}
			delay = d
		}
	}
	// do the liss math!
	freq := rand.Float64() * 3.0 // relative frequency of y oscillator
	anim := gif.GIF{LoopCount: nframes}
	phase := 0.0 // phase difference
	for i := 0; i < nframes; i++ {
		rect := image.Rect(0, 0, int(2*size+1), int(2*size+1))
		img := image.NewPaletted(rect, palette)
		for t := 0.0; t < cycles*2*math.Pi; t += res {
			x := math.Sin(t)
			y := math.Sin(t*freq + phase)
			img.SetColorIndex(int(size+(x*size+0.5)), int(size+(y*size+0.5)),
				blackIndex)
		}
		phase += 0.1
		anim.Delay = append(anim.Delay, delay)
		anim.Image = append(anim.Image, img)
	}
	gif.EncodeAll(out, &anim) // NOTE: ignoring encoding errors
}

//!-main
