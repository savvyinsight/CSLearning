package main

import (
	"context"
	"fmt"
	"sync"
	"time"
)

// what will happen if i try to send different types to channel
// asnwer: compile error
// so use Empty interface to send different or any types (interface{})
func chanDiffTypes() {
	c := make(chan interface{}) // Can hold any type

	go func() {
		c <- 42                  // int
		c <- "hello"             // string
		c <- 3.14                // float64
		c <- struct{ X int }{99} // struct
		c <- []int{1, 2, 3}      // slice
		close(c)
	}()

	for v := range c {
		fmt.Printf("Value: %v, Type: %T\n", v, v)
	}
}

// Type Assertion Required
// When receiving, you need type assertions to use the values:
func typeAssertion() {
	ch := make(chan interface{})

	go func() {
		ch <- "hello"
		ch <- 24
		ch <- '!'
		close(ch)
	}()

	for v := range ch {
		switch val := v.(type) {
		case int:
			fmt.Printf("Got int: %d (double: %d)\n", val, val*2)
		case string:
			fmt.Printf("Got string: %s (length: %d)\n", val, len(val))
		default:
			fmt.Printf("Got unknown type: %T\n", val)
		}
	}
}

// Safer: Custom Types/Structs
type Message struct {
	Type   string
	IntVal int
	StrVal string
	Data   interface{}
}

func typeAssertionCustomType() {
	ch := make(chan Message)

	go func() {
		ch <- Message{Type: "int", IntVal: 42}
		ch <- Message{Type: "string", StrVal: "hello"}
		ch <- Message{Type: "slice", Data: []int{1, 2, 3}}
		close(ch)
	}()

	for msg := range ch {
		switch msg.Type {
		case "int":
			fmt.Printf("Integer: %d\n", msg.IntVal)
		case "string":
			fmt.Printf("String: %s\n", msg.StrVal)
		case "slice":
			if slice, ok := msg.Data.([]int); ok {
				fmt.Printf("Slice: %v\n", slice)
			}
		}
	}

}

// Using Generics
// Go 1.18 introduced generics for type-safe channels:
// Get any type channels then print types and values
func Process[T any](ch chan T) {
	for v := range ch {
		fmt.Printf("Processing %T:%v\n", v, v)
	}
}

func typedChannels() {
	intChan := make(chan int)
	StrChan := make(chan string)

	go func() {
		intChan <- 24
		intChan <- 42
		close(intChan)
	}()

	go func() {
		StrChan <- "hello"
		StrChan <- "world"
		close(StrChan)
	}()

	go Process(intChan)
	go Process(StrChan)

	time.Sleep(time.Second)
}

// channel for communication between goroutine
func fibonacci(c, quit chan int) {
	x, y := 0, 1
	for {
		select {
		case c <- x:
			x, y = y, x+y
		case <-quit:
			fmt.Println("quit")
			return
		}
	}
}

func selection() {
	c := make(chan int)
	quit := make(chan int)
	var wg sync.WaitGroup
	wg.Add(1)
	go func() {
		defer wg.Done()
		for i := 0; i < 10; i++ {
			fmt.Println(<-c)
		}
		quit <- 0
	}()
	go fibonacci(c, quit)
	wg.Wait()
}

// Channel closure: A cleaner approach would be
// to close the channel instead of using a separate quit channel:
func fibonacci_1(ch chan int) {
	x, y := 0, 1
	for {
		ch <- x
		x, y = y, x+y
	}
}

func selection_1() {
	ch := make(chan int)
	var wg sync.WaitGroup
	wg.Add(1)
	go func() {
		defer wg.Done()
		for i := 0; i < 10; i++ {
			fmt.Println(<-ch)
		}
	}()
	go fibonacci_1(ch)
	wg.Wait()
}

// Best practice solution (using context for cancellation)
func fibonacci_2(ctx context.Context, ch chan int) {
	x, y := 0, 1
	for {
		select {
		case ch <- x:
			x, y = y, x+y
		case <-ctx.Done():
			return
		}
	}
}

func selection_2() {
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel() //Cleanup/defensive programming

	ch := make(chan int)
	var wg sync.WaitGroup

	wg.Add(1)
	go func() {
		defer wg.Done()
		for i := 0; i < 10; i++ {
			fmt.Println(<-ch)
		}
		cancel()
	}()
	go fibonacci_2(ctx, ch)
	wg.Wait()
}

// Pattern: Result/Error Channels
type Result struct {
	Value interface{}
	Err   error
}

func worker(id int, results chan<- Result) {
	// Simulate work
	if id%2 == 0 {
		results <- Result{Value: 2 * id, Err: nil}
	} else {
		results <- Result{Value: nil, Err: fmt.Errorf("error on worker %d", id)}
	}
}

func task2Worker() {
	results := make(chan Result, 5)

	for i := 0; i < 5; i++ {
		go worker(i, results)
	}

	for i := 0; i < 5; i++ {
		result := <-results
		if result.Err != nil {
			fmt.Printf("Error: %v\n", result.Err)
		} else {
			fmt.Printf("Result: %v\n", result.Value)
		}
	}

}

// Consider your use case - if you need multiple types,
// ask if you really need one channel or multiple specialized channels
func example() {
	// Usually better to use separate channels
	// intCh := make(chan int)
	// stringCh := make(chan string)
	// errorCh := make(chan error)

	// Or use a struct
	type Event struct {
		Type string
		Data interface{}
	}
}

// communicate or data sharing through channel communication
func testCommuWithChan() {
	ch := make(chan string)

	go sendData(ch)
	go recvDate(ch)
	time.Sleep(time.Second)
}

func sendData(ch chan string) {
	ch <- "good"
	ch <- "boy"
}

func recvDate(ch chan string) {
	var input string
	for {
		input = <-ch
		fmt.Printf("%s ", input)
	}
}

func main() {
	task2Worker()
}
