package main

import "fmt"

func studentMap(name, gender string) map[string]string {
	x, y := name, gender
	if x == "" || y == "" {
		return nil
	} else {
		return map[string]string{
			name:   name,
			gender: gender,
		}
	}
}
func main() {
	var student map[string]string = studentMap("cahya", "")
	if student == nil {
		fmt.Println("DATA KOSONG")
	}
	fmt.Println(student)
}
