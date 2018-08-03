package main

import (
	"fmt"
	"unicode/utf8"
)

func main() {
	s := "神奇大叶子" // UTF-8
	fmt.Println(len(s))
	for _, b := range []byte(s) {
		fmt.Printf("%X ", b) //  16进制
	}
	fmt.Println()

	for i, ch := range s {
		fmt.Printf("(%d %X)", i, ch)
	}
	fmt.Println()

	fmt.Println("Rune count:", utf8.RuneCountInString(s))

	bytes := []byte(s)
	for len(bytes) > 0 {
		ch, size := utf8.DecodeRune(bytes)
		bytes = bytes[size:]
		fmt.Printf("%c ", ch) // 字符
	}
	fmt.Println()

	for i, ch := range []rune(s) {
		fmt.Printf("(%d %c) ", i, ch)
	}
	fmt.Println()
}
