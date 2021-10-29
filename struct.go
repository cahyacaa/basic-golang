package main

import "fmt"

type Profiles struct {
	Name, Address, Phone string
	Gender               int
}

func main() {
	var user Profiles
	user.Name = "cahya"
	user1 := Profiles{
		Name:    "cahya",
		Address: "Sampit",
		Phone:   "08132432432",
		Gender:  1,
	}

	fmt.Println(user1)
	fmt.Println(user)
}
