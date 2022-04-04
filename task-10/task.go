package main

import "fmt"

/*
	Дана последовательность температурных колебаний: -25.4, -27.0 13.0, 19.0, 15.5, 24.5, -21.0, 32.5.
	Объединить данные значения в группы с шагом в 10 градусов. Последовательность в подмножноствах не важна.
	Пример: -20:{-25.0, -27.0, -21.0}, 10:{13.0, 19.0, 15.5}, 20: {24.5}, etc.
*/


func main() {
	temperature := []float32{-29.4, -27.0,  13.0, 19.0, 15.5, 24.5, -21.0, 32.5, 123.3}
	sets := make(map[int][]float32)

	for _, value := range temperature {
		key := 10*int(value/10)
		//fmt.Println(key)
		if  _, ok := sets[key]; !ok {
				arr := make([]float32, 0)
				sets[key] = arr
		}
		sets[key] = append(sets[key], value)
	}
	fmt.Println(sets)

}