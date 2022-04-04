package main

import "fmt"

/*
Разработать конвейер чисел. Даны два канала: в первый пишутся числа (x) из массива,
во второй — результат операции x*2, после чего данные из второго канала должны выводиться в stdout.
 */

func ToChanOne(nums []int) <-chan int {
	out := make(chan int)
	go func() { // используем горутину иначе deadlock
		for _, value := range nums {
			out <- value
		}
		close(out)
	}()
	return out
}

func FromOneToSecond(in <- chan int) <-chan int {
	out := make(chan int)
	go func() {
		for value := range in {
			out <- value*value
		}
		close(out)
	}()
	return out
}

func main() {
	nums := []int{1, 2, 3, 4, 5}
	for value := range FromOneToSecond(ToChanOne(nums)) {
		fmt.Println(value)
	}

}