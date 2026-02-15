package main

import (
	"errors"
	"fmt"
	"runtime"
	"time"
)

// type File interface {
// 	Read(b Buffer) bool
// 	Write(b Buffer) bool
// 	Close()
// }

type Person struct {
	Name string
}

func a() {
	var dayOfWeek int = 4

	switch dayOfWeek {
	case 1:
		fmt.Println("Monday")
		fallthrough
	case 2:
		fmt.Println("Tuesday")
		fallthrough
	case 3:
		fmt.Println("Wednesday")
		fallthrough
	case 4:
		fmt.Println("Thursday")
		fallthrough
	case 5:
		fmt.Println("Friday")
		fallthrough
	case 6:
		fmt.Println("Saturday")
	case 7:
		fmt.Println("Sunday")
	default:
		fmt.Println("Invalid Day")
	}
}

func t() {
	for i := 0; i < 10; i++ {
		if i == 5 {
			goto end // 跳到 end 位置执行
		}
		fmt.Println(i)
	}

end:
	fmt.Println("end")
}

func fibonacci(c, quit chan int) {
	x, y := 0, 1
	for {
		select {
		case c <- x:
			x, y = y, x+y
		case value := <-quit:
			fmt.Printf("quit信号值: %d\n", value)
			return
		}
	}
}

func slection() {
	c := make(chan int)
	quit := make(chan int)
	go func() {
		for i := 0; i < 10; i++ {
			fmt.Println(<-c)
		}
		quit <- 0
	}()
	fibonacci(c, quit)
}

func aboutError() {
	// 使用 New 创建错误
	err1 := errors.New("first error")

	// 使用 fmt 创建错误
	err2 := fmt.Errorf("second %s", "error")

	fmt.Println(err1, err2)
}

func aboutArray() {
	// var data []int

	// for i := 0; i < 10; i++ {
	// 	data = append(data, i)
	// }

	// fmt.Println(data)
}

func aboutMap() {
	// 初始化 Map
	var dataMap map[string]string
	dataMap = make(map[string]string)

	// 也可以这样初始化
	//dataMap := map[string]string{}

	// 添加键值
	dataMap["first"] = "first value"
	dataMap["second"] = "second value"
	dataMap["third"] = "third value"

	fmt.Println("print key and value: ")
	delete(dataMap, "first") // you can delete
	fmt.Println(dataMap)
	/*
		// 使用 range 遍历 key value
		for key, val := range dataMap {
			fmt.Printf("key: %s  -  value: %s \n", key, val)
		}

		fmt.Println("print key: ")

		// 使用 range 遍历 key
		for key := range dataMap {
			fmt.Printf("key: %s \n", key)
		}

		fmt.Println("print value: ")

		// 使用 range 遍历 value
		for _, val := range dataMap {
			fmt.Printf("Value: %s \n", val)
		}
	*/
}

func aboutSlices() {
	// 第一种方式： 直接声明
	var dataSlice []string

	// 第二种方式： 直接初始化
	dataSlice1 := []string{}

	// 第三种方式： 使用 make
	dataSlice2 := make([]string, 10)

	fmt.Println(dataSlice, dataSlice1, dataSlice2)
}

func aboutTimer() {
	t := time.NewTicker(3 * time.Second)
	fmt.Println("start")
	<-t.C
	fmt.Println("finish")
}

func aboutGoroutineCommunicate() {
	ch := make(chan string)

	go sendData(ch)

	go getData(ch)
	time.Sleep(time.Second)
}

func aboutGoroutine() {
	ch := make(chan int)

	go func() {
		for {
			select {
			case <-ch:
				fmt.Println("exit")
				return
			default:
				fmt.Println("running...")
				time.Sleep(time.Second)
			}
		}
	}()

	time.Sleep(5 * time.Second)
	ch <- 1
}

func getCPUNum() {
	fmt.Println("CPU逻辑核心数:", runtime.NumCPU())
	fmt.Println("当前GOMAXPROCS:", runtime.GOMAXPROCS(0))
}

func aboutPanic() {
	for i := range 4 {
		go testPanic(i)
	}

	time.Sleep(time.Second)
}

func testPanic(i int) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("recover panic: ", i)
			panic("panic in defer")
		}
	}()

	if i == 1 {
		panic(fmt.Sprintf("panic %d", i))
	}

	fmt.Println("test panic: ", i)
}

func main() {
	// ch := make(chan int, 20)
	// fmt.Println(ch)

	// m := make(map[string]string, 20)
	// m["s"] = "good"
	// fmt.Println(m["s"])

	// a()

	// // for range 遍历
	// array := []int{1, 2, 23, 4, 5}
	// for i, v := range array {
	// 	fmt.Println(i, v)
	// }

	// // break
	// for i, v := range array {
	// 	if i >= 2 {
	// 		break
	// 	}
	// 	fmt.Println(i, v)
	// }

	// t()

	// slection()

	// go func() {
	// 	fmt.Println("no.2")
	// }()
	// go func() {
	// 	slection()
	// }()
	// fmt.Println("start goroutine")
	// time.Sleep(2 * time.Second)
}

func sendData(ch chan string) {
	ch <- "good"
	ch <- "boy"
}

func getData(ch chan string) {
	var input string
	for {
		input = <-ch
		fmt.Println(input)
	}
}
