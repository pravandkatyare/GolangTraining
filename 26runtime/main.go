package main

import (
	"fmt"
	"runtime"
)

func main() {
	// go func() {
	// 	time.Sleep(5)
	// }()
	fmt.Println("runtime.NumGoroutine(): ", runtime.NumGoroutine())

	var m runtime.MemStats
	runtime.ReadMemStats(&m)
	fmt.Println("m.Alloc: ", bToMb(m.Alloc))

	fmt.Println("runtime.Version(): ", runtime.Version())
	// fmt.Println("runtime.GC(): ", runtime.GC())
	fmt.Println("runtime.GOOS: ", runtime.GOOS)
	fmt.Println("runtime.GOARCH: ", runtime.GOARCH)
	fmt.Println("runtime.NumCPU(): ", runtime.NumCPU())
	fmt.Println("runtime.GOMAXPROCS(): ", runtime.GOMAXPROCS(4))
	fmt.Println("runtime.GOMAXPROCS(): ", runtime.GOMAXPROCS(8))
	
	// stack trace
	buf := make([]byte, 1<<16)
	stackSize := runtime.Stack(buf, true)
	fmt.Printf("Stack trace:\n%s", buf[:stackSize])
}

func bToMb(b uint64) uint64 {
	return b / 1024 / 1024
}
