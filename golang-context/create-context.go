package main

import (
	"context"
	"fmt"
)

func main() {
	contextA := context.Background()
	contextB := context.WithValue(contextA, "key1", "contextB")
	contextC := context.WithValue(contextB, "key", "contextC")

	fmt.Println(contextA, contextB, contextC.Value("key"), contextC.Value("key1"))
}
