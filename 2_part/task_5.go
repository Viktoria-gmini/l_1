package main

import (
	"fmt"
	"unsafe"
)

// Какой размер у структуры struct{}{}?
// Ответ: 0 байт. Размер структуры определяется размером её полей
func main() {
	s := struct{}{}

	fmt.Printf("s size: %v\n", unsafe.Sizeof(s)) // 0 байт

}
