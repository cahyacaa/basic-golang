package main

import "fmt"

func main() {
	counter := 0
	incrementResult := func() {
		fmt.Println("Incrementing", counter)
		counter++
	}
	for i := 0; i < 100; i++ {
		fmt.Println(i)
		incrementResult()
	}
}
