package main

import (
	"fmt"
	"sync"
)

type Account struct {
	id      int
	balance int
	mu      sync.Mutex
}

func transfer(from, to *Account, amount int) {
	// Lock the accounts in a consistent order to prevent deadlocks
	if from == to {
		return
	}

	if from.id < to.id {
		from.mu.Lock()
		defer from.mu.Unlock()

		to.mu.Lock()
		defer to.mu.Unlock()
	} else {
		to.mu.Lock()
		defer to.mu.Unlock()

		from.mu.Lock()
		defer from.mu.Unlock()
	}

	if from.balance >= amount {
		from.balance -= amount
		to.balance += amount
	}
}

func main() {
	a1 := Account{id: 1, balance: 1000}
	a2 := Account{id: 2, balance: 1000}
	wg := sync.WaitGroup{}
	wg.Add(2000)
	for i := 0; i < 1000; i++ {
		go func() {
			defer wg.Done()
			transfer(&a1, &a2, 1)
		}()
		go func() {
			defer wg.Done()
			transfer(&a2, &a1, 1)
		}()
	}
	wg.Wait()
	fmt.Printf("Final Balance\n")
	fmt.Printf("A1 balance: %d\n", a1.balance)
	fmt.Printf("A2 alance: %d\n", a2.balance)
}
