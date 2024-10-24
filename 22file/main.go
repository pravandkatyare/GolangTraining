package main

import (
	"fmt"
	"os"
)

// writing to file
func main() {
	file, err := os.Create("output.txt")
	if err != nil {
		fmt.Println("Error: ", err)
		return
	}
	defer file.Close()

	_, err = file.WriteString("Hello from Golang training...!!")
	if err != nil {
		fmt.Println("Error writing into file: ", err)
		return
	}
	// fmt.Println("temp: ", temp)

	data := []byte("This is byte slice")

	temp, err := file.Write(data)
	if err != nil {
		fmt.Println("Error writing byte slice: ", err)
	}

	fmt.Println("temp: ", temp)
}

// reading from file
// func main() {
// 	file, err := os.Open("output.txt")
// 	if err != nil {
// 		fmt.Println("error opening: ", err)
// 		return
// 	}

// 	data, err := io.ReadAll(file)
// 	if err != nil {
// 		fmt.Println("Error reading file: ", err)
// 		return
// 	}

// 	fmt.Println("Data from file: ", string(data))
// }
