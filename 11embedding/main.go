package main

import "fmt"

type Shape interface {
	Area() float64
}

type ColouredShape interface {
	Area(string) float64
	Color() string
}

type Rectangle struct {
	Width  float64
	Height float64
}

type ColouredRectangle struct {
	Rectangle
	Colour string
}

func (c ColouredRectangle) Color() string {
	return c.Colour
}

func (r Rectangle) Area() float64 {
	return r.Width * r.Height
}

func main() {
	rect := Rectangle{
		Width:  5,
		Height: 4,
	}

	fmt.Println(rect.Width)
	fmt.Println(rect.Height)
	fmt.Println(rect.Area())

	var shape Shape
	shape = rect

	fmt.Println("Shape Area: ", shape.Area())

	cr := ColouredRectangle{
		Rectangle: rect,
		Colour:    "RED",
	}

	fmt.Println("ColouredRectangle Area", cr.Area())
	fmt.Println("ColouredRectangle Color", cr.Color())

	var cs ColouredShape
	cs = cr

	fmt.Println("ColouredShape Area", cs.Area())
	fmt.Println("ColouredShape Color", cs.Color())
}
