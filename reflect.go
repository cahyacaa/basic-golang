package main

import (
	"fmt"
	"reflect"
)

type Sample struct {
	Name string `required:"true"`
}

func main() {
	sample := Sample{
		"cahya",
	}

	var data = reflect.TypeOf(sample)
	fmt.Println(data.NumField())
}
