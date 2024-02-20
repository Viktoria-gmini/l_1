package main

// Разработать программу, которая будет последовательно отправлять
// значения в канал, а с другой стороны канала — читать.
// По истечению N секунд программа должна завершаться.
import (
	"fmt"
	"time"
)

var (
	channel = make(chan string)
)

func main() {
	waitTime := 5
	done := make(chan struct{}) // Канал для прерывания выполнения

	go func() {
		i := 0
		for {
			select {
			case <-time.After(time.Second * time.Duration(waitTime)):
				close(done)
				return
			default:
				i++
				inputData := fmt.Sprintf("data №%d", i)
				channel <- inputData
			}
		}

	}()

	go func() {
		for {
			select {
			case <-done: // Если получен сигнал о прерывании, выходим из функции
				return
			case outputData := <-channel:
				fmt.Println(outputData)
			}
		}
	}()
	time.Sleep(time.Duration(waitTime) * time.Second) // Ожидаем N секунд
	close(done)

	fmt.Println("Программа завершилась")
}
