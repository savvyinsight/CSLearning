package main

import (
	"fmt"
	"runtime"
	"sync"
	"time"
)

func main() {
	// 1. 查看CPU信息
	fmt.Printf("逻辑CPU核心数: %d\n", runtime.NumCPU())
	fmt.Printf("GOMAXPROCS当前值: %d\n", runtime.GOMAXPROCS(0))

	// 2. 创建大量goroutine观察调度
	var wg sync.WaitGroup
	start := time.Now()

	for i := 0; i < 10000; i++ {
		wg.Add(1)
		go func(id int) {
			defer wg.Done()
			// 模拟一些工作
			time.Sleep(10 * time.Millisecond)
		}(i)
	}
	wg.Wait()

	fmt.Printf("执行10000个goroutine耗时: %v\n", time.Since(start))

	// 3. 查看当前goroutine数量
	fmt.Printf("当前goroutine数量: %d\n", runtime.NumGoroutine())
}
