package main

import (
	"fmt"
	"math"
)

//C comment
func C(y, w, h float64) float64 {
	return w/2*math.Sqrt(1-y*y/(h*h)) + 0.5
}
func main() {
	r := 32.0
	p := r / 2.0
	q := p / 5.0
	s := ""
	for y := -p; y < p; y++ {
		for x := -r; x < r; x++ {
			d := C(y, r*2, p)
			e := C(y+q, r/5, q)
			f := e - p
			g := e + p
			h := C(y, r*1.3, r/3)
			if x >= d || x <= -d || (x > -g && x < f) || (x < g && x > -f) || (y > q && (x > -h && x < h)) {
				s += "0"
			} else {
				s += " "
			}
		}
		s += "\n"
	}
	fmt.Println(s)
}
