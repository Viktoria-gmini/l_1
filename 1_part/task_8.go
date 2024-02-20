// Дана переменная int64. Разработать программу которая устанавливает i-й бит в 1 или 0.
package main

import "fmt"

func setBit(num int64, i int, value int) int64 {
	if value == 1 {
		num |= 1 << i
	} else {
		num &= ^(1 << i)
	}
	return num
}

func main() {
	var num int64 = 100000000
	i := 6
	value := 1
	result := setBit(num, i, value)
	fmt.Printf("The result after setting bit %d to %d: %d\n", i, value, result)
}
