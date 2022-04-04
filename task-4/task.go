package main

import (
	"fmt"
	"math/rand"
	"os"
	"os/signal"
	"runtime"
	_ "sync"
	"time"
)

/*
Реализовать постоянную запись данных в канал (главный поток).
Реализовать набор из N воркеров, которые читают произвольные данные из канала и выводят в stdout.
Необходима возможность выбора количества воркеров при старте.
Программа должна завершаться по нажатию Ctrl+C. Выбрать и обосновать способ завершения работы всех воркеров.
*/


func worker(numWorker int, in chan int) {
	// defer wg.Done()
	for value := range in {
		fmt.Printf("Worker: %d value: %d\n", numWorker, value)
		runtime.Gosched()
	}
	fmt.Println("Finish Worker: ", numWorker)
	time.Sleep(2*time.Millisecond)
}


func main() {
	//ctx, cancel := context.WithCancel(context.Background())
	var N int
	fmt.Scanln(&N)

	// wg := &sync.WaitGroup{}

	in := make(chan int)
	sigChannel := make(chan os.Signal)
	signal.Notify(sigChannel, os.Interrupt)
	for i := 0; i < N; i++ {
		// wg.Add(1)
		go worker(i, in)
	}
	//time.Sleep(time.Millisecond)
	//

	LOOP:
	for  {
		select {
		case msg := <-sigChannel:
			switch msg {
			case os.Interrupt:
				break LOOP
			}
		case in <- rand.Int() % 1000:
		}
	}
	//wg.Wait()
	close(in)
	time.Sleep(1000*time.Millisecond)
}