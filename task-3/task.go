package main

import (
	"fmt"
	"sync"
)

func main() {
	arr := []int{2, 4, 6, 8, 10}
	wg := sync.WaitGroup{}

	sum := 0
	for i := 0; i < len(arr); i++ {
		wg.Add(1)
		go func(value int) {
			sum += value* value
		// 	fmt.Println(value* value)
			wg.Done()
		}(arr[i])

	}

	wg.Wait()
	fmt.Println(sum)


}
