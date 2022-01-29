package main

import (
	"fmt"
	"math"
	"math/cmplx"
)

func triangle() {
	a, b := 3, 4
	var c int
	c = int(math.Round(math.Sqrt(float64(a*a + b*b))))
	println(c)
}

func main() {
	fmt.Printf("%.3f\n", cmplx.Exp(1i*math.Pi)+1)
	triangle()
}
