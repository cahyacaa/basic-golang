package main

import (
	"errors"
	"fmt"
	"os"
)

func generateErr1() error {
	return errors.New("lala")
}

func generateErr2() error {
	return errors.New("lala-2")
}

func main() {
	f, err := os.Open("name")
	if err != nil {
		fmt.Errorf(err.Error())
	}
	d, err := f.Stat()
	if err != nil {
		f.Close()
		fmt.Errorf(err.Error(), d)
	}

}
