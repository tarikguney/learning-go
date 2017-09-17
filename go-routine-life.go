package main

import (
	"fmt"
	"time"
)

func main() {
	foo()
	time.Sleep(4 * time.Second)
}

func foo() {
	go func() {
		time.Sleep(2 * time.Second)
		fmt.Println("From go routine.")
	}()
	fmt.Println("Exiting boo")
}
