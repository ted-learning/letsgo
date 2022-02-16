package main

import (
	"fmt"
	"unicode/utf8"
)

func main() {
	s := "Yes我爱慕课网!"
	fmt.Println(len(s))
	for _, b := range []byte(s) {
		fmt.Printf("%X ", b)
	}
	fmt.Println()

	for i, ch := range s {
		fmt.Printf("(%d, %X) ", i, ch)
	}
	fmt.Println()
	fmt.Println(utf8.RuneCountInString(s))

	bytes := []byte(s)
	for len(bytes) > 0 {
		decodeRune, size := utf8.DecodeRune(bytes)
		bytes = bytes[size:]
		fmt.Printf("%c ", decodeRune)
	}
	fmt.Println()

	for i, r := range []rune(s) {
		fmt.Printf("(%d %c) ", i, r)
	}
	fmt.Println()
}
