package main

import (
	"errors"
	"fmt"
)

var (
	wrogn = errors.New("Something is wrogn")
)

type CustomError struct {
	Code    int
	Message string
}

func (e *CustomError) Error() string {
	return fmt.Sprintf("Error: %d, %s", e.Code, e.Message)
}

func main() {
	err := doSomething()
	if err != nil {
		fmt.Println("Error: ", err)
	}

	fmt.Println("errors.Unwrap(err): ", errors.Unwrap(err))

	fmt.Println("errors.Is(err, wrogn): ", errors.Is(err, wrogn))

	err = doSomethingCustom()
	var customErr *CustomError
	fmt.Println("error.As(err, &customErr): ", errors.As(err, &customErr))
}

func doSomethingCustom() error {
	if err := customError(); err != nil {
		return fmt.Errorf("doSomething: %w", err)
	}
	return nil
}

func doSomething() error {
	if err := anotherFunc(); err != nil {
		return fmt.Errorf("doSomething: %w", err)
	}
	return nil
}

func anotherFunc() error {
	return wrogn
}

func customError() error {
	return &CustomError{Code: 10, Message: "Invalid"}
}
