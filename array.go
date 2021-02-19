package main 

import "fmt"

func main(){
	var words[4] string

	words[0]="I"
	words[1]="love"
	words[2]="U"
	words[3]="SomeOne"

	fmt.Println(words[0], words[1], words[2], words[3])
	fmt.Println(words)

	var fruits = [4]string{"apple", "grape", "banana", "melon"}
	//vertical style
	fruits  = [4]string{
		"apple",
		"grape",
		"banana",
		"melon",
	}

	fmt.Println("Jumlah element ", len(fruits))
	fmt.Println("Isi semua element \t", fruits)

	var numbers = [...]int{2, 3, 2, 4, 3}

	fmt.Println("data array \t:", numbers)
	fmt.Println("jumlah elemen \t:", len(numbers))

	var numbers1 = [2][3]int{[3]int{3, 2, 3}, [3]int{3, 4, 5}}
	var numbers2 = [2][3]int{{3, 2}, {3, 4, 5}}

	fmt.Println("numbers1", numbers1)
	fmt.Println("numbers2", numbers2)

	var fruit = [4]string{"apple", "grape", "banana", "melon"}

	for i, fruit := range fruit {
    	fmt.Printf("elemen %d : %s\n", i, fruit)

		var fruits = [4]string{"apple", "grape", "banana", "melon"}

		for i, fruit := range fruits {
			fmt.Printf("elemen %d : %s\n", i, fruit)
		}

		var car = make([]string, 2)
		car[0] = "BMW"
		car[1] = "Merceddes"
		fmt.Println(car	) 
}
}