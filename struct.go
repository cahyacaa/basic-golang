package main 

import (
	"fmt"
)

type student struct {
	name string 
	grade int
	languages
}

type languages struct{
	intermediate []string
}

func main(){
	var s1= student{} 
	s1.name = "Ryu Hye-Young"
	s1.grade = 10
	s1.intermediate = []string{"lol","lal"}
	fmt.Println("name :", s1.name)
	fmt.Println("grade :", s1.grade)
	fmt.Println("languages :", s1.languages.intermediate)
}