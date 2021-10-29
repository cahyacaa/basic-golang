package main

import (
	"errors"
	"fmt"
)

func pembagi(nilaiAtas, nilaiBawah float32) (float32, error) {
	if nilaiAtas == 0 {
		return nilaiAtas, errors.New("Angka Tidak valid")
	} else {
		x := nilaiAtas / nilaiBawah
		return x, nil
	}
}
func main() {
	var contoh = errors.New("dodol kamu")
	fmt.Println(contoh.Error())

	hasil, err := pembagi(10, 3.99)
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println(hasil)
	}
}
