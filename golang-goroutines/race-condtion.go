package main

import (
	"fmt"
	"sync"
)

func main() {
	x := 0
	var mutex sync.Mutex
	for i := 0; i <= 1000; i++ {
		go func() {
			for y := 0; y <= 100; y++ {
				mutex.Lock()
				x = x + 1
			}
			mutex.Unlock()

		}()
	}
	fmt.Println(x)
}
