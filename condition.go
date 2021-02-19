package main

import "fmt"

func main(){
	point:=10
	if point >=9{
		fmt.Println("value greater than initial value, your value is",point)
	}
//repeating
	i:=1
	for i := 0; i < 5; i++ {
		fmt.Println("Number", i)
	}

	for i < 5 {
		fmt.Println("Angka", i)
		i++
	}
	for i := 0; i < 5; i++ {
		for j := i; j < 5; j++ {
			fmt.Print(j)
		}
	}
	outerLoop:
	for i := 0; i < 5; i++ {
   	 	for j := 0; j < 5; j++ {
        	if i == 3 {
            	break outerLoop
        }
        	fmt.Print("matriks [", i, "][", j, "]", "\n")
    }
	
}

}