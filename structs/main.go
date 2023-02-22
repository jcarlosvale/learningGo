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
	joel := person{
		firstName: "joel",
		lastName:  "ellie",
		contact: contactInfo{
			email:   "test@test.com",
			zipCode: 1223,
		},
	}

	joel.updateName("jimmy")
	joel.print()
}

func (p *person) updateName(newFirstaName string) {
	(*p).firstName = newFirstaName
}

func (p person) print() {
	fmt.Printf("%+v", p)
}
