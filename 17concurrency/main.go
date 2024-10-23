package main

import (
	"fmt"
	"sync"
	"time"
)

// var wg sync.WaitGroup

// func main() {
// 	for i := 0; i < 10; i++ {
// 		wg.Add(1)
// 		go print(i)
// 	}

// 	wg.Wait()
// }

// func print(idx int) {
// 	defer wg.Done()
// 	for i := 0; i < 10; i++ {
// 		fmt.Println("idx: ",idx," i: ",  i)
// 	}

// }

func doWork(id int, wg *sync.WaitGroup) {
	defer wg.Done() // Decrement the counter when the goroutine completes
	fmt.Printf("Goroutine %d is starting\n", id)
	time.Sleep(time.Second) // Simulate work
	fmt.Printf("Goroutine %d is done\n", id)
}

func main() {
	var wg sync.WaitGroup

	const numGoroutines = 5

	// Launching goroutines
	for i := 1; i <= numGoroutines; i++ {
		wg.Add(1)         // Increment the WaitGroup counter
		go doWork(i, &wg) // Start a new goroutine
	}

	wg.Wait() // Wait for all goroutines to finish
	fmt.Println("All goroutines have completed.")
}
