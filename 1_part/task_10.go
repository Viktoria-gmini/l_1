// Дана последовательность температурных колебаний:
// -25.4, -27.0 13.0, 19.0, 15.5, 24.5, -21.0, 32.5.
// Объединить данные значения в группы с шагом в 10 градусов.
// Последовательность в подмножноствах не важна.

package main

import (
	"fmt"
	"math"
)

func main() {
	super_map := make(map[int][]float64)
	values := []float64{-25.4, -27.0, 13.0, 19.0, 15.5, 24.5, -21.0, 32.5}
	for _, value := range values {
		rounded_value := int(math.Floor(value/10) * 10)
		super_map[rounded_value] = append(super_map[rounded_value], value)
	}
	fmt.Println(super_map)
}
