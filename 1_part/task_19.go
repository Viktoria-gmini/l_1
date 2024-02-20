package main

import (
	"fmt"
	"unicode/utf8"
)

// переворачиваем строку
func reverseString(s string) string {
	runes := []rune(s)
	for i, j := 0, utf8.RuneCountInString(s)-1; i < j; i, j = i+1, j-1 {
		runes[i], runes[j] = runes[j], runes[i]
	}
	return string(runes)
}

func main() {
	str := "главрыба"
	reversed := reverseString(str)
	fmt.Printf("%s -> %s\n", str, reversed)
}
