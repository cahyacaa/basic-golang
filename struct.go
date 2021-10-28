package main

import "fmt"

type Profile struct {
	Name, Address, Phone string
	Gender               int
}

func main() {
	var user Profile
	user.Name = "cahya"
	user1 := Profile{
		Name:    "cahya",
		Address: "Sampit",
		Phone:   "08132432432",
		Gender:  1,
	}

	fmt.Println(user1)
	fmt.Println(user)
}
