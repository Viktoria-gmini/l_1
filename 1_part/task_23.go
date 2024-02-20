// Удалить i-ый элемент из слайса.

package main

import "fmt"

func removeAtIndex(slice []int, index int) []int {
	return append(slice[:index], slice[index+1:]...)
}

func main() {
	// Создаем слайс
	slice := []int{1, 2, 3, 4, 5}

	// Удаляем 3-ий элемент (индекс 2)
	index := 2
	result := removeAtIndex(slice, index)

	fmt.Printf("Слайс до удаления элемента: %v\n", slice)
	fmt.Printf("Слайс после удаления элемента: %v\n", result)
}
