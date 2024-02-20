package main

import (
	"errors"
	"fmt"
)

// бинарный поиск
func main() {
	var array = []int{1, 4, 7, 22, 23, 63, 63, 70, 223, 1111}
	fmt.Println(array)
	fmt.Println("Введите число")
	var x int
	fmt.Scan(&x)
	ind, err := binarySearch(x, array, 0)
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println("Индекс первого вхождения в массие:")
		fmt.Println(ind)
	}
}
func binarySearch(x int, arr []int, step int) (int, error) {
	med := int(len(arr) / 2)
	if arr[med] == x {
		return med + step, nil
	}
	if (len(arr) <= 2) && (arr[0] != x) {
		return 0, errors.New("В массиве нет такого числа")
	}
	if x < arr[med] {
		return binarySearch(x, arr[:med], step)
	} else {
		return binarySearch(x, arr[med:], med+step)
	}

}
