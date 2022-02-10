package main

import (
	"bufio"
	"fmt"
	"os"
)

func sum(max int) int {
	sum := 0
	for i := 0; i <= max; i++ {
		sum += i
	}
	return sum
}

func printFileLines(filename string) {
	if open, err := os.Open(filename); err != nil {
		panic(err)
	} else {
		scanner := bufio.NewScanner(open)
		for scanner.Scan() {
			fmt.Println(scanner.Text())
		}
	}
}

//goland:noinspection GoUnusedFunction
func deathLoop() {
	for {
		fmt.Println("no")
	}
}

func main() {
	fmt.Println(sum(100))

	printFileLines("a.txt")
	//deathLoop()
}
