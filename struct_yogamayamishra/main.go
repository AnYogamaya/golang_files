package main

import "fmt"

type contactInfo struct {
	email   string
	zipCode int
}

type person struct {
	firstName string
	lastName  string
	contact   contactInfo
}

func main() {
	maya := person{
		firstName: "Yogamaya",
		lastName:  "Mishra",
		contact: contactInfo{
			email:   "maya@gmail.com",
			zipCode: 752120,
		},
	}
	mayaPointer := &maya
	mayaPointer.updateName("chinu")
	maya.print()
}

func (p *person) updateName(firstname string) {
	(*p).firstName = firstname
}

func (p person) print() {
	fmt.Printf("%+v", p)
}
