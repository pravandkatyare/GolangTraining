package main

import (
	"0112accessmodifiers/employee"
	"fmt"
)

func main() {
	var emp employee.Employer
	emp = &employee.Employee{}
	emp.SetName("FOO")
	fmt.Println("Name: ", emp.GetName())

	// var shape rectangle // not able to access because it is private in package shape
}
