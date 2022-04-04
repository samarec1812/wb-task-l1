package main

import "fmt"

/*
	Разработать программу, которая переворачивает подаваемую на ход строку
	(например: «главрыба — абырвалг»). Символы могут быть unicode.
*/

// version 1
func Reverse(str string) string {
	res := ""
	for _, value := range str {
		res = string(value) + res
	}
	return res
}

// version 2
func ReverseB(str string) string {
	arr := []rune(str)
	res := ""
	for i := len(arr)-1; i > -1; i-- {
		res += string(arr[i])
	}
	return res
}

func main() {
	fmt.Println("function result reverse: ", ReverseB("function"))
	fmt.Println("String with chines symbol result reverse: ", ReverseB("The quick brown 狐 jumped over the lazy 犬"))
	fmt.Println("String with russian symbol result reverse: ", Reverse("главрыба"))
}