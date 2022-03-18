package main

import (
	"sync"
	"time"
)

func main() {
	done := make(chan bool, 1)
	var mu sync.Mutex

	// goroutine 1
	go func() {
		for {
			select {
			case <-done:
				return
			default:
				mu.Lock()
				time.Sleep(100 * time.Microsecond)
				mu.Unlock()
			}
		}
	}()

	// goroutine 2
	for i := 0; i < 10; i++ {
		t0 := time.Now()
		time.Sleep(100 * time.Microsecond)
		mu.Lock()
		dt := time.Since(t0)
		mu.Unlock()
		_ = dt
	}
	done <- true
}
