package main

import (
	"fmt"
	"sync"
)

func IncrementAngka() int {
	var angka int = 0
	return angka
}

func main() {
	var mutex sync.Mutex
	angka := 0
	for i := 0; i < 10900; i++ {
		go func() {
			mutex.Lock()
			angka++
			mutex.Unlock()
		}()
	}
	fmt.Println("angka sesudah", angka)

}
