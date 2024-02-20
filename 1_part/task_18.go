package main

import (
	"fmt"
	"sync"
	"time"
)

// структура-счётчик, инкрементирующаяся в конкурентной среде
type Counter struct {
	mu      sync.Mutex
	counter int
}

func (c *Counter) Increment() {
	c.mu.Lock()
	defer c.mu.Unlock()
	c.counter++
}

func main() {
	done := make(chan bool)
	var c Counter
	c.counter = 0
	for i := 0; i < 2; i++ {
		go func() {
			for {
				select {
				case <-done:
					return
				default:
					c.Increment()
				}
			}
		}()
	}
	time.Sleep(3 * time.Second) // Ожидаем N секунд
	done <- true
	fmt.Println("Значение счётчика:")
	fmt.Println(c.counter)

}
