package main

import "fmt"

type Set []string

func NewSet() *Set {
	return &Set{}
}

func (s Set) InSet(str string) bool {
	for _, value := range s {
		if value == str {
			return true
		}
	}
	return false
}

func (s *Set) Insert(str string) {
	*s = append(*s, str)
}

func main() {
	arr_str := []string{"cat", "cat", "dog", "cat", "tree"}
	set := NewSet()

	for _, value := range arr_str {
		if ok := set.InSet(value); !ok {
			set.Insert(value)
		}
	}

	fmt.Println("Исходный набор: ", arr_str)
	fmt.Println("Множество: ", set)
}
