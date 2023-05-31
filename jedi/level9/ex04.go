package main

import (
	"fmt"
	"sync"
)

// Using goroutines, create an incrementer program
// You can prove that there is no race condition by using the `--race` flag
func main() {
	var wg sync.WaitGroup

	// have a variable to hold the incrementer value
	incrementer := 0
	gs := 100
	wg.Add(gs)

	// Create Mutex.
	var mu sync.Mutex

	// launch a bunch of goroutines
	for i := 0; i < gs; i++ {

		// each goroutine should
		go func() {
			// lock
			mu.Lock()

			// need the incrementer value - store it in a new variable
			v := incrementer
			// increment the new variable
			v++

			// write the value in th e new variable back to the incrementer variable
			incrementer = v

			// unlock
			mu.Unlock()
			wg.Done()
		}()
	}
	wg.Wait()

	fmt.Println("incrementer :", incrementer) // "incrementer : 100"
}
