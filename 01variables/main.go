package main

import (
	"fmt"
)

var (
	test = "test"
	abc  = "abc"
)

const (
	xyz = "xyz"
)

func main() {
	// var str string
	// str = "Hello"
	// fmt.Println(str)
	// str := "Hello"
	// fmt.Println(str)
	// fmt.Printf("%T", str)

	// Rune

	var r rune = '😀'

	fmt.Printf("Rune: %c, Integer value: %d\n", r, r)

	str = "Hello, 🌍🌟🚀"

	for i, c := range str {
		fmt.Printf("Character %d: %c (Rune: %d)\n", i, c, c)
	}

	runes := []rune{'👋', 'H', 'e', 'l', 'l', 'o'}
	fmt.Printf("Runes slice: %c\n", runes)

	runeSlice := []rune(str)
	fmt.Printf("Runes from string: %v\n", runeSlice)

	fmt.Printf("The first character of the string: %c\n", runeSlice[0]) // H
	fmt.Printf("The character at index 7: %c\n", runeSlice[7])          // 🌍
	fmt.Printf("The character at index 8: %c\n", runeSlice[8])          // 🌟
	fmt.Printf("The character at index 9: %c\n", runeSlice[9])          // 🚀

}
