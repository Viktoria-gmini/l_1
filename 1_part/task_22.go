// Разработать программу, которая перемножает, делит, складывает,
// вычитает две числовых переменных a,b, значение которых > 2^20.
package main

import (
	"fmt"
	"math/big"
)

func main() {
	a := big.NewInt(1)
	b := big.NewInt(1)

	// Установим значения a и b больше чем 2^20
	a.Exp(big.NewInt(2), big.NewInt(20), nil)
	b.Exp(big.NewInt(2), big.NewInt(21), nil)

	// Умножение
	result := new(big.Int).Mul(a, b)
	fmt.Printf("Умножение: %s * %s = %s\n", a.String(), b.String(), result.String())

	// Деление
	result = new(big.Int).Div(a, b)
	fmt.Printf("Деление: %s / %s = %s\n", a.String(), b.String(), result.String())

	// Сложение
	result = new(big.Int).Add(a, b)
	fmt.Printf("Сложение: %s + %s = %s\n", a.String(), b.String(), result.String())

	// Вычитание
	result = new(big.Int).Sub(a, b)
	fmt.Printf("Вычитание: %s - %s = %s\n", a.String(), b.String(), result.String())
}
