package main

import (
	"fmt"
	"math"
)

const (
	width, height = 600, 320
	cells         = 100
	xyrange       = 30.0
	xyscale       = width / 2 / xyrange
	zscale        = height * 0.4
	angle         = math.Pi / 6
)

var sin30, cos30 = math.Sin(angle), math.Cos(angle)

func main() {
	fmt.Printf("<svg xmlns='http://www.w3.org/2000/svg' "+
		"style='stroke: grey; fill: white; stroke-width: 0.7' "+
		"width='%d' height='%d'>", width, height)
	for i := 0; i < cells; i++ {
		for j := 0; j < cells; j++ {
			ax, ay, aok := corner(i+1, j)
			if !aok {
				continue
			}
			bx, by, bok := corner(i, j)
			if !bok {
				continue
			}
			cx, cy, cok := corner(i, j+1)
			if !cok {
				continue
			}
			dx, dy, dok := corner(i+1, j+1)
			if !dok {
				continue
			}

			fmt.Printf("<polygon points='%g,%g %g,%g %g,%g %g,%g'/>\n",
				ax, ay, bx, by, cx, cy, dx, dy)
		}
	}
	fmt.Println("</svg>")
}

func corner(i, j int) (v1 float64, v2 float64, ok bool) {
	x := xyrange * (float64(i)/cells - 0.5)
	y := xyrange * (float64(j)/cells - 0.5)

	z, ok := f(x, y)
	if !ok {
		return
	}

	v1 = width/2 + (x-y)*cos30*xyscale
	v2 = height/2 + (x+y)*sin30*xyscale - z*zscale
	return
}

func f(x, y float64) (value float64, ok bool) {
	r := math.Hypot(x, y)
	s := math.Sin(r)
	if math.IsNaN(s) {
		return 0, false
	}
	return s, true
}
