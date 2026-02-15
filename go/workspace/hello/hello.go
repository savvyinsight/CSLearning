package main

import (
	"fmt"

	"strconv"

	"golang.org/x/example/hello/reverse"
)

// Int returns the decimal reversal of the integer i.
func Int(i int) int {
	i, _ = strconv.Atoi(reverse.String(strconv.Itoa(i)))
	return i
}

func main() {
	fmt.Println(reverse.String("Hello"))
	fmt.Println(Int(12345))
}
