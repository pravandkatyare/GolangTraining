package main

import (
	"fmt"
	"math"
)

type Shape interface {
	Area() float64
}

type Rect struct {
	w float64
	h float64
}

func (r Rect) Area() float64 {
	return r.w * r.h
}

type Circle struct {
	r float64
}

func (c Circle) Area() float64 {
	return math.Pi * c.r * c.r
}

func main() {
	printArea(Rect{4, 5})
	printArea(Circle{5})
	// fmt.Println("shape.Area(): ", shape.Area())
}

func printArea(shape Shape) {
	fmt.Println("shape.Area(): ", shape.Area())
}
