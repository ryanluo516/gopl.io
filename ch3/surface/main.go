// Copyright © 2016 Alan A. A. Donovan & Brian W. Kernighan.
// License: https://creativecommons.org/licenses/by-nc-sa/4.0/

// See page 58.
//!+

// Surface computes an SVG rendering of a 3-D surface function.
package main

import (
	"errors"
	"fmt"
	"log"
	"math"
	"net/http"
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
	// Exercise 3.4
	handler := func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "image/svg+xml")

		fmt.Fprintf(w, "<svg xmlns='http://www.w3.org/2000/svg' "+
			"style='stroke: grey; fill: white; stroke-width: 0.7' "+
			"width='%d' height='%d'>", width, height)
		min, max := minmax()

		for i := 0; i < cells; i++ {
			for j := 0; j < cells; j++ {
				ax, ay, err := corner(i+1, j)
				if err != nil {
					continue
				}

				bx, by, err := corner(i, j)
				if err != nil {
					continue
				}

				cx, cy, err := corner(i, j+1)
				if err != nil {
					continue
				}

				dx, dy, err := corner(i+1, j+1)
				if err != nil {
					continue
				}

				fmt.Fprintf(w, "<polygon style='stroke: %s; fill: white' points='%g,%g %g,%g %g,%g %g,%g'/>\n",
					color(i, j, min, max), ax, ay, bx, by, cx, cy, dx, dy)
			}
		}
		fmt.Fprintf(w, "</svg>")
	}
	http.HandleFunc("/", handler)
	//!-http
	log.Fatal(http.ListenAndServe("localhost:8000", nil))
}

func corner(i, j int) (float64, float64, error) {
	// Find point (x,y) at corner of cell (i,j).
	x := xyrange * (float64(i)/cells - 0.5)
	y := xyrange * (float64(j)/cells - 0.5)

	// Compute surface height z.
	z := f(x, y)
	// Exercise 3.1
	if math.IsInf(z, 0) {
		return 0, 0, errors.New("infinite height")
	}

	// Project (x,y,z) isometrically onto 2-D SVG canvas (sx,sy).
	sx := width/2 + (x-y)*cos30*xyscale
	sy := height/2 + (x+y)*sin30*xyscale - z*zscale
	return sx, sy, nil
}

func f(x, y float64) float64 {
	r := math.Hypot(x, y) // distance from (0,0)
	return math.Sin(r) / r
}

// Exercise 3.2
func eggbox(x, y float64) float64 {
	return 0.2 * (math.Cos(x) + math.Cos(y))
}

// Exercise 3.3
func minmax() (float64, float64) {
	min, max := math.NaN(), math.NaN()
	for i := 0; i < cells+1; i++ {
		for j := 0; j < cells+1; j++ {
			x := xyrange * (float64(i)/cells - 0.5)
			y := xyrange * (float64(j)/cells - 0.5)

			z := f(x, y)

			if math.IsNaN(min) || z < min {
				min = z
			}

			if math.IsNaN(max) || z > max {
				max = z
			}
		}
	}

	return min, max
}

func color(i, j int, zmin, zmax float64) string {
	min, max := math.NaN(), math.NaN()
	for k := 0; k <= 1; k++ {
		for l := 0; l <= 1; l++ {
			x := xyrange * (float64(i+k)/cells - 0.5)
			y := xyrange * (float64(j+l)/cells - 0.5)

			z := f(x, y)

			if math.IsNaN(min) || z < min {
				min = z
			}

			if math.IsNaN(max) || z > max {
				max = z
			}
		}
	}

	color := ""
	if math.Abs(max) > math.Abs(min) {
		red := math.Exp(math.Abs(max)) / math.Exp(math.Abs(zmax)) * 255
		if red > 255 {
			red = 255
		}
		color = fmt.Sprintf("#%02x0000", int(red))
	} else {
		blue := math.Exp(math.Abs(min)) / math.Exp(math.Abs(zmin)) * 255
		if blue > 255 {
			blue = 255
		}
		color = fmt.Sprintf("#0000%02x", int(blue))
	}
	return color
}

//!-
