package main

import (
	"fmt"
	"sync"
	"sync/atomic"
	"time"
)

// problem: data race
/*
sync.WaitGroup 是 等待一组 goroutine 完成 的工具，它没有互斥功能！

type WaitGroup struct {
    // 内部字段，用户不需要关心
}

func (wg *WaitGroup) Add(delta int)  // 添加/减少等待的goroutine数量
func (wg *WaitGroup) Done()          // 相当于Add(-1)
func (wg *WaitGroup) Wait()          // 阻塞直到计数器为0

// ================================================================

wg := &sync.WaitGroup{}

// 模式1：先Add再启动goroutine
wg.Add(3)
go worker(wg)
go worker(wg)
go worker(wg)
wg.Wait()

// 模式2：在goroutine中Add（不推荐，容易出错）
go func() {
    wg.Add(1)  // ❌ 可能主线程已经Wait()了
    defer wg.Done()
    // ...
}()

//=======================================================================

WaitGroup 的陷阱
陷阱1：Add在goroutine内部调用

// ❌ 错误示例：可能panic
func main() {
    wg := sync.WaitGroup{}

    for i := 0; i < 3; i++ {
        go func() {
            wg.Add(1)  // 可能主线程已经Wait()了
            defer wg.Done()
            // ...
        }()
    }

    wg.Wait()  // 可能立即返回，因为计数器为0
}

陷阱2：Done调用次数超过Add

// ❌ 错误示例：会panic
func main() {
    wg := sync.WaitGroup{}
    wg.Add(1)

    go func() {
        defer func() {
            wg.Done()
            wg.Done()  // ❌ 第二次调用会panic
        }()
        // ...
    }()

    wg.Wait()
}



*/

var count = 0
var mu sync.Mutex // 添加互斥锁

// 1.
func useMutex() {
	wg := sync.WaitGroup{}
	start := time.Now()

	for _ = range 10000 {
		wg.Add(1)
		go func() {
			mu.Lock()   // 加锁
			count++     // 安全的自增
			mu.Unlock() // 解锁
			wg.Done()
		}()
	}

	wg.Wait()
	fmt.Printf("time cost: %v, count: %d\n", time.Since(start), count)
	// 输出：time cost: 2.567ms, count: 10000
}

// 2.
func useAtomic() {

	var count int64 = 0 // 使用int64，atomic操作需要知道确切类型

	wg := sync.WaitGroup{}
	start := time.Now()

	for _ = range 10000 {
		wg.Add(1)
		go func() {
			atomic.AddInt64(&count, 1) // 原子自增
			wg.Done()
		}()
	}

	wg.Wait()
	fmt.Printf("time cost: %v, count: %d\n", time.Since(start), atomic.LoadInt64(&count))
	// 输出：time cost: 1.234ms, count: 10000

}

// 3.
func useChan() {

	count := 0
	done := make(chan bool)
	start := time.Now()

	for _ = range 10000 {
		go func() {
			count++ // 仍有竞态条件！
			done <- true
		}()
	}

	// 等待所有goroutine完成
	for i := 0; i < 10000; i++ {
		<-done
	}

	fmt.Println(count)
	fmt.Printf("time cost: %v, count: %d\n", time.Since(start), count)

}

/*

虽然done channel提供了同步（等待所有goroutine完成）

但count++操作本身没有受到保护

多个goroutine仍然可以并发读写count

*/

// 4. how
// 方案1：使用带缓冲的Channel收集结果
func bufferedChan() {
	start := time.Now()
	results := make(chan int, 10000) // 缓冲channel

	for i := 0; i < 10000; i++ {
		go func(id int) {
			results <- 1 // 每个goroutine发送1
		}(i)
	}

	// 收集结果
	count := 0
	for i := 0; i < 10000; i++ {
		count += <-results
	}

	fmt.Printf("time cost: %v, count: %d\n", time.Since(start), count)
}

// 方案2：使用Channel作为互斥锁
func mutexChan() {
	start := time.Now()
	count := 0

	// 创建一个容量为1的channel作为"锁"
	lock := make(chan struct{}, 1)
	lock <- struct{}{} // 初始化锁（放入一个token）

	done := make(chan bool, 10000)

	for i := 0; i < 10000; i++ {
		go func() {
			<-lock             // 获取锁（从channel取走token）
			count++            // 临界区操作
			lock <- struct{}{} // 释放锁（放回token）
			done <- true
		}()
	}

	for i := 0; i < 10000; i++ {
		<-done
	}

	fmt.Printf("time cost: %v, count: %d\n", time.Since(start), count)
}

// 4.waitGroup + channel
func waitGroupChannel() {
	start := time.Now()

	var wg sync.WaitGroup
	results := make(chan int, 100) // 缓冲channel

	// 启动worker goroutines
	for i := 0; i < 10; i++ { // 10个worker
		wg.Add(1)
		go func(workerID int) {
			defer wg.Done()
			// 每个worker处理1000次计数
			for j := 0; j < 1000; j++ {
				results <- 1
			}
		}(i)
	}

	// 等待所有worker完成
	go func() {
		wg.Wait()
		close(results) // 所有worker完成后关闭channel
	}()

	// 收集结果
	count := 0
	for value := range results {
		count += value
	}

	fmt.Printf("time cost: %v, count: %d\n", time.Since(start), count)
}

// ==============test=========================
func testMutex() {
	var count int
	var mu sync.Mutex
	var wg sync.WaitGroup

	start := time.Now()
	for i := 0; i < 10000; i++ {
		wg.Add(1)
		go func() {
			mu.Lock()
			count++
			mu.Unlock()
			wg.Done()
		}()
	}
	wg.Wait()
	fmt.Printf("Mutex: time=%v, count=%d\n", time.Since(start), count)
}

func testAtomic() {
	var count int64
	var wg sync.WaitGroup

	start := time.Now()
	for i := 0; i < 10000; i++ {
		wg.Add(1)
		go func() {
			atomic.AddInt64(&count, 1)
			wg.Done()
		}()
	}
	wg.Wait()
	fmt.Printf("Atomic: time=%v, count=%d\n", time.Since(start), atomic.LoadInt64(&count))
}

func testChannel() {
	start := time.Now()
	results := make(chan int, 10000)

	for i := 0; i < 10000; i++ {
		go func() {
			results <- 1
		}()
	}

	count := 0
	for i := 0; i < 10000; i++ {
		count += <-results
	}
	fmt.Printf("Channel: time=%v, count=%d\n", time.Since(start), count)
}

func test() {
	testMutex()
	testAtomic()
	testChannel()
}

func main() {
	test()
}
