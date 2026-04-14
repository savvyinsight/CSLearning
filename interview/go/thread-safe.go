package main

import "sync"

// Questions
// 1.Counter is a simple counter that is not thread-safe.
type Counter struct {
	value int
}

func (c *Counter) Increment() {
	c.value++
}

func (c *Counter) Value() int {
	return c.value
}

// Your task is to make the Counter thread-safe by using a mutex.
type SafeCounter struct {
	value int
	mu    sync.Mutex
}

func (c *SafeCounter) Increment() {
	c.mu.Lock()
	defer c.mu.Unlock()
	c.value++
}

func (c *SafeCounter) Value() int {
	c.mu.Lock()
	defer c.mu.Unlock()
	return c.value
}

// 2. multi-threadAccountTransferer is a simple bank account that allows for concurrent transfers.
// See multi-threadAccountTransferer.go for the implementation.
