package main

import "fmt"

type Animal struct{}

type Birds struct{}

func main() {
	var m = make(map[string]interface{})
	// var a Animal
	// var b Birds
	m["abc"] = Animal{}
	m["pqr"] = Birds{}

	// fmt.Println(m)

	// for k, v := range m {
	// 	fmt.Println("key:", k, " value: ", v)
	// }

	// while substitute
	// i := 0
	// for ; i < 10; i++ {
	// 	fmt.Println(i)
	// }

	// break
	// for i := 0; i < 10; i++ {
	// 	fmt.Println(i)
	// 	if i == 5 {
	// 		break
	// 	}
	// }

	// continue
	for i := 0; i < 10; i++ {
		if i < 5 {
			continue
		}
		fmt.Println(i)
	}

}
