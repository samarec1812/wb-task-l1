package main

/*
	К каким негативным последствиям может привести данный фрагмент кода,
	и как это исправить? Приведите корректный пример реализации.
*/

import (
	"fmt"
	"golang.org/x/exp/rand"
	"strings"
)
//
//var justString string
//var justString2 string
//
//const alphabet = "ABCDEFGHIJKLMNOPQRSTUVWXYZabcdefghijklmnopqrstuvwxyz0123456789"
//
//// в целом ничего особенного
//func SomeFunc() {
//	v := createHugeString(1 << 10)
//	builder :=  strings.Builder{}
//	builder.WriteString(v[:100])
//	justString = builder.String()
//}
//
//func SomeFuncB() {
//	v := createHugeString(1 << 10)
//	justString2 = v[:100]
//}
//
//// 1 is binary 0001
//// 1 << 10 mean: 00010000000000
//// createHugeString возвращает строку
//func createHugeString(value int) string {
//	var s strings.Builder
//	for i := 0; i < value; i++ {
//		s.WriteByte(alphabet[rand.Int()%len(alphabet)])
//	}
//	return s.String()
//}

//// createHugeStringB возвращает slice
//func createHugeStringB(value int) []byte {
//	var s strings.Builder
//	for i := 0; i < value; i++ {
//		s.WriteByte(alphabet[rand.Int()%len(alphabet)])
//	}
//	return []byte(s.String())
//}

func main() {
	fmt.Println(1 << 10)
	SomeFunc()
	SomeFuncB()
}

// во-первых это не скомпилируется из-за наличия необъявленной функции
// Решение объявить функцию createHugeString

// во-вторых описать функцию(что она делает)
// потому что если она например возвращает пустую строку
// в функции someFunc возникнет паника, т.к слайс имеет длину 0

// Реализуем логику. Скорее всего имеется в виду создать строку длиной 1 << 10

// также есть недочёт в плане unused variable justString
// это не вызовет ошибки компиляции
