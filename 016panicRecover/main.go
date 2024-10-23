package main // package main

import (
	"fmt"
	"runtime"
)

func main() {

	fmt.Println("Testing recover 1")
	divide(5, 0)
	fmt.Println("Testing recover 2")
	divide(10, 1)
	fmt.Println("Program continues...")
}

func divide(a, b int) int {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recover: ", r)
			buf := make([]byte, 1024)
			stackSize := runtime.Stack(buf, true)
			fmt.Println("Stack trace:\n", string(buf[:stackSize]))
		}
	}()
	if b == 0 {
		panic("divide by zero")
	}
	return a / b
}
