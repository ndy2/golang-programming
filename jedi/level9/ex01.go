// In addition to the main goroutine,
// launch two additional goroutines.
// Each additional goroutine should print something out.

// Use waitgroups to make sure each goroutine finishes before your program exit.
package main

import (
	"fmt"
	"runtime"
	"sync"
)

func main() {
	fmt.Println("begin - Goroutines", runtime.NumGoroutine())
	fmt.Println("begin - CPUs", runtime.NumCPU())

	var wg sync.WaitGroup
	wg.Add(2)

	fmt.Println("Hello, Golang 1")
	go func() {
		fmt.Println("Hello, Golang 2")
		wg.Done()
	}()
	go func() {
		fmt.Println("Hello, Golang 3")
		wg.Done()
	}()

	fmt.Println("mid - Goroutines", runtime.NumGoroutine())
	fmt.Println("mid - CPUs", runtime.NumCPU())

	wg.Wait()

	fmt.Println("end - Goroutines", runtime.NumGoroutine())
	fmt.Println("end - CPUs", runtime.NumCPU())
}
