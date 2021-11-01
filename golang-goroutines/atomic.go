package main

import (
	"sync"
	"sync/atomic"
)

func main() {
	var int int32 = 0
	group := &sync.WaitGroup{}

	for i := 0; i < 1000; i++ {
		go func() {
			group.Add(1)
			for j := 0; j < 1000; j++ {
				atomic.AddInt32(&int, 1)
			}
			group.Done()
		}()
	}
	group.Wait()
	println(int)
}
