package main

import "fmt"

type Animal struct {
}

func main() {
	// var isLoggedIn bool

	// if !isLoggedIn {
	// 	fmt.Println("Unauthorized user")
	// }

	// if isLoggedIn {
	// 	fmt.Println("Logged In user")
	// }
	switchCondition("RED")
	switchCondition(123)
	// var animal Animal
	// if &animal == nil {
	// 	fmt.Println("nil")
	// }
}

func switchCondition(color interface{}) {
	switch color {
	case "RED":
		fmt.Println("RED")
	case "Green":
		fmt.Println("Green")
	case "Blue":
		fmt.Println("Blue")
	case 123:
		fmt.Println(123)
	default:
		fmt.Println("Invalid")
	}
}
