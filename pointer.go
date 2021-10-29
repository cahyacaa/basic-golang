package main

import (
	"fmt"
)

type DetailBarang struct {
	Name string
	Type string
}

func main() {
	var numberA int = 190
	var numberB *int = &numberA
	fmt.Println("numberA (value)   :", numberA)  // 4
	fmt.Println("numberA (address) :", &numberA) // 0xc20800a220

	fmt.Println("numberB (value)   :", *numberB) // 4
	fmt.Println("numberB (address) :", numberB)  // 0xc20800a220

	barangA := DetailBarang{
		Name: "Nokia",
		Type: "HandPhone",
	}

	barangB := &barangA
	barangB.Name = "samsung"

	fmt.Println(barangA, barangB)
}
