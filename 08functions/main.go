package main

import "fmt"

func main() {
	// _, _, _, _ = personFunc(nil)

	// n := []int{1, 2, 3, 4, 5}
	// variadicFunc(n...)

	// // Anonymous function
	// func(n int) {
	// 	fmt.Println("Anonymous function: ", n)
	// }(10)

	// //closure functions, it holds the context of execution
	// makeCounter := func() func() int {
	// 	count := 0
	// 	return func() int {
	// 		count++
	// 		return count
	// 	}
	// }

	// counter := makeCounter()
	// fmt.Println("counter: ", counter())
	// fmt.Println("counter: ", counter())

	// counter1 := makeCounter()
	// fmt.Println("counter1: ", counter1())
	// fmt.Println("counter1: ", counter1())
	// fmt.Println("counter1: ", counter1())

	// fmt.Println("counter: ", counter())

	// callback function
	num := []int{1, 2, 3, 4, 5}

	double := func(n int) int {
		return n * 2
	}

	square := func(n int) int {
		return n * n
	}

	doubledNum := processNumbers(num, double)
	fmt.Println("doubledNum: ", doubledNum)

	squaredNum := processNumbers(num, square)
	fmt.Println("squaredNum: ", squaredNum)
}

func processNumbers(num []int, callback func(int) int) []int {
	res := make([]int, len(num))
	for idx, val := range num {
		res[idx] = callback(val)
	}
	return res
}

func personFunc(p interface{}) (interface{}, error, error, error) {

	return p, nil, nil, nil
}

func variadicFunc(n ...int) {
	for key, val := range n {
		fmt.Println("key: ", key, " value: ", val)
	}
}
