// Copyright © 2016 Alan A. A. Donovan & Brian W. Kernighan.
// License: https://creativecommons.org/licenses/by-nc-sa/4.0/

// See page 58.
//!+

// Surface computes an SVG rendering of a 3-D surface function.
package main

import (
	"errors"
	"fmt"
	"math"
	"os"
)

const (
	width, height = 600, 320            // canvas size in pixels
	cells         = 100                 // number of grid cells
	xyrange       = 30.0                // axis ranges (-xyrange..+xyrange)
	xyscale       = width / 2 / xyrange // pixels per x or y unit
	zscale        = height * 0.4        // pixels per z unit
	angle         = math.Pi / 6         // angle of x, y axes (=30°)
)

var sin30, cos30 = math.Sin(angle), math.Cos(angle) // sin(30°), cos(30°)

func main() {
	if len(os.Args) > 1 {
		switch os.Args[1] {
		// case "eggcrate":
		// 	graph(eggcrate)
		// case "moguls":
		// 	graph(moguls)
		// case "saddle":
		// 	graph(saddle)
		default:
			graph(wave)
		}
	}
	graph(wave)
}

func graph(f func(i, j float64) (float64, error)) {
	fmt.Printf("<svg xmlns='http://www.w3.org/2000/svg' "+
		"style='stroke: grey; fill: white; stroke-width: 0.7' "+
		"width='%d' height='%d'>", width, height)
	for i := 0; i < cells; i++ {
		for j := 0; j < cells; j++ {
			ax, ay, err := corner(i+1, j, f)
			if err != nil {
				continue
			}
			bx, by, err := corner(i, j, f)
			if err != nil {
				continue
			}
			cx, cy, err := corner(i, j+1, f)
			if err != nil {
				continue
			}
			dx, dy, err := corner(i+1, j+1, f)
			if err != nil {
				continue
			}
			fmt.Printf("<polygon points='%g,%g %g,%g %g,%g %g,%g'/>\n",
				ax, ay, bx, by, cx, cy, dx, dy)
		}
	}
	fmt.Println("</svg>")
}

func corner(i, j int, f func(i, j float64) (float64, error)) (float64, float64, error) {
	// Find point (x,y) at corner of cell (i,j).
	x := xyrange * (float64(i)/cells - 0.5)
	y := xyrange * (float64(j)/cells - 0.5)

	// Compute surface height z.
	z, err := f(x, y)
	if err != nil {
		return x, y, err
	}

	// Project (x,y,z) isometrically onto 2-D SVG canvas (sx,sy).
	sx := width/2 + (x-y)*cos30*xyscale
	sy := height/2 + (x+y)*sin30*xyscale - z*zscale
	return sx, sy, nil
}

func wave(x, y float64) (float64, error) {
	r := math.Hypot(x, y) // distance from (0,0)
	r = math.Sin(r) / r
	if math.IsNaN(r) {
		return r, errors.New("Result is NaN")
	}
	if math.IsInf(r, 0) {
		return r, errors.New("Result overflows float64")
	}
	return r, nil
}

func eggcrate(x, y float64) (float64, error) {
	r := 4 * (math.Pow(math.Sin(x), 2) + math.Pow(math.Sin(y), 2))
	// r := (x * x) + (y * y) + 25*(math.Pow(math.Sin(x), 2)+math.Pow(math.Sin(y), 2))
	if math.IsNaN(r) {
		return r, errors.New("Result is NaN")
	}
	if math.IsInf(r, 0) {
		return r, errors.New("Result overflows float64")
	}
	return r, nil
}

//!-
