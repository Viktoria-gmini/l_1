// Реализовать пересечение двух неупорядоченных множеств.
package main

import (
	"fmt"
)

func main() {
	set1 := []int{1, 2, 3, 4, 5, 0}
	set2 := []int{6, 3, 4, 1, 7, 8}

	result := conjunction(set1, set2)
	fmt.Println(result)
}

func conjunction(set1, set2 []int) []int {
	dict := make(map[int]bool)
	for _, val := range set1 {
		dict[val] = true
	}
	result := []int{}
	for _, val := range set2 {
		if dict[val] {
			result = append(result, val)
		}
	}

	return result
}
