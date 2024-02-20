package main

import (
	"fmt"
	"sync"
)

// Что выведет данная программа и почему?
func main() {
	wg := sync.WaitGroup{}
	for i := 0; i < 5; i++ {
		wg.Add(1)
		go func(wg sync.WaitGroup, i int) {
			fmt.Println(i)
			wg.Done()
		}(wg, i)
	}
	wg.Wait()
	fmt.Println("exit")
}

//выведет в рандомном порядке цифры от 0 до 4 и выдаст deadlock.
//Это связано с тем, что вейтгруппа не передаётся в анонимую функцию,
//и на каждй итерации создаётся лишь её копия. Эту проблему можно
// решить с помощью указателей
