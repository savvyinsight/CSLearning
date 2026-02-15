package main

import (
	"fmt"
	"sync"
)

// func main() {
// 	ch := make(chan struct{})

// 	var wg sync.WaitGroup

// 	wg.Add(2)

// 	go printNumbers(ch, &wg)
// 	go printLetters(ch, &wg)

// 	ch <- struct{}{} //Signal start
// 	wg.Wait()
// }

func main() {
	// str := []int{1, 2, 3}
	// t(str)
	// fmt.Println(str)

	m := map[int]string{1: "good", 2: "boy", 3: "hello", 4: "fine"}
	for _, v := range m {
		fmt.Println(v)
	}
}

func t(str []int) {
	str[1] = 23
}

func printNumbers(ch chan struct{}, wg *sync.WaitGroup) {
	defer wg.Done()
	for i := 1; i <= 10; i++ {
		<-ch // Wait for turn
		fmt.Printf("%d ", i)
		ch <- struct{}{} // Pass turn
	}
}

func printLetters(ch chan struct{}, wg *sync.WaitGroup) {
	defer wg.Done()
	for i := 0; i < 10; i++ {
		<-ch // Wait for turn
		fmt.Printf("%c ", 'a'+i)
		if i < 9 {
			ch <- struct{}{} // Pass turn back(except last)
		}
	}

}
