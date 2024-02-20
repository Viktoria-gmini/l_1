// var justString string
// func someFunc() {
//   v := createHugeString(1 << 10)
//   justString = v[:100]
// }

//	func main() {
//	  someFunc()
//	}
//
// как я поняла, проблема данного кода заключается в том,
// что мы пытаемся обратиться к элементам слайса, которых
// может и не быть. Добавила проверку на длину v
package main

var justString string

func someFunc() {
	v := createHugeString(1 << 10)
	if len(v) >= 100 {
		justString = v[:100]
	}
}

func main() {
	someFunc()
}
