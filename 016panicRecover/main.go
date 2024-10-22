package main

import "fmt"

func main() {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recover: ", r)
		}
	}()
	divide(5, 0)
	fmt.Println("Testing recover 1")
	
	fmt.Println("Testing recover 2")
}

func divide(a, b int) int {
	if b == 0 {
		panic("divide by zero")
	}
	return a / b
}
