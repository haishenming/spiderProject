package main

import (
	"fmt"
	"unicode/utf8"
)

func main() {
	s := "这是一句中文"
	fmt.Printf("%X\n", []byte(s))
	fmt.Println(s, len(s))

	for i, ch := range s {  //ch is a rune
		fmt.Printf("(%d %X)", i, ch)
	}

	fmt.Println("Rune count:",
		utf8.RuneCountInString(s))

	bytes := []byte(s)

	for len(bytes) > 0 {
		ch, size := utf8.DecodeRune(bytes)
		bytes = bytes[size:]
		fmt.Printf("%c ", ch)

	}
	fmt.Println()

	for i, ch := range []rune(s) {
		fmt.Printf("(%d %c)", i, ch)
	}
	fmt.Println()


}
