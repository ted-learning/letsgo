package main

import (
	"fmt"
	"reflect"
	"runtime"
)

func eval2(a, b int, op string) (result int, err error) {
	switch op {
	case "+":
		return a + b, nil
	case "*":
		return a * b, nil
	case "-":
		return a - b, nil
	case "/":
		q, _ := div(a, b)
		return q, nil
	default:
		return 0, fmt.Errorf("not supported op: %s", op)
	}
}

func div(a, b int) (q, r int) {
	return a / b, a % b
}

func apply(function func(a, b int) int, a, b int) int {
	functionRef := reflect.ValueOf(function)
	name := runtime.FuncForPC(functionRef.Pointer()).Name()
	fmt.Println("function name: ", name)
	return function(a, b)
}

func s(a ...int) int {
	rs := 0
	for i := range a {
		rs += a[i]
	}
	return rs
}

func add(a, b int) int {
	return a + b
}

func main() {
	if result, err := eval2(1, 2, "+"); err == nil {
		fmt.Println(result)
	} else {
		fmt.Println(err)
	}

	if result, err := eval2(1, 2, "["); err == nil {
		fmt.Println(result)
	} else {
		fmt.Println(err)
	}

	fmt.Println(s(1, 2, 3, 4, 5))

	a := apply(add, 1, 2)
	b := apply(func(a, b int) int {
		return a + b
	}, 1, 2)
	c := apply(func(a, b int) int {
		return a - b
	}, 1, 2)

	fmt.Println(a, b, c)
}
