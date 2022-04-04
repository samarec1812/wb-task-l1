package main

import "fmt"

func main() {
	var value interface{}
	value = make(chan interface{})
	switch value.(type) {
	case int:
		fmt.Println("type int")
	case string:
		fmt.Println("type string")
	case bool:
		fmt.Println("type bool")
	case chan interface{}:
		fmt.Println("type chan")
	default:
		fmt.Println("Don't know this type")


	}


}
