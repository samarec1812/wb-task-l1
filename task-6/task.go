package main

import (
	"context"
	"fmt"
	"sync"
	"time"
)

/*
	Реализовать все возможные способы остановки выполнения горутины.
*/

// 2 через каналы 1) (канал отмены)
// 2 через context


// посылаем сигнал отмены через канал отмены
func PrintA(data chan int, quit chan struct{}) {
	for {
		select {
		case m := <- data:
			fmt.Println(m)
		case <-quit:
			return
		}
	}
}

// через закрытие канала
func PrintB(data chan int) {
	for {
		select {
		case m, ok := <- data:
			if !ok {
				data = nil
			} else {
				fmt.Println("case %d", m)
				for i := 0; i < 10; i++ {
					fmt.Print(i, " ")

				}
				fmt.Println()
			}
		default:
			return
		}
	}
}


func PrintC(ctx context.Context) {
	for {
		select {
		case <-ctx.Done():
			return
		default:
			for i := 0; i < 10; i++ {
				fmt.Print(i, " ")
			}
			fmt.Println()
		}
	}
}

func PrintD(ctx context.Context) {
	time.Sleep(5*time.Second)
	for {
		select {
		case <-ctx.Done():
			return
		default:
			for i := 0; i < 10; i++ {
				fmt.Print(i, " ")
			}
			fmt.Println()
		}
	}
}




func main() {
	// case A
	//data := make(chan int)
	//quit := make(chan struct{})
	//go Print(data, quit)
	//
	//data <- 1
	//data <- 2
	//quit <- struct{}{}

	// case B
	//data := make(chan int)
	//go PrintB(data)
	//data <- 2
	//close(data)
	//time.Sleep(5 * time.Millisecond)
	//

	// case C
	//ctxWithCancel, cancelFunc := context.WithCancel(context.Background())
	//go PrintC(ctxWithCancel)
	//time.Sleep(10*time.Second)
	//cancelFunc()
	//time.Sleep(5*time.Second)

	// case D
	ctx, cancelFunc := context.WithCancel(context.Background())

	wg := sync.WaitGroup{}
	wg.Add(1)
	go func() {
		PrintD(ctx)
		wg.Done()
	}()
	//go PrintC(ctxWithCancel)
	time.Sleep(6*time.Second)
	cancelFunc()
	time.Sleep(6*time.Second)


}