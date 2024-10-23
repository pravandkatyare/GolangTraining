package main

import (
	"fmt"
	"sync"
	"time"
)

var wg sync.WaitGroup

// func main() {
// 	ch := make(chan int)
// 	go func() {
// 		for i := 0; i < 10; i++ {
// 			ch <- i
// 		}
// 		close(ch)
// 	}()

// 	for val := range ch {
// 		fmt.Println("val:", val)
// 	}
// }

// unbuffered channels
// func worker(i int, ch chan<- string) {
// 	defer wg.Done()
// 	time.Sleep(time.Second)
// 	ch <- fmt.Sprintf("worked %d finished", i)
// }

// func print(ch <-chan string) {
// 	defer wg.Done()

// 	fmt.Println("msg: ", <-ch)
// }

// func main() {
// 	ch := make(chan string)

// 	for i := 0; i < 3; i++ {
// 		wg.Add(1)
// 		go worker(i, ch)
// 	}
	
// 	for i := 0; i < 3; i++ {
// 		wg.Add(1)
// 		go print(ch)
// 	}
// 	wg.Wait()
// }


