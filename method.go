package main 

import (
	"fmt"
	"strings"
)

type student struct {
	name string
	age int
}

func (s student) sayHello(){
	fmt.Println("hallo", s.name)
}

func (s student) getNameAt(i int) string {
    return strings.Split(s.name, " ")[i-1]
}

type drakor struct{
	Nama string
	Rating string
	Genre string
}

func(d drakor) detailDrama(){
	fmt.Printf("Judul Drama :%s, Rating : %s, Genre : %s \n", d.Nama, d.Rating, d.Genre)
}

func main(){
	var s1 = student{"Cahya Nugroho", 24}
	s1.sayHello()

	drama1:= drakor{"Reply 1988", "17.84%" ,"Comedy, Romance"}
	drama1.detailDrama()

	var name = s1.getNameAt(2)
    fmt.Println("nama panggilan :", name)
}

