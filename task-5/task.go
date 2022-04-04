package main

import (
	"context"
	"fmt"
	"math/rand"
	"time"
)

/*
Разработать программу, которая будет последовательно отправлять значения в канал,
а с другой стороны канала — читать.
По истечению N секунд программа должна завершаться.
*/

func ReadAndWrite(ctx context.Context) {
	ch := make(chan int, 1)
	for {
		if ctx.Err() != nil {
			return
		}
		ch <- rand.Int()
		fmt.Println(<-ch)
	}
}

func ReadAndWriteB(ctx context.Context) {
	ch := make(chan int, 1)
	for {
		select {
		case <-ctx.Done():
			return
		case val := <-ch:
			fmt.Println(val)
			time.Sleep(10 * time.Millisecond)
		default:
			ch <- rand.Int()
		}
	}
}

func ReadAndWriteC(duration time.Duration) {
	cancel := make(chan struct{})
	data := make(chan int)
	flag := false

	go func(cancel chan struct{}, data chan int) {
		val := 0
		for {
			select {
			case <-cancel:
				return
			case data <- val:
				val++
			}
		}
	}(cancel, data)

	time.AfterFunc(duration * time.Second, func() {
		cancel <- struct{}{}
		flag = true
	})

	LOOP:
	for value := range data {
		fmt.Println("value: ", value)
		time.Sleep(1*time.Second)
		if flag {
			break LOOP
		}
 	}

}

func main() {
	var N int
	fmt.Scanln(&N)
	//ctx, cancel := context.WithTimeout(context.Background(), time.Duration(N)*time.Second)
	//defer cancel()
	//// ch := make(chan int)
	//ReadAndWriteB(ctx)

	ReadAndWriteC(time.Duration(N))
	//time.Sleep(5 * time.Second)

}
