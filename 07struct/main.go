package main

import "fmt"

type Person struct {
	Name string
	Age  int
}

func main() {
	var p1 Person         //holding an object
	var nilP *Person      // nil object without any memory address
	person := new(Person) // it assigns memory address

	nilP = &p1
	nilP.Name = "Bar"
	nilP.Age = 22

	fmt.Println("p: ", &nilP)

	if person == nil {
		fmt.Println("nil? : ")
	}

	person.Name = "Foo"
	person.Age = 20

	fmt.Println(person)

	

	
}
