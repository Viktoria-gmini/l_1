package main

import "fmt"

// Что выведет данная программа и почему?
func main() {
	slice := []string{"a", "a"}

	func(slice []string) {
		slice = append(slice, "a") //создаётся новый слайс
		slice[0] = "b"             //манипуляции с новым слайсом
		slice[1] = "b"
		fmt.Print(slice) //выводится новый слас
	}(slice)
	fmt.Print(slice) //на старый слайс ничего не повлияло
}
