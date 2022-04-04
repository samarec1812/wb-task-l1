package main

import (
	"fmt"
	"strings"
)

/*
	Разработать программу, которая проверяет, что все символы в строке уникальные
	(true — если уникальные, false etc).
	Функция проверки должна быть регистронезависимой.
	Например:
		abcd — true
		abCdefAaf — false
		aabcd — false
*/

func Check(str string) bool {
	str = strings.ToLower(str) // преобразуем всё к нижнему регистру для удобства
	table := make(map[string]bool)
	for _, value := range str {
		if flag, ok := table[string(value)]; !ok {
			table[string(value)] = true
		} else if flag {
			return false
		}
	}
	return true
}

func main() {
	fmt.Println("Check abcd: ", Check("abcd"))
	fmt.Println("Check abCdefAaf: ", Check("abCdefAaf"))
	fmt.Println("Check aabcd: ", Check("aabcd"))
}