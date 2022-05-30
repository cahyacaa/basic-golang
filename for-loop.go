package main

import "fmt"

func main() {
	sekumpulan := [...]int{1, 2, 3, 4, 4, 3, 3}
	for index, number := range sekumpulan {
		fmt.Printf("%d adalah index %d \n", number, index)
	}
}
