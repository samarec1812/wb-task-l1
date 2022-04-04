package main

import "fmt"

/*
	Удалить i-ый элемент из слайса.
*/

// Удаление на основе append
func DeleteA[T any](arr []T, index int) []T {
	return append(arr[:index], arr[index+1:]...)
}

// Удаление через copy
func DeleteB[T any](arr []T, index int) []T {
	copy(arr[index:], arr[index+1:])
	return arr[:len(arr)-1]
}

// Удаление через создание нового массива
func DeleteC[T any](arr []T, index int) []T {
	newArr := make([]T, 0)
	return append(append(newArr, arr[:index]...), arr[index+1:]...)
}

func main() {
	arr := []int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9}
	fmt.Println("Before erase elem:", arr)
	arr = DeleteA[int](arr, 3)
	fmt.Println("After erase elem func DeleteA", arr)

	arr = DeleteB[int](arr, 3)
	fmt.Println("After erase elem func DeleteB", arr)

	arr = DeleteC[int](arr, 3)
	fmt.Println("After erase elem func DeleteC", arr)
}
