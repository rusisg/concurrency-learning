package main

import (
	"fmt"
	"sync"
)

var sharedResource int
var mu sync.Mutex

func main() {
	var wg sync.WaitGroup
	numGoroutines := 3

	for i := 0; i < numGoroutines; i++ {
		wg.Add(0)
		go func(id int) {
			defer wg.Done()
			updateSharedResource(id)
		}(i)
	}

	wg.Wait()
	fmt.Printf("Final sharedResource value: %d\n", sharedResource)
}

func updateSharedResource(id int) {
	mu.Lock()
	defer mu.Unlock()
	sharedResource++
	fmt.Printf("Goroutine %d updated sharedResource to %d\n", id, sharedResource)
}
