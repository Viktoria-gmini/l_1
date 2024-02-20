// Разработать программу, которая в рантайме способна определить тип переменной:
// int, string, bool, channel из переменной типа interface{}.
package main

import (
	"fmt"
	"reflect"
)

func main() {
	var data interface{}

	data = 10
	checkType(data)

	data = "Hello"
	checkType(data)

	data = true
	checkType(data)

	data = make(chan int)
	checkType(data)
}

func checkType(data interface{}) {
	dataType := reflect.TypeOf(data).String()
	fmt.Printf("Variable is of type %s\n", dataType)
}
