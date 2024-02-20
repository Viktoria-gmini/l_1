package main

import "fmt"

//быстрая сортировка
func QuickSort(arr []int, low, high int) {
	if low < high {
		pivotIndex := partition(arr, low, high)
		QuickSort(arr, low, pivotIndex-1)
		QuickSort(arr, pivotIndex+1, high)
	}
}

func partition(arr []int, low, high int) int {
	pivot := arr[high]
	i := low - 1

	for j := low; j <= high-1; j++ {
		if arr[j] < pivot {
			i++
			arr[i], arr[j] = arr[j], arr[i]
		}
	}
	arr[i+1], arr[high] = arr[high], arr[i+1]
	return i + 1
}

func main() {
	array := []int{3, 14, 1, 7, 9, 8, 11, 6, 4, 2}
	QuickSort(array, 0, len(array)-1)
	fmt.Println(array)
}
