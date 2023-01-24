package main

import (
	"fmt"
	"runtime"
	"sync"
)

type counter struct {
	val int
	*sync.Mutex
}

func (c *counter) Add(x int) {
	c.Lock()
	c.val++
	c.Unlock()
}

func (c *counter) Value() (x int) {
	return c.val
}

func main() {
	runtime.GOMAXPROCS(2)
	var wg sync.WaitGroup

	count := counter{
		val: 1,
	}
	wg.Add(10)
	for i := 0; i < 10; i++ {
		go func() {
			defer wg.Done()
			count.Add(count.val)
		}()
	}

	wg.Wait()
	fmt.Println(count.Value())
}
