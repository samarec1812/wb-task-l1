package main

import (
	"errors"
	"fmt"
	"golang.org/x/exp/constraints"
)

/*
	Реализовать быструю сортировку массива (quicksort) встроенными методами языка.
*/


// QuickSortRange сортирует в границах
func QuickSortRange[T constraints.Ordered](arr []T, left, right int) {
	if len(arr) <= 1 {
		return
	}
	if left < right {
		index := partition[T](arr, left, right)
		QuickSortRange(arr, left, index-1)
		QuickSortRange(arr, index+1, right)
	}
}

func QuickSort[T constraints.Ordered](arr []T) {
	QuickSortRange[T](arr, 0, len(arr)-1)
}

// Compare two values for less
func LessEqual[T constraints.Ordered] (value, value2 T) (bool, error) {
	switch (interface{})(value).(type) {
	case int:
		return (interface{})(value).(int) <= (interface{})(value2).(int), nil
	case string:
		return (interface{})(value).(string) <= (interface{})(value2).(string), nil
	case float32:
		return (interface{})(value).(float32) <= (interface{})(value2).(float32), nil
	case float64:
		return (interface{})(value).(float64) <= (interface{})(value2).(float64), nil
	}
	return false, errors.New("Invalid type")
}

func partition[T constraints.Ordered](arr []T, left, right int) int {
	index := left - 1
	mainPoint := arr[right]
	for i := left; i < right; i++ {
		if  ok, _ := LessEqual[T](arr[i], mainPoint); ok{
			index += 1
			arr[index], arr[i] = arr[i], arr[index]
		}
	}
	arr[index+1], arr[right] = arr[right], arr[index+1]
	return index + 1
}

func main() {
	arr := []int{3, 4, 1, 2, 5, 7, -1, 0}
	fmt.Println(arr)
	QuickSort[int](arr)
	fmt.Println(arr)
}