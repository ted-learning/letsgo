package main

import "fmt"

var (
	aa = "aa"
	bb = "bb"
	cc = "cc"
)

func variableZeroValue() {
	var a int
	var s string
	fmt.Println(a, s)
}

func variableInitialValue() {
	var a = 1
	var s = "s"
	fmt.Printf("%d %q\n", a, s)
}

func variableTypeDeduction() {
	var a, s = 1, "s"
	fmt.Printf("%d %q\n", a, s)
}

func variableShorter() {
	a, s := 1, "s"
	fmt.Println(a, s)
}
func variableIn() {
	var (
		aa = "aa"
		bb = "bb"
		cc = "cc"
	)
	fmt.Println(aa, bb, cc)
}

func main() {
	variableZeroValue()
	variableInitialValue()
	variableTypeDeduction()
	variableShorter()
	variableIn()
	fmt.Println(aa, bb, cc)
}
