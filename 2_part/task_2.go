// Интерфейсы в Go - это общий способ определения абстракций, которые не привязаны
// к конкретным типам данных. Они позволяют определить набор методов,
// которые должны быть реализованы в других типах данных,
// чтобы они могли быть использованы в качестве аргументов функций или объектов.
package main

import (
	"fmt"
	"strings"
)

type martian struct{}

type laser int

func (m martian) talk() string {
	return "nack nack"
}
func (l laser) talk() string {
	return strings.Repeat("pew ", int(l))
}

func main() {
	var t interface {
		talk() string
	}
	t = martian{}
	fmt.Println(t.talk()) // Выводит: nack nack

	t = laser(3)
	fmt.Println(t.talk()) // Выводит: pew pew pew
}

// В отличие от Java, в Go martian и laser напрямую не объявляют,
//  что они имплементируют интерфейс.
