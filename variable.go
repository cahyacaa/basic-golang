package main

import "fmt"

func main(){
	var firstName = "Cahya"
	var lastName = "Nugroho"

	lastName = "Abel"

	fmt.Printf("Hello %s %s! \n", firstName, lastName)

//type Interference (just can do in func bloc)

familyName:= "Ryu"

fmt.Printf("Hello %s %s your familiy name is %s\n", firstName, lastName, familyName)

//declaration multivariable

var first, second , third string

first, second, third = "cahya", "love", "someone"

// using type interference

// fourth, fifth, sixth := "cahya", "love", "someone"

// seventh:=7

fmt.Printf("Hello %s %s your love name is %s\n", first, second, third)



//cretae var using keyword new

name := new(string)

fmt.Println(name)
fmt.Println(*name)

}