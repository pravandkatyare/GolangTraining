package main

import (
	"fmt"
	"os"
	"os/signal"
	"syscall"
	"time"
)

func main() {
	sigs := make(chan os.Signal, 1)

	signal.Notify(sigs, syscall.SIGINT, syscall.SIGTERM)

	// signal.Ignore(syscall.SIGTERM)
	go func() {
		for sig := range sigs {
			switch sig {
			case syscall.SIGINT:
				fmt.Println("Received SIGINT (Ctrl+C). Exiting...")
				os.Exit(0)
			case syscall.SIGTERM:
				fmt.Println("Received SIGTERM. Performing cleanup...")
				// Cleanup code here
				os.Exit(0)
			case syscall.SIGHUP:
				fmt.Println("Received SIGHUP. Reloading configuration...")
				// Reload configuration code here
			}
		}
	}()

	fmt.Println("ctrl+C to exit...")
	for {
		time.Sleep(time.Second)
	}
}
