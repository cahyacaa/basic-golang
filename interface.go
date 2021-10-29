package main

import "fmt"

type HasName interface {
	GetName() string
}

//interface kosong
func testInterfaceNull(inp int) interface{} {
	if inp == 1 {
		return inp
	} else {
		return "Bukan 1"
	}
}

func sayHello(hasName HasName) {
	fmt.Println("Hello Nama Saya", hasName.GetName())
}

type Person struct {
	Name string
}

func (person Person) GetName() string {
	return person.Name
}

func main() {
	cahya := Person{
		Name: "Cahya",
	}
	var testInterfaceNull interface{} = testInterfaceNull(0)
	fmt.Println(testInterfaceNull)
	sayHello(cahya)
}
