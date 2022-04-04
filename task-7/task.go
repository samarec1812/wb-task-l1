package main

import (
	"fmt"
	"sync"
)

/*
	Реализовать конкурентную запись данных в map.
*/


func main() {
	var counters = map[int]int{}
	mu := &sync.Mutex{}
	for i := 0; i < 5; i++ {
		go func(counter map[int]int, value int, mu *sync.Mutex) {
			for j := 0; j < 5; j++ {
				mu.Lock()
				counter[value*5+j]++
				mu.Unlock()
			}
		}(counters, i, mu)
	}

	mu.Lock()
	fmt.Println("counters result", counters)
	mu.Unlock()
}