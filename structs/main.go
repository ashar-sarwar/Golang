package main

import "fmt"

type contactInfo struct {
	email string
	zip   int
}

type person struct {
	firstName string
	lastName  string
	contact   contactInfo
}

func main() {
	alex := person{
		firstName: "Alex",
		lastName:  "Anderson",
		contact: contactInfo{
			email: "abc@gail.com",
			zip:   354435,
		},
	}
	//fmt.Printf("%+v", alex)

	//	alexPointer := &alex
	//	alexPointer.updateName("jimmy")
	alex.updateName("jimmy")
	alex.print()
}

func (p person) print() {
	fmt.Printf("%+v", p)
}

func (pointerToPerson *person) updateName(newFirstName string) {
	(*pointerToPerson).firstName = newFirstName
}
