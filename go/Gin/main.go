package main

import (
	"fmt"
	"sync"
	"time"
)

var totalNum int
var lock sync.RWMutex

func read() {
	defer wg.Done()
	lock.RLock()
	fmt.Println("SR")
	time.Sleep(time.Second)
	fmt.Println("ER")
	lock.RUnlock()
}

func write() {
	defer wg.Done()
	lock.Lock()
	fmt.Println("SW")
	time.Sleep(time.Second)
	fmt.Println("EW")
	lock.Unlock()
}

var wg sync.WaitGroup

func main() {
	wg.Add(6)

	for i := 0; i < 5; i++ {
		go read()
	}
	go write()
	wg.Wait()
	fmt.Println(totalNum)
}
