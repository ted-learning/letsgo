package main

import (
	"fmt"
	"math"
)

const name = "name"

func consts() {
	const (
		c = 1
		d = "d"
	)

	const a, b, f = 3, 4, false
	result := int(math.Round(math.Sqrt(a*a + b*b)))
	fmt.Println(name)
	fmt.Println(result, f)
	fmt.Println(c, d)
}

func enums() {
	const (
		a = iota
		_
		c
		d
		e
	)
	fmt.Println(a, c, d, e)
}

func bytes() {
	const (
		b = 1 << (iota * 10)
		kb
		mb
		gb
		tb
		pb
	)
	fmt.Println(b, kb, mb, gb, tb, pb)
}

func main() {
	consts()
	enums()
	bytes()
}
