package main

import "fmt"

func CreateChannel(channel chan string) {
	fmt.Println("Mengirim Data Ke Channel")
	channel <- "Data"

}

func ChannelNonParameter() {
	channel := make(chan string)
	go func() {
		channel <- "Data1"
	}()
	defer close(channel)
	fmt.Println(<-channel)
}

func CreateSecondChannel(channel chan string) {
	fmt.Println("Mengirim Data Ke Channel")
	channel <- "Data2"
}
func main() {
	channel := make(chan string)
	go CreateChannel(channel)
	go CreateSecondChannel(channel)
	defer close(channel)
	data := <-channel
	fmt.Println(data)
	ChannelNonParameter()

}
