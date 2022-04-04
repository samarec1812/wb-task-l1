package main

/*
	Реализовать собственную функцию sleep.
*/

import (
	"fmt"
	"os/exec"
	_ "os/signal"
	"time"
)



// #include <unistd.h>
// void callSleep(unsigned int seconds) {
//  sleep(seconds);
//}
import "C"

// Sleep вызывает системную команду sleep
func Sleep(seconds uint) {
	cmd := exec.Command("sleep", fmt.Sprintf("%d", seconds))
	cmd.Run()
}


// SleepB вызывает си-ую функцию sleep
func SleepB(seconds uint) {
	C.callSleep(C.uint(seconds))
}


//func SleepC(seconds uint) {
//	bin, _ := exec.LookPath("sleep")
//	args := []string{"sleep", fmt.Sprintf("%d", seconds)}
//	env := os.Environ()
//	err := syscall.Exec(bin, args, env)
//	if err != nil {
//		log.Fatalf("Error syscall: %s", err.Error())
//	}
//}

func SleepD(seconds time.Duration) {
	<- time.After(seconds * time.Second)
}

func main() {

	fmt.Println("Hello")
	SleepD(20)
	fmt.Println("World!")
}