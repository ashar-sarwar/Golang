package main

import (
	"fmt"
	"reflect"
)

type Doctor struct {
	number     int
	actorName  string
	companions []string
}

type Animal struct {
	Name   string `required max:"100"`
	Origin string
}

type Bird struct {
	Animal
	speed float32
	fly   bool
}

func pain() {
	aDoctor := Doctor{
		number:    3,
		actorName: "kskja",
		companions: []string{
			"yhuds",
			"iud9w",
		},
	}
	fmt.Println(aDoctor)           // {3 kskja [yhuds iud9w]}
	fmt.Println(aDoctorcompanions) // [yhuds iud9w]

	b := Bird{
		Animal: Animal{Name: "Emu", Origin: "asu"},
		speed:  48,
		fly:    true,
	}
	fmt.Println(b.Name)

	t := reflect.TypeOf(Animal{})
	feild, _ := t.FieldByName("Name")
	fmt.Println(feild.Tag) // required max:"100"

}
  