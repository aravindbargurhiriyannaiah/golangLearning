package main

import (
	"fmt"
)

type contactInfo struct {
	email   string
	zip     string
	address string
}

type person struct {
	firstName string
	lastName  string
	contactInfo
}

func main() {
	p := person{
		firstName: "Aravind",
		lastName:  "Bargur",
		contactInfo: contactInfo{
			email: "aravind_bh@yahoo.com",
			zip:   "560085",
		},
	}
	p.update("Deepika")
	p.print()
}

func (pointerToPerson *person) update(firstName string) {
	(*pointerToPerson).firstName = firstName
}

func (p person) print() {
	fmt.Printf("%+v", p)
}
