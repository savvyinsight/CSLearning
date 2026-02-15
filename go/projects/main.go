package main

import (
	"fmt"
	"sync"
)

func main() {
	ch := make(chan struct{})
	var wg sync.WaitGroup
	wg.Add(2)

	// Start numbers first to guarantee order
	go printNumbers(ch, &wg)
	go printLetters(ch, &wg)

	// Signal start to numbers
	ch <- struct{}{}

	wg.Wait()
	fmt.Println()
}

func printNumbers(ch chan struct{}, wg *sync.WaitGroup) {
	defer wg.Done()
	for i := 1; i <= 10; i++ {
		<-ch // Wait for turn
		fmt.Printf("%d ", i)
		ch <- struct{}{} // Always pass turn
	}
}

func printLetters(ch chan struct{}, wg *sync.WaitGroup) {
	defer wg.Done()
	for i := 0; i < 10; i++ {
		<-ch // Wait for turn
		fmt.Printf("%c ", 'a'+i)
		// Pass turn back (including on last iteration)
		ch <- struct{}{}
	}
}
