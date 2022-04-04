package main

import (
	"fmt"
	"sync"
)

func calculate(value int) {
	fmt.Print(value*value, " ")
}

func main() {
	arr := []int{2, 4, 6, 8, 10}
	wg := new(sync.WaitGroup)
	//mu := new(sync.Mutex)

	for _, value := range arr {
		wg.Add(1)
		go func(wg *sync.WaitGroup, value int) {
			defer wg.Done()
			//mu.Lock()
			fmt.Println(value* value)
			//mu.Unlock()

		}(wg, value)

	}
	wg.Wait()

}