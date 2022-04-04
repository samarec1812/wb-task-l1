package main

import (
	"fmt"
	"strings"
)

/*
	Разработать программу, которая переворачивает слова в строке.
	Пример: «snow dog sun — sun dog snow».
*/


// Version 1
// если отсутствуют знаки препинания и т.п
func ReverseWordInString(str string) string {
	strArr := strings.Split(str, " ")

	return strings.Join(ReverseSliceStr(strArr), " ")
}

func ReverseSliceStr(strArr []string) []string {
	for i := 0; i < len(strArr)/2; i++{
		strArr[i], strArr[len(strArr)-1-i] = strArr[len(strArr)-1-i], strArr[i]
	}
	return strArr
}

func main() {
	fmt.Println("Word 1: ", ReverseWordInString("snow dog sun"))

}