// Реализовать все возможные способы остановки выполнения горутины.
package main

import (
	"context"
	"fmt"
	"sync"
	"time"
)

func main() {
	channel_way_to_stop_goroutine()
	context_way_to_stop_goroutine()
	boolean_way_to_stop_goroutine()

}
func channel_way_to_stop_goroutine() {
	stop := make(chan bool)

	go func() {
		for {
			select {
			case <-stop:
				return
			default:
				fmt.Println("Горутина 1 работает")
				time.Sleep(time.Second)
			}
		}
	}()

	time.Sleep(5 * time.Second)
	stop <- true
	fmt.Println("Горутина 1 остановлена")
}
func context_way_to_stop_goroutine() {
	ctx, cancel := context.WithCancel(context.Background())

	go func() {
		for {
			select {
			case <-ctx.Done():
				return
			default:
				fmt.Println("Горутина 2 работает")
				time.Sleep(time.Second)
			}
		}
	}()

	time.Sleep(5 * time.Second)
	cancel()
	fmt.Println("Горутина 2 остановлена")
}
func boolean_way_to_stop_goroutine() {
	var wg sync.WaitGroup
	stop := false

	wg.Add(1)
	go func() {
		defer wg.Done()
		for {
			fmt.Println("Горутина 3 работает")
			time.Sleep(time.Second)
			if stop {
				return
			}
		}
	}()

	time.Sleep(5 * time.Second)
	stop = true

	wg.Wait()
	fmt.Println("Горутина 3 остановлена")
}
