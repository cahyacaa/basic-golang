package main

import (
	"fmt"
	"testing"
	"time"
)

func RunHelloWorld() {
	fmt.Println("Hello ini proses 1")
}

func DisplayNumber(number int) {
	fmt.Println(number)
}

func TestLoopingGoRoutines(t *testing.T) {
	for i := 0; i < 10000; i++ {
		go DisplayNumber(i)
	}
}
func TestCreateGoRoutines(t *testing.T) {
	go RunHelloWorld()
	fmt.Println("Process Beriktnya")
	time.Sleep(2 * time.Second)
}
