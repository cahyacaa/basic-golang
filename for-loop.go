package main

import "fmt"

type Users struct {
	name string
	age  int32
}

func main() {
	sekumpulan := [...]int{1, 2, 3, 4, 4, 3, 3}
	for index, number := range sekumpulan {
		fmt.Printf("%d adalah index %d \n", number, index)
	}

	kvs := map[string]string{"a": "apple", "b": "banana"}
	akvs := []Users{{
		name: "cahya",
		age:  12,
	},
		{
			name: "agil",
			age:  19,
		},
	}
	for k, v := range kvs {
		fmt.Printf("%s -> %s\n", k, v)
	}

	for k, v := range akvs {
		fmt.Println(k, v)
	}
}
