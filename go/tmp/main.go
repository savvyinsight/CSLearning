package main

import (
	"fmt"
	"reflect"
	"time"
)

func lengthOfLongestSubstring(s string) int {
	charIndex := make(map[byte]int)

	maxLen := 0

	for r := 0; r < len(s); r++ {
		// if idx, exists := charIndex[s[r]]; exists && idx >= l {
		// 	l = idx + 1
		// }
		charIndex[s[r]] = r
		// if r-l+1 > maxLen {
		// 	maxLen = r - l + 1
		// }
	}
	fmt.Println(charIndex)
	return maxLen
}

func main() {
	// s := "pwwkew"
	// b := s[0]
	// fmt.Println(b)
	// b1 := []rune(s)
	// fmt.Println(b1)
	// t := lengthOfLongestSubstring(s)
	// fmt.Println(t)

	// s := 'S'
	// fmt.Println((rune)(s))
	// fmt.Println((rune)('a'))
	// fmt.Println((rune)('A'))
	// fmt.Println((rune)('a' - 'A'))
	// fmt.Println(toLower(s))

	// s := "good"
	// runes := []rune(s)
	// bytes := []byte(s)
	// fmt.Println(runes, bytes)

	// num := 100
	// testReflect(num)

	// timer()
	str := fmt.Sprintf("%d",time.Now().Unix())
	fmt.Printf(str)

}

func testReflect(i interface{}) {
	fmt.Println(i)
	ty := reflect.TypeOf(i)
	fmt.Printf("type is %T\n", ty)
	tv := reflect.ValueOf(i)
	fmt.Println(tv)
	k1 := ty.Kind()
	fmt.Println(k1)
	fmt.Println("--------------")
	t1 := tv.Interface()
	fmt.Println("interface: ", t1)
	n, flag := t1.(float32)
	if flag == false {
		fmt.Println("dead", n)
	}
}

func toLower(ch rune) rune {
	if ch >= 'A' && ch <= 'Z' {
		return ch + ('a' - 'A')
	}
	return ch
}

func timer() {
	t := time.NewTicker(3 * time.Second)
	fmt.Println("Start")
	<-t.C
	fmt.Println("End")
}
