package main

import (
    "fmt"
    "math/rand"
    "time"
    "strings"
    "math"
)

func main() {
    var avg = calculate(2, 4, 3, 5, 4, 3, 3, 5, 5, 3)
    var msg = fmt.Sprintf("Average : %.2f", avg)
    fmt.Println(msg)

    names:=[]string{"Jhon","wick"}
    printMessage("hello", names)
    rand.Seed(time.Now().Unix())
    var randomValue int

    randomValue = randomWithRange(2, 10)
    fmt.Println("random number:", randomValue)
    randomValue = randomWithRange(2, 10)
    fmt.Println("random number:", randomValue)
    randomValue = randomWithRange(2, 10)
    fmt.Println("random number:", randomValue)

    dividenumber(10, 2)
    dividenumber(4, 0)
    dividenumber(8, -4)

    var diameter float64 = 15
    var area, circumference = circle(diameter)

    fmt.Printf("luas lingkaran\t\t: %.2f \n", area)
    fmt.Printf("keliling lingkaran\t: %.2f \n", circumference)

    fmt.Println(split(145))
    // even,odd=countEvenandOddVal(90,1000)
    // fmt.Println(even)

    //closure
    var getMinMax = func(n []int) (int, int) {
        var min, max int
        for i, e := range n {
            switch {
            case i == 0:
                max, min = e, e
            case e > max:
                max = e
            case e < min:
                min = e
            }
        }
        return min, max
    }

    var numbers = []int{2, 3, 4, 3, 4, 2, 3}
    var min, max = getMinMax(numbers)
    fmt.Printf("data : %v\nmin  : %v\nmax  : %v\n", numbers, min, max)
}

func calculate(numbers ...int) float64 {
    var total int = 0
    for _, number := range numbers {
        total += number
    }

    var avg = float64(total) / float64(len(numbers))
	fmt.Println(total, len(numbers))
    return avg
}

func printMessage(message string, input[]string){
    nameString := strings.Join(input," ")
    fmt.Println(message, nameString)
}

func randomWithRange(min, max int )int{
    var value = rand.Int()%(max-min +1)+min
    return value
}

func dividenumber (m, n int){
    if n==0{
        fmt.Printf("invalid divider. %d cannot divided by %d\n",m,n)
        return
    }
        var res= m/n
        fmt.Printf("%d / %d = %d\n", m, n, res)
    
}

func circle(d float64) (float64, float64)  {
    // hitung luas
    var area = math.Pi * math.Pow(d / 2, 2)
    // hitung keliling
    var circumference = math.Pi * d

    // kembalikan 2 nilai
    return area, circumference
}

// predefined return value 
// func countEvenandOddVal(startVal, endVal int)(even,odd int){
//     var even=0
//     var odd=0
//     for i:=startVal; i<endVal; i++{
//         if i%2==0{
//             even=+1
//         }else{
//             odd=+1
//         }
//     }
//     return 
// }

func split(sum int) (x, y int) {
	x = sum * 4 / 9
	y = sum - x
	return
}