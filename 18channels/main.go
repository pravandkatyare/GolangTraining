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

// buffered channel

// func main() {
// 	ch := make(chan int, 10)
// 	wg.Add(1)
// 	go func() {
// 		defer wg.Done()
// 		for i := 0; i < 10; i++ {
// 			ch <- i
// 			fmt.Println("Sent: ", i)
// 		}
// 		close(ch)
// 	}()

// 	wg.Add(1)
// 	go func() {
// 		defer wg.Done()
// 		for val := range ch {
// 			fmt.Println("Received: ", val)
// 		}
// 	}()

// 	wg.Wait()
// }

// select on channel

func main() {
	number := make(chan int)
	msg := make(chan string)

	go channelNumber(number)
	go channelMessage(msg)

	for {
		select {
		case numCh := <-number:
			fmt.Println("Channel Data:", numCh)
		case msgCh := <-msg:
			fmt.Println("Channel Data: ", msgCh)
		default:
			fmt.Println("No Data")
			time.Sleep(time.Second *2)
		}
	}
}

func channelNumber(num chan int) {
	time.Sleep(time.Second)
	num <- 10
}

func channelMessage(msg chan string) {
	msg <- "Select on channel"
}
