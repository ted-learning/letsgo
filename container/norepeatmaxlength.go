package main

import "fmt"

func noRepeatMaxlength(s string) int {
	lastOccurred := make(map[rune]int)
	start := 0
	maxLength := 0
	for i, c := range []rune(s) {
		if lastI, ok := lastOccurred[c]; ok && lastI >= start {
			start = lastI + 1
		}
		if i-start+1 > maxLength {
			maxLength = i - start + 1
		}
		lastOccurred[c] = i
	}
	return maxLength
}

func main() {
	fmt.Println(noRepeatMaxlength("abc"))
	fmt.Println(noRepeatMaxlength("bbbbbb"))
	fmt.Println(noRepeatMaxlength("pwrfkekpw"))
	fmt.Println(noRepeatMaxlength("pwwkew"))
	fmt.Println(noRepeatMaxlength(""))
	fmt.Println(noRepeatMaxlength("b"))
	fmt.Println(noRepeatMaxlength("abcdefc"))
	fmt.Println(noRepeatMaxlength("一二一"))
}
