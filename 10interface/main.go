package main

import "fmt"

type Employer interface {
	test(string)
}

type Person struct{}

func (p Person) test(str string) {
	fmt.Println(str)
}

func main() {
	var e Employer // zero value interface
	//e = Person{} // assigning an object to an interface will no longer be zero value

	e.test("Foo") // if accessed with assigning any valu it will panic

	var i interface{} // empty interface

	fmt.Println("i: ", i)

	m := map[string]interface{}{
		"command": "dir",
		"timeout": 10,
		"abc":     Person{},
	}

	fmt.Println(m)
}
