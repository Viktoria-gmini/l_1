package main

import "fmt"

// Что выведет данная программа и почему?
func main() {
	var a = []int8{1, 2, 3, 4, 5}
	someAction(a, 6)
	fmt.Println(a)
}
func someAction(v []int8, b int8) {
	v[0] = 100
	v = append(v, b)
}
