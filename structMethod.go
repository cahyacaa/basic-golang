package main

import "fmt"

type Customer struct {
	Name string
}

func sayHi(customer Customer, name string) {
	fmt.Println("hi", name, "my name", customer.Name)
}

func (customer Customer) sayHello(name string) {
	fmt.Println("hello", name, "my name", customer.Name)
}
func main() {
	var user Customer
	user.Name = "cahya"
	sayHi(user, "lola")
	user.sayHello("bambang")
}
