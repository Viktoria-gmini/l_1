//программа, которая конкурентно рассчитывает значение
//квадратов чисел взятых из массива (2,4,6,8,10) и выведет их квадраты в stdout.

package main

import (
	"fmt"
	"time"
)

func sqr(num int) {
	fmt.Println(num * num)
	return
}

func main() {
	array := [4]int{2, 4, 6, 8}
	go sqr(array[0])
	go sqr(array[3])
	go sqr(array[2])
	go sqr(array[1])

	time.Sleep(time.Millisecond)
}
