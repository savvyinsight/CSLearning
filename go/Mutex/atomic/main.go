package main

import (
	"fmt"
	"sync"
	"sync/atomic"
	"time"
)

var count int64 = 0

// No mutex and atomic
func original() {
	wg := sync.WaitGroup{}
	start := time.Now()
	for _ = range 10000 {
		wg.Add(1)
		go func() {
			count++
			wg.Done()
		}()
	}
	wg.Wait()
	fmt.Printf("time cost: %v, count: %d\n", time.Since(start), count)
}

func withMutex() {
	wg := sync.WaitGroup{}
	lock := sync.Mutex{}
	start := time.Now()

	for _ = range 10000 {
		wg.Add(1)
		go func() {
			lock.Lock()
			count++
			lock.Unlock()
			wg.Done()
		}()
	}
	wg.Wait()
	fmt.Printf("time cost: %v, count: %d\n", time.Since(start), count)
}

func withAtomic() {
	wg := sync.WaitGroup{}
	start := time.Now()

	for _ = range 10000 {
		wg.Add(1)
		go func() {
			atomic.AddInt64(&count, 1)
			wg.Done()
		}()
	}
	wg.Wait()
	fmt.Printf("time cost: %v, count: %d\n", time.Since(start), count)
}

func main() {
	// original()
	// withMutex()
	withAtomic()

}
