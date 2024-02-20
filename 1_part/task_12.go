// Имеется последовательность строк - (cat, cat, dog, cat, tree)
// создать для нее собственное множество.
package main

import (
	"fmt"
)

func main() {
	strings := []string{"cat", "cat", "dog", "cat", "tree"}
	set_ := go_set(strings)
	fmt.Println(set_)
}
func go_set(strings []string) []string {

	m := make(map[string]bool)
	for _, item := range strings {
		m[item] = true
	}
	var strings_ []string
	for key := range m {
		strings_ = append(strings_, key)
	}
	return strings_
}
