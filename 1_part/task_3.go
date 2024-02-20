// Дана последовательность чисел: 2,4,6,8,10. Найти сумму их квадратов(2^2+3^2+4^2….)
// с использованием конкурентных вычислений.
package main

import (
	"fmt"
	"sync"
)

func sqr(num int) int {
	return num * num
}
func main() {
	array := [4]int{2, 4, 6, 8}
	sqrs := make(chan int, 4)
	var wg sync.WaitGroup
	for _, num := range array {
		wg.Add(1)
		go func(c int) {
			defer wg.Done()
			sqrs <- sqr(c)
			fmt.Println("num ", c)
			return
		}(num)
	}
	go func() {
		wg.Wait()
		close(sqrs)
	}()
	sum := 0
	for s := range sqrs {
		sum += s
	}

	fmt.Println(sum)

}
