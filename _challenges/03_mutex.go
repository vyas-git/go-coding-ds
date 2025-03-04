package main

import (
	"fmt"
	"sync"
)

// Assignment 3: Mutex for Synchronization Problem:
// Write a Go program that uses mutexes to synchronize access to a shared variable.
// Multiple goroutines should increment the variable concurrently, and the final value should be printed.

func mains() {
	type MU struct {
		v  int
		mu sync.RWMutex
	}
	no_of_goroutines := 500
	var data MU
	var wg sync.WaitGroup
	for range no_of_goroutines {
		wg.Add(1)
		go func() {
			data.mu.RLocker().Lock()
			data.mu.RLock()

			data.v++
			data.mu.Unlock()

			defer wg.Done()
		}()
	}

	wg.Wait()
	fmt.Println(data.v)
}
