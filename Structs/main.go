package main

import "fmt"

//import "fmt"
type contactinfo struct {
	email   string
	zipcode int
}
type person struct {
	firstName string
	lastName  string
	contact   contactinfo
}

func main() {
	srini := person{firstName: "Kamatham", lastName: "Srinivasa", contact: contactinfo{email: "kamathams565@gmail.com", zipcode: 75035}}
	//srini.updateName("kamat")
	//sriniPointer := &srini
	//sriniPointer.updateName("kmt")
	srini.updateName("kmt")
	srini.print()
	//fmt.Println(srini)
	/*var srini person
	srini.firstName = "Kamatham"
	srini.lastName = "Srinivasa"
	fmt.Println(srini)
	fmt.Printf("%+v", srini)*/
}
func (p person) print() {
	fmt.Println(p)
}
func (pointerToPerson *person) updateName(newFirstName string) {
	pointerToPerson.firstName = newFirstName
}
