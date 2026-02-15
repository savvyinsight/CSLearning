package main

import (
	"fmt"
	"sync"
	"time"
)

func main() {
	lock := sync.RWMutex{}
	go Read(&lock)
	go Read(&lock)
	go Read(&lock)
	go Read(&lock)
	go Write(&lock)
	time.Sleep(7 * time.Second)
}

func Read(lock *sync.RWMutex) {
	lock.RLock()
	fmt.Println("Read lock at:", time.Now())
	time.Sleep(time.Second)
	fmt.Println("Read at :", time.Now())
	lock.RUnlock()
	fmt.Println("Read unlock at :", time.Now())
	return
}

func Write(lock *sync.RWMutex) {
	lock.Lock()
	fmt.Println("Write lock at: ", time.Now())
	time.Sleep(time.Second)
	fmt.Println("Write at: ", time.Now())
	lock.Unlock()
	fmt.Println("Write unlock at: ", time.Now())
	return
}
