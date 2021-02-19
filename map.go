package main

import "fmt"

func main(){


	var stock= map[string]int{
		"Mouse":1,
		"Car":10,
		"Keybpard":20,
		"LCD Monitor":12,
	}

	fmt.Println(stock)
	var chicken map[string]int
	chicken = map[string]int{}

	chicken["januari"] = 50
	chicken["februari"] = 40
	chicken["mei"]=10

	fmt.Println("januari", chicken["januari"]) // januari 50
	fmt.Println("mei",     chicken["mei"])     // mei 0

	var value, isExist = chicken["mei"]

	if isExist {
		fmt.Println(value)
	} else {
		fmt.Println("item is not exists")
}

}