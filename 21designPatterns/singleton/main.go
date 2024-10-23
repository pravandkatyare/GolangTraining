package main

import (
	"fmt"
	"sync"
)

var db *DB
var once sync.Once

// singleton with init()
// func init() {
// 	db := &DB{}
// 	fmt.Println("DB", db)
// }

// // constructor
// func GetDB() *DB {
// 	return db
// }

// Service
type DB struct{}

// singleton with Once.Do

func GetDB() *DB {
	once.Do(func() {
		db = &DB{}
		fmt.Println("Invoked")
	})
	return db
}

func main() {
	fmt.Println(GetDB())
	fmt.Println(GetDB())
	fmt.Println(GetDB())
}
