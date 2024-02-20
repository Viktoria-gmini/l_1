// Разработать конвейер чисел. Даны два канала: в первый пишутся числа (x) из массива,
//
//	во второй — результат операции x*2,
//
// после чего данные из второго канала должны выводиться в stdout.
package main

import (
	"fmt"
	"time"
)

func main() {
	channel_1 := make(chan int)
	channel_2 := make(chan int, 1)
	go func() {
		i := 0
		for {
			i++
			channel_1 <- i
			time.Sleep(time.Second * 1)
		}
	}()
	go func() {
		for {
			select {
			case x := <-channel_1:
				fmt.Println(x)
				y := x * x
				channel_2 <- y
			case y := <-channel_2:
				fmt.Println(y)
			}
		}
	}()
	time.Sleep(time.Second * 60)
}
