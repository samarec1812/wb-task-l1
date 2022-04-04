package main

import (
	"math/rand"
	"strings"
	"testing"
)


var justString string
var justString2 string

const alphabet = "ABCDEFGHIJKLMNOPQRSTUVWXYZabcdefghijklmnopqrstuvwxyz0123456789"

// в целом ничего особенного
func SomeFunc() {
	v := createHugeString(1 << 10)
	builder :=  strings.Builder{}
	builder.WriteString(v[:100])
	justString = builder.String()
}

func SomeFuncB() {
	v := createHugeString(1 << 10)
	justString2 = v[:100]
}

// 1 is binary 0001
// 1 << 10 mean: 00010000000000
// createHugeString возвращает строку
func createHugeString(value int) string {
	var s strings.Builder
	for i := 0; i < value; i++ {
		s.WriteByte(alphabet[rand.Int()%len(alphabet)])
	}
	return s.String()
}


func BenchmarkSomeFunc(b *testing.B) {
	for i := 0; i < b.N; i++ {
		SomeFunc()
	}
}

func BenchmarkSomeFuncB(b *testing.B) {
	for i := 0; i < b.N; i++ {
		SomeFuncB()
	}
}

// run
//  go test -bench . -benchmem -cpuprofile="cpu.out" -memprofile="mem.out" -memprofilerate="1" task_test.go