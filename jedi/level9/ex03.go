package main

import (
	"fmt"
	"runtime"
	"sync"
)

// Using goroutines, create an incrementer program
// You can prove that there is a race condition by using the `--race` flag
func main() {
	var wg sync.WaitGroup

	// have a variable to hold the incrementer value
	incrementer := 0
	gs := 100
	wg.Add(gs)

	// launch a bunch of goroutines
	for i := 0; i < gs; i++ {

		// each goroutine should
		go func() {
			// need the incrementer value - store it in a new variable
			v := incrementer

			// yield the processor with runtime.Gosched()
			// Gosched yiilds the processor, allowing other goroutines to run.
			// It does not suspend the current goroutine, so execution resumes automatically.
			runtime.Gosched()

			// increment the new variable
			v++

			// write the value in th e new variable back to the incrementer variable
			incrementer = v
			wg.Done()
		}()
	}
	wg.Wait()

	fmt.Println("incrementer :", incrementer)
}
