package main

import (
	"fmt"
)

func main() {
	var m = make(map[int]string)

	m[1] = "abc"
	//add
	m[2] = "xyz"

	//update
	m[1] = "pqr"

	// get
	fmt.Println(m[2])

	// check if present
	_, ok := m[3]
	if !ok {
		fmt.Println("not present")
	}

	// delete element
	delete(m, 2)

	fmt.Println(m)
}
