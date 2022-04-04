package main

import "fmt"

func CheckIn(value int, arr []int) bool {
	for _, v := range arr {
		if v == value {
			return true
		}
	}
	return false
}

func main() {

	a := []int{1, 5, 3, 4}
	b := []int{2, 3, 4, 5}
	c := make([]int, 0)


		for _, value := range a {
			for _, value2 := range b {
				if value == value2 && !CheckIn(value, c) {
					c = append(c, value)
				}
			}
		}
	fmt.Println(c)
}
