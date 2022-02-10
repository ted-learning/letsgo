package main

import (
	"fmt"
	"io/ioutil"
)

func simpleReadFile() {
	const filename = "abc.txt"
	file, err := ioutil.ReadFile(filename)
	fmt.Println(file)
	fmt.Println(err)
}

func readFile() {
	const filename = "a.txt"
	if file, err := ioutil.ReadFile(filename); err == nil {
		fmt.Printf("%s\n", file)
		fmt.Printf(string(file))
	} else {
		fmt.Println(err)
	}
}

func eval(a, b int, op string) int {
	var result int
	switch op {
	case "+":
		result = a + b
	case "*":
		result = a * b
	case "-":
		result = a - b
	case "/":
		result = a / b
	default:
		panic("Not supported op:" + op)
	}
	return result
}

func main() {
	simpleReadFile()
	readFile()

	fmt.Println(eval(1, 2, "+"))
	fmt.Println(eval(112, 228, "+")) //340
	fmt.Println("5 - 3 =", eval(5, 3, "-"))
	fmt.Println("521 - 309 =", eval(521, 309, "-"))
	fmt.Println("2 * 3 =", eval(2, 3, "*"))
	fmt.Println("200 * 300 =", eval(200, 300, "*"))
}
