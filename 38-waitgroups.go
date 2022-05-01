package main

import (
	"fmt"
	"sync"
	"time"
)

func worker(id int) {
	fmt.Printf("#%d starting \n", id)

	time.Sleep(time.Second)
	fmt.Printf("#%d done\n", id)
}

func main() {
	// NOTE: if a WaitGroup is explicitly passed into functions,
	// it should be done by pointer.
	var wg sync.WaitGroup

	for i := 1; i <= 5; i++ {
		wg.Add(1)

		// Avoid re-use of the same i value in each goroutine closure.
		// https://go.dev/doc/faq#closures_and_goroutines
		i := i

		go func() {
			defer wg.Done()
			worker(i)
		}()
	}

	wg.Wait()
}
