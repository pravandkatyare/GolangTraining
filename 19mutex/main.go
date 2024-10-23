package main

import (
	"fmt"
	"sync"
	"time"
)

var counter = 0
var mutex sync.RWMutex
var wg sync.WaitGroup

// func increment(i int) {
// 	defer func() {
// 		// wg.Done()
// 		fmt.Println("wg.Done(): ", i)
// 	}()
// 	mutex.Lock()
// 	defer func() {
// 		mutex.Unlock()
// 		fmt.Println("mutex.Unlock(): ", i)
// 	}()
// 	counter++
// }

// func main() {
// 	expCounter := 1000
// 	for i := 0; i < expCounter; i++ {
// 		// wg.Add(1)
// 		go increment(i)
// 	}

// 	// wg.Wait()

// 	fmt.Println("Expected Counter:", expCounter)
// 	fmt.Println("Actual counter: ", counter)

// 	if expCounter != counter {
// 		fmt.Println("Race Condition")
// 	} else {
// 		fmt.Println("No Race Condition")
// 	}
// }

func main() {
	for i := 0; i < 3; i++ {
		wg.Add(1)
		go readData()
	}
	wg.Add(1)
	go writeData(20)
	wg.Wait()
	fmt.Println("Final Data: ", counter)
}

func readData() {
	defer wg.Done()
	mutex.RLock()
	fmt.Println("Read Data: ", counter)
	mutex.RUnlock()
	time.Sleep(time.Second)
}

func writeData(value int) {
	defer wg.Done()
	mutex.Lock()
	counter = value
	mutex.Unlock()
	time.Sleep(time.Second)
}
