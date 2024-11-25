package main

import (
	"fmt"
	"sync"
	"time"
)

var (
	wg sync.WaitGroup
	mu sync.Mutex
	t  = time.Now()
)

func Counter(arg *int, i int) {
	defer wg.Done()
	mu.Lock()
	*arg++
	mu.Unlock()
	fmt.Printf("[Goroutine %d]: executed at %v\n", i, time.Since(t))
}

func main() {
	var val int

	wg.Add(10_000)
	for i := 0; i < 10_000; i++ {
		go Counter(&val, i)
	}

	wg.Wait()
	fmt.Println(val)
}
