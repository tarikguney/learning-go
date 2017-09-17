package main

import (
	"fmt"
	"sync"
)

func main() {
	sizes := make(chan int)
	var group sync.WaitGroup
	// we are creating new go routines but we don't use the cap in the main go routine
	// for-loop at the very bottom. So, we need a mechanism to close the channel when semantically
	// there is no more value coming. Otherwise, a deadlock will occur since no one is writing to
	// sizes but the bottom loop still is blocked.
	for i := 0; i < 100; i++ {
		group.Add(1)
		// the number is for capturing an otherwise a closure value.
		go func(number int) {
			defer group.Done()
			sizes <- number
		}(i)
	}

	// not using group.Wait() and close(sizes), then the bottom loop will not know
	// if there is not any more data coming so will be blocked indefinitely.
	go func() {
		group.Wait()
		close(sizes)
	}()

	for size := range sizes {
		fmt.Println(size)
	}

}
