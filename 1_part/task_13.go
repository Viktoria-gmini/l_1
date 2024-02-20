// Поменять местами два числа без создания временной переменной.
package main

import (
	"fmt"
)

var a int = -100
var b int = 9

func main() {

	a = a + b
	b = a - b
	a = a - b

	fmt.Println(a)
	fmt.Println(b)

}


