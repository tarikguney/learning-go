package main

import "fmt"
import "time"

func main() {
	naturals := make(chan int)
	squares := make(chan int)

	// Experimenting channels between goroutines
	go func() {
		for x := 0; ; x++ {
			naturals <- x
		}
	}()

	go func() {
		for {
			x := <-naturals
			squares <- x * x
		}
	}()

	for {
		fmt.Println(<-squares)
		time.Sleep(1 * time.Second)
	}
}
