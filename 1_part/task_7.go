// Реализовать конкурентную запись данных в map.
package main

import (
	"fmt"
	"sync"
	"time"
)

var m sync.Map

func main() {
	i := 0
	j := 0
	stop := make(chan bool)
	go func() {
		for {
			select {
			case <-stop:
				return
			default:
				m.Store(i, "value1")
				value, _ := m.Load(i)
				fmt.Println(i, " ", value)
				i++
				time.Sleep(time.Second)
			}
		}
	}()
	go func() {
		for {
			select {
			case <-stop:
				return
			default:
				m.Store(j, "value2")
				value, _ := m.Load(j)
				fmt.Println(j, " ", value)
				j--
				time.Sleep(time.Second)
			}
		}
	}()
	time.Sleep(time.Second * 10)
	stop <- true
}
