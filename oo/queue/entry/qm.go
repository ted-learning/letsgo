package main

import (
	"fmt"
	"oo/queue"
)

func main() {
	q := queue.Queue{1, 2, 3}
	q.Push(4)
	q.Push(5)
	fmt.Println(q.Pop())
	fmt.Println(q.Pop())
	fmt.Println(q.IsEmpty())
	fmt.Println(q.Pop())
	fmt.Println(q.Pop())
	fmt.Println(q.Pop())
	fmt.Println(q.IsEmpty())
}
