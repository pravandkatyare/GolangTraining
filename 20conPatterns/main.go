package main

import (
	"context"
	"fmt"
	"time"
)

// fan-out pattern
// func main() {

// 	numJobs := 10

// 	jobCh := make(chan int, numJobs)
// 	resultCh := make(chan int, numJobs)

// 	for w := 0; w < 3; w++ {
// 		go worker(w, jobCh, resultCh)
// 	}

// 	for i := 0; i < numJobs; i++ {
// 		jobCh <- i
// 	}

// 	for i := 0; i < numJobs; i++ {
// 		// fmt.Println("Result: ", <-resultCh)
// 		<-resultCh
// 	}
// }

// func worker(name int, jobCh, resultCh chan int) {
// 	for job := range jobCh {
// 		fmt.Println("Worker: ", name, "processing job: ", job)
// 		time.Sleep(time.Millisecond * 200)
// 		resultCh <- job * 2
// 	}
// }

// fan-in pattern & signal channel

// func main() {
// 	sigCh := make(chan bool)
// 	output := make(chan int)
// 	input1 := make(chan int)
// 	input2 := make(chan int)
// 	input3 := make(chan int)

// 	go producer(input1)
// 	go producer(input2)
// 	go producer(input3)
// 	go fanIn(sigCh, input1, input2, input3, output)

// 	go func() {
// 		sigCh <- true
// 		time.Sleep(time.Second * 10)
// 	}()
// 	for n := range output {
// 		fmt.Println("n: ", n)
// 		time.Sleep(time.Second)
// 	}
// }

// func fanIn(sigCh chan bool, input1, input2, input3, output chan int) {
// 	for {
// 		select {
// 		case in1, ok := <-input1:
// 			if ok {
// 				output <- in1
// 			}
// 		case in2, ok := <-input2:
// 			if ok {
// 				output <- in2
// 			}
// 		case in3, ok := <-input3:
// 			if ok {
// 				output <- in3
// 			}
// 		case <-sigCh:
// 			fmt.Println("Signal chan")
// 			close(sigCh)
// 			close(output)

// 			return
// 		}
// 	}
// }

// func producer(inputCh chan int) {
// 	defer close(inputCh)
// 	for i := 0; i < 3; i++ {
// 		inputCh <- i
// 	}
// }

// context and cancellation

func main() {
	// ctx, cancel := context.WithCancel(context.Background())
	ctx, _ := context.WithTimeout(context.Background(), time.Duration(time.Second*5))
	// sigCh := make(chan bool)
	output := make(chan int)
	input1 := make(chan int)
	input2 := make(chan int)
	input3 := make(chan int)

	go producer(input1)
	go producer(input2)
	go producer(input3)
	go fanIn(ctx, input1, input2, input3, output)

	go func() {
		time.Sleep(time.Second * 10)
		// cancel()
	}()
	for n := range output {
		fmt.Println("n: ", n)
		time.Sleep(time.Second)
	}
}

func fanIn(ctx context.Context, input1, input2, input3, output chan int) {
	for {
		select {
		case in1, ok := <-input1:
			if ok {
				output <- in1
			}
		case in2, ok := <-input2:
			if ok {
				output <- in2
			}
		case in3, ok := <-input3:
			if ok {
				output <- in3
			}
		case <-time.After(time.Duration(time.Second)):
			fmt.Println("Waiting....")
			// case <-ctx.Done():
			// 	fmt.Println("ctx.Done()")
			// 	close(output)
			// 	return
		}
	}
}

func producer(inputCh chan int) {
	defer close(inputCh)
	for i := 0; i < 3; i++ {
		inputCh <- i
	}
}
