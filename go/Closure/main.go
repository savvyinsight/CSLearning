package main

import "fmt"

func main() {
	count := func() func() int {
		i := 0
		return func() int {
			i++
			return i
		}
	}()
	fmt.Println(count())
	fmt.Println(count())
}
