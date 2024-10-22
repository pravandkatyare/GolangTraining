package main

import "fmt"

type Color int

const (
	RED Color = iota
	GREEN
	BLUE
)

func (c Color) String() string {
	// return []string{"Red", "GREEN", "BLUE"}[c]
	colorMap := map[Color]string {
		RED: "Red",
		GREEN: "Green",
		BLUE: "Blue",
	}
	return colorMap[c]
}

func main() {
	var c Color
	c = GREEN
	switch c {
	case RED:
		fmt.Println(RED)
	case GREEN:
		fmt.Println(GREEN)
	case BLUE:
		fmt.Println(BLUE)
	}
}
