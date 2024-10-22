package main

import (
	"fmt"
	"math"
)

type Shape interface {
	Area()
}

type Rect struct{}

func (r Rect) Area() {}

func main() {
	var i interface{}
	i = 123

	str, ok := i.(string)
	fmt.Println("str: ", str)
	fmt.Println("ok: ", ok)

	var shape Shape
	shape = Rect{}  // Rect
	shape = &Rect{} // Shape

	switch s := shape.(type) {
	case Rect:
		fmt.Println("Rect: ", s)
	case Shape:
		fmt.Println("Shape: ", s)
	default:
		fmt.Println("Invalid")
	}

	printType(11)
	printType("hello")
	printType(true)
	printType(math.Pi)
}

func printType(i interface{}) {
	switch v := i.(type) {
	case int:
		fmt.Printf("Type int, Value: %d\n", v)
	case string:
		fmt.Printf("Type string, Value: %s\n", v)
	case bool:
		fmt.Printf("Type bool, Value: %t\n", v)
	default:
		fmt.Printf("Unknown type\n")
	}
}
