package main

import "fmt"

func main() {
	map1()
}

func f(s []int) {
	for i := range s {
		s = append(s, i)
	}
}

func f1(s *[]int) {
	for i := range *s {
		*s = append(*s, i)
	}
}

func map1() {
	s := map[int]string{1: "good", 2: "h", 3: "sf", 4: "aA", 78: "sdfasdfsf", 12: "sfasdf", 8: "8", 7: "ha7"}
	for i, v := range s {
		fmt.Println(i, v)
	}
}
