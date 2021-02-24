package main

import (
    "fmt"
    "runtime"
    "sync"
)

type counter struct {
    val int
    sync.Mutex
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
    m := make(map[int]int)
    var wg sync.WaitGroup
    // var meter counter
    var mtx sync.Mutex
    wg.Add(10)
    for i := 0; i < 10; i++ {
      

        go func() {
                mtx.Lock()
                m[i] += i
                mtx.Unlock()
                fmt.Println(m[i])
            wg.Done()
        }()
    }

    wg.Wait()
    println(len(m))
}