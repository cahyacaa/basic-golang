package main

import "fmt"

type School struct {
	Name    string
	Address string
}

type Profile struct {
	Name, Gender, Country string
}

func changeSchoolName(name string, school School) School {
	school.Name = name
	return school
}

func (school *School) changeSchoolData(field, value string) {
	if field == "Name" {
		school.Name = value
	} else if field == "Address" {
		school.Address = value
	}
}

func changeCountryToIndonesia(profile *Profile) {
	profile.Country = "Indonesia"
}

func main() {
	school := School{
		Name:    "SMANSA",
		Address: "SAMPIT",
	}
	result := changeSchoolName("SMANDA", school)
	school2 := School{"MAN", "JOGJA"}
	school2.changeSchoolData("Name", "MTS")
	profile := Profile{
		Name:    "Hoki",
		Gender:  "Male",
		Country: "Hongkong",
	}
	fmt.Println(school2)
	fmt.Println(result)
	fmt.Println(profile)
	changeCountryToIndonesia(&profile)
	fmt.Println(profile)
}
