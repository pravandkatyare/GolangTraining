package main

import (
	"errors"
	"fmt"
)

func main() {

	_, err := doSomething()
	if err != nil {
		fmt.Println("Error: ", err)
	}

	// var val int
	// var err1 error
	if val, err1 := doSomething(); err1 != nil {
		fmt.Println("Error: ", err1)
	} else {
		fmt.Println("Val: ", val)
	}

	_, err = divideByZero(5, 0)
	if err != nil {
		fmt.Println("Error: ", err)
	}

	if err := doSomethingCustom(10); err != nil {
		fmt.Println("Error: ", err)
	}

	if err := doSomethingCustom(0); err != nil {
		fmt.Println("Error: ", err.Error())
	}

	if err := doSomethingCustom(20); err != nil {
		fmt.Println("Error: ", err.Error())
	}
}

func doSomethingCustom(n int) error {
	if n <= 0 {
		return errors.New("argument must be greater than zero")
	} else if n == 10 {
		return &CustomError{Code: 10, Message: "not valid number"}
	}
	return nil
}

type CustomError struct {
	Code    int
	Message string
}

func (e *CustomError) Error() string {
	return fmt.Sprintf("Error: %d, %s", e.Code, e.Message)
}

func divideByZero(a, b int) (int, error) {
	if b == 0 {
		return 0, errors.New("division by zero")
	}
	return a / b, nil
}

func doSomething() (int, error) {
	return 123, errors.New("Something is wrogn")
}
