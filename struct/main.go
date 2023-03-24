package main

import "fmt"

type contactInfo struct {
	email   string
	zipCode int
}

type person struct {
	firstName string
	lastName  string
	contactInfo
}

func main() {
	myself := person{
		firstName: "David",
		lastName:  "Gagliardi",
		contactInfo: contactInfo{
			email:   "dgags2012@gmail.com",
			zipCode: 55025,
		},
	}

	myself.printName()

	myself.updateName("Johnny", "Willamsberg", &myself.lastName)

	myself.printName()
}

func (p *person) printName() {
	fmt.Printf("%v %v the man the myth the legend\n", p.firstName, p.lastName)
}

func (p *person) updateName(newFirstName string, newLastName string, lastNamePointer *string) {
	p.firstName = newFirstName
	*lastNamePointer = newLastName
}
