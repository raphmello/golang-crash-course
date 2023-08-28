package main

import "fmt"

type person struct {
	firstName string
	lastName  string
	contact   contactInfo
}

type contactInfo struct {
	email   string
	zipCode int
}

// another way to embedding a struct
type person2 struct {
	firstName string
	lastName  string
	contactInfo
}

func main() {
	//declaring a struct
	raphael := person{firstName: "Raphael", lastName: "de Mello"}
	fmt.Println(raphael)

	//another way of declaring a struct
	var raphael2 person
	fmt.Println(raphael2)
	fmt.Printf("%+v", raphael2)
	fmt.Println()

	//assigning values to struct fields
	raphael2.firstName = "Raphael"
	raphael2.lastName = "de Mello"
	fmt.Printf("%+v", raphael2)
	fmt.Println()

	jim := person{
		firstName: "Jim",
		lastName:  "Party",
		contact: contactInfo{
			email:   "jim@jim.com",
			zipCode: 94000,
		},
	}
	fmt.Printf("%+v", jim)
	fmt.Println()

	jim2 := person2{
		firstName: "Jim",
		lastName:  "Party",
		contactInfo: contactInfo{
			email:   "jim@jim.com",
			zipCode: 94000,
		},
	}
	fmt.Printf("%+v", jim2)
	fmt.Println()

	//using receiver function and pointer
	jim2.print()

	////calling with a pointer person
	//raphaelPointer := &raphael
	//raphaelPointer.updateName("Rapha")

	////calling with the person
	raphael.updateName("Rapha")

	raphael.print()
}

func (p person) print() {
	fmt.Printf("%+v", p)
	fmt.Println()
}

func (p person2) print() {
	fmt.Printf("%+v", p)
	fmt.Println()
}

func (pointerToPerson *person) updateName(newFirstName string) {
	pointerToPerson.firstName = newFirstName
}
