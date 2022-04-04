package main

import (
	"fmt"
	"sync"
	"sync/atomic"
)

/*
	Реализовать структуру-счетчик, которая будет инкрементироваться в конкурентной среде.
	По завершению программа должна выводить итоговое значение счетчика.
*/

type Ticker struct {
	value int64
}

func (t *Ticker) inc(wg *sync.WaitGroup) {
	defer wg.Done()
	atomic.AddInt64(&t.value, 1)
}

// проблема: состояние гонки
func main() {
	t := Ticker{0}
	wg := &sync.WaitGroup{}
	for i := 0; i < 1000; i++ {
		wg.Add(1)
		go t.inc(wg)
	}
	wg.Wait()
	fmt.Println(t)
}