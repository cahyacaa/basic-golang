package main

import (
	"context"
	"fmt"
	"runtime"
	"time"
)

//try goroutines leak

func CreateCounter() chan int {
	output := make(chan int)
	go func() {
		defer close(output)
		counter := 1
		for {
			output <- counter
			counter++
		}
	}()
	return output
}

func CreateCounterWithContex(ctx context.Context) chan int {
	output := make(chan int)
	go func() {
		defer close(output)
		counter := 1
		for {
			select {
			case <-ctx.Done():
				return
			default:
				output <- counter
				counter++
				time.Sleep(1 * time.Second)
			}
		}
	}()
	return output
}

func main() {
	//go routines leak
	fmt.Println("Total Go Routines", runtime.NumGoroutine())
	destination := CreateCounter()
	for n := range destination {
		fmt.Println("counter", n)
		if n == 10 {
			break
		}
	}
	fmt.Println("Total Go Routines", runtime.NumGoroutine())

	//fix leak issue with contex signal
	parent := context.Background()
	// ctx, cancel := context.WithCancel(parent)
	ctx, cancel := context.WithTimeout(parent, 5*time.Second)
	fmt.Println("Total Go Routines", runtime.NumGoroutine())
	destination1 := CreateCounterWithContex(ctx)
	for n := range destination1 {
		fmt.Println("counter", n)
		if n == 10 {
			break
		}
	}
	cancel()
	// time.Sleep(3 * time.Second)
	fmt.Println("Total Go Routines", runtime.NumGoroutine())

}
