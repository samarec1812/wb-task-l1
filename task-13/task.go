package main

import "fmt"

func Swap(i, j *int) {
	*i, *j = *j, *i
}



func main() {
	a, b := 1, 2
	fmt.Println(a, b)
	Swap(&a, &b)
	fmt.Println(a, b)

}
