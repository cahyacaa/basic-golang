package main

import (
	"fmt"
	"sync"
)

func runAsync(num int, group *sync.WaitGroup) int {
	defer group.Done()
	group.Add(1)
	result := 1 + num
	fmt.Println("hasil", result)

	return result

}

func main() {
	group := &sync.WaitGroup{}
	for i := 0; i < 10; i++ {
		res := runAsync(i, group)
		fmt.Println(res)
	}

	group.Wait()
	fmt.Println("selesai")
}
