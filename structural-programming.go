package main

import (
	"fmt"
)

func main() {
	fmt.Println("ini ada")
	// time.Sleep(1 * time.Second)
	fmt.Println("selesai")
	for i := 0; i < 100; i++ {
		go func() {
		}()
	}
}
