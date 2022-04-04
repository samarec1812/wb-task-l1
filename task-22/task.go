package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"os/exec"
)

/*
	Разработать программу, которая перемножает, делит, складывает,
	вычитает две числовых переменных a,b, значение которых > 2^20.
*/



// Numbers вызывает питоновский скрипт
func Numbers(a, b int) {
	cmd := exec.Command("py", "task-22/task.py", fmt.Sprintf("%d", a), fmt.Sprintf("%d", b))
	cmd.Stdout = bufio.NewWriter(os.Stdout)
	err := cmd.Run()
	if err != nil {
		log.Fatalf("Error exec: %s", err.Error())
	}
}



func main() {
	Numbers(40417401740, 5724086204870)
}
