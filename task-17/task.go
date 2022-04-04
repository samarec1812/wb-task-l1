package main

import (
	"errors"
	"fmt"
	"golang.org/x/exp/constraints"
)

/*
	Реализовать бинарный поиск встроенными методами языка
*/

// Compare two values for less
func Less[T constraints.Ordered] (value, value2 T) (bool, error) {
	switch (interface{})(value).(type) {
	case int:
		return (interface{})(value).(int) < (interface{})(value2).(int), nil
	case string:
		return (interface{})(value).(string) < (interface{})(value2).(string), nil
	case float32:
		return (interface{})(value).(float32) < (interface{})(value2).(float32), nil
	case float64:
		return (interface{})(value).(float64) < (interface{})(value2).(float64), nil
	}
	return false, errors.New("Invalid type")
}

// Iterative model
func BinarySearch[T constraints.Ordered](arr []T, value T) int {
	first := 0
	last := len(arr) - 1
	for first <= last {
		middle := first + (last - first)/2
		if arr[middle] == value {
			return middle
		} else if ok, err := Less[T](arr[middle], value); ok && err == nil {
			first = middle + 1
		} else if err != nil {
			return -1
		} else {
			last = middle - 1
		}
	}
	return -1
}

// Recursive model
func BinarySearchR(arr []int, value, first, last int) int {
	if first <= last {
		middle := first + (last-first)/2
		if arr[middle] == value  {
			return middle
		} else if arr[middle] < value {
			return BinarySearchR(arr, value, middle + 1, last)
		} else {
			return BinarySearchR(arr, value, first, middle - 1)
		}
	}
	return -1
}

//
func Search(size int, f func(int) bool) int {
	i, j := 0, size
	for i < j {
		h := i + (j - i)/2
		if !f(h) {
			i = h + 1
		} else {
			j = h
		}
	}
	return i
}

func main() {

	data := []int{-1, 0, 1, 3,  5, 7}
	data2 := []string{"aba", "bob", "cisco", "exel"}
	data3 := []byte{1, 2, 3, 4, 5}
	index := BinarySearch[int](data, 3)
	fmt.Println(index)
	fmt.Println(data)

	index2 := BinarySearch[string](data2, "cisco")
	fmt.Println(index2)
	fmt.Println(data2)

	index3 := BinarySearch[byte](data3, 1)
	fmt.Println(index3)
	fmt.Println(data3)
}