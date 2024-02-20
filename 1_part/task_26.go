// Разработать программу, которая проверяет, что все символы
// в строке уникальные (true — если уникальные, false etc).
// Функция проверки должна быть регистронезависимой.

package main

import (
	"fmt"
	"strings"
)

func isUnique(str string) bool {
	// Приводим строку к нижнему регистру
	str = strings.ToLower(str)

	// Создаем мапу для отслеживания уникальных символов
	uniqueChars := make(map[rune]bool)

	// Перебираем каждый символ в строке
	for _, char := range str {
		// Если символ уже присутствует в мапе, значит он не уникален
		if uniqueChars[char] {
			return false
		}

		// Добавляем символ в мапу
		uniqueChars[char] = true
	}

	// Если все символы уникальны, возвращаем true
	return true
}

func main() {
	str1 := "abcd"
	fmt.Println(isUnique(str1)) // Вывод: true

	str2 := "abCdefAaf"
	fmt.Println(isUnique(str2)) // Вывод: false

	str3 := "aabcd"
	fmt.Println(isUnique(str3)) // Вывод: false

	str4 := "golang"
	fmt.Println(isUnique(str4)) // Вывод: true
}
