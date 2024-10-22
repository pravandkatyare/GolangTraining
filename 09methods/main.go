package main

import "fmt"

type Person struct {
	Name string
	// age  int
}

func (p *Person) getName() string {
	return p.Name
}

func (p *Person) setName(name string) {
	p.Name = name
}

func main() {
	var p Person

	p.setName("FOO")
	fmt.Println("Name: ", p.getName())

}
