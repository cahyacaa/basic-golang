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

func main(){
	var s1 = student{"Cahya Nugroho", 24}
	s1.sayHello()

	var name = s1.getNameAt(2)
    fmt.Println("nama panggilan :", name)
}

