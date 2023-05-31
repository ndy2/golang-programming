package main

import (
	"fmt"
	"sync"
	"sync/atomic"
)

// Using goroutines, create an incrementer program
// You can prove that there is no race condition by using the `--race` flag
func main() {
	var wg sync.WaitGroup

	// have a variable to hold the incrementer value
	var incrementer int64
	gs := 100
	wg.Add(gs)

	// launch a bunch of goroutines
	for i := 0; i < gs; i++ {

		// each goroutine should
		go func() {
			atomic.AddInt64(&incrementer, 1)
			wg.Done()
		}()
	}
	wg.Wait()

	fmt.Println("incrementer :", incrementer) // "incrementer : 100"
}
