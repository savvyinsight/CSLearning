package main

import (
	"context"
	"fmt"
	"time"
)

func test1() {
	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()
	done := make(chan bool)

	go func(ctx1 context.Context) {
		defer func() {
			done <- true // 通知主goroutine
		}()

		for {
			select {
			case <-ctx1.Done():
				fmt.Println("time out")
				return
			default:
				fmt.Println("running...")
				time.Sleep(time.Second)
			}
		}
	}(ctx)

	<-done // 等待goroutine结束
	fmt.Println("程序结束")
}

func test() {
	ctx, _ := context.WithTimeout(context.Background(), 3*time.Second)

	go func(ctx context.Context) {
		for {
			select {
			case <-ctx.Done():
				fmt.Println("time out")
				return
			default:
				fmt.Println("running....")
				time.Sleep(time.Second)
			}
		}
	}(ctx)
}

func testWithCancel() {
	ctx, cancel := context.WithCancel(context.Background())

	go func(ctx context.Context) {
		for {
			select {
			case <-ctx.Done():
				fmt.Println("canceled...")
				return
			default:
				fmt.Println("running...")
				time.Sleep(time.Second)
			}
		}
	}(ctx)

	time.Sleep(time.Second)
	cancel()
	time.Sleep(time.Second)
}

func main() {
	testWithCancel()
	time.Sleep(5 * time.Second)
}
