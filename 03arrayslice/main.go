package main

import "fmt"

var (
	arr = [5]int{1,2,3,4,5}
)

func main() {
	var arr [5]int

	arr[0] = 1
	arr[1] = 2
	arr[2] = 3
	arr[3] = 4
	fmt.Println(arr)
	// arr[3] = 3
	// fmt.Println(arr)

	
	// slice
	
	s1 := arr[0:4]
	fmt.Println("copy: ", s1)
	s2 := s1


	s2[2] = 10
	fmt.Println("copy s2: ", s2)
	fmt.Println("arr: ", arr)

	testSlice(&s2, 15)
	fmt.Println("testSlice: ", s2)

	s3 := append(arr[0:2], arr[3:]...)
	fmt.Println("s3: ", s3)


}

func testSlice(s *[]int, v int) {
	*s = append(*s, v)
	fmt.Println("testSlice internal: ", s)
}
