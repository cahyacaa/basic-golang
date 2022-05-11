package main

import (
	"encoding/json"
	"fmt"
)

type User struct {
	FullName       string           `json:"name" json:"required"`
	Age            int              `json:"age"`
	WorkExperience []WorkExperience `json:"workExperience"`
}

type WorkExperience struct {
	Company string
	Year    int
}

func main() {
	var data = `[{"name": "cahya", "age":10}]`
	var object = []User{
		{
			FullName: "cahya nugroho",
			Age:      25,
			WorkExperience: []WorkExperience{{
				Company: "Skyshi",
				Year:    2022,
			}},
		},
		{
			FullName: "Andi Morana",
			Age:      20,
		},
	}

	user := []User{}

	var jsonBytes = []byte(data)

	errUnmarshal := json.Unmarshal(jsonBytes, &user)

	if errUnmarshal != nil {
		panic(errUnmarshal.Error())
	}

	fmt.Println(user[0].FullName)

	var jsonData, err = json.Marshal(object)
	if err != nil {
		fmt.Println(err.Error())
		return
	}

	var jsonString = string(jsonData)
	fmt.Println(jsonString)
}
