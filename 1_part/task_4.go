package main

import (
	"context"
	"fmt"
	"os"
	"os/signal"
	"syscall"
	"time"
)

var (
	N       int
	channel chan string
)

func main() {
	fmt.Print("Введите количество работников: ")
	fmt.Scan(&N)
	channel = make(chan string, N)

	ctx, cancel := context.WithCancel(context.Background())

	// Обработка сигнала завершения
	signalCh := make(chan os.Signal, 1)
	signal.Notify(signalCh, os.Interrupt, syscall.SIGTERM)
	go func() {
		<-signalCh
		cancel()
	}()

	// Запуск N воркеров
	startWorkers(ctx)

	// Запись данных в канал
	for i := 1; ; i++ {
		data := fmt.Sprintf("data №%d", i)
		select {
		case channel <- data:
			// данные успешно записаны в канал
		case <-ctx.Done():
			// если канал закрыт
			return
		}
	}

}

func startWorkers(ctx context.Context) {
	for i := 0; i < N; i++ {
		go func(workerNum int) {
			for {
				select {
				case data := <-channel:
					fmt.Println("worker №", workerNum, ": ", data)
				case <-ctx.Done():
					fmt.Println("Worker", workerNum, "stopped")
					return
				}
				time.Sleep(time.Second)
			}
		}(i)
	}

}
