package main

import "fmt"

// 739. Daily Temperatures
// https://leetcode.com/problems/daily-temperatures/

/*
Thinking process:
1. stack:
	- because we need to look at future temperature, so when we get to future element, we need to
	compare with the past element. so use stack to store it.
	- when stack is empty, we can add element. then when future element greater than stack element,
	we can compute , so it needs to store value and index. and pop the past element, afterall we got
	future warmer day.

	TC:O(n),SC:O(n)
*/

func dailyTemperatures(temperatures []int) []int {
	type pair struct {
		index int
		val   int
	}
	stk := []pair{}
	res := make([]int, len(temperatures), len(temperatures))

	for i, v := range temperatures {
		for len(stk) > 0 && stk[len(stk)-1].val < v {
			idx := stk[len(stk)-1].index
			stk = stk[:len(stk)-1]
			res[idx] = i - idx
		}
		stk = append(stk, pair{i, v})
	}
	return res
}

func main() {
	tempe := []int{73, 74, 75, 71, 69, 72, 76, 73}
	res := dailyTemperatures(tempe)
	fmt.Print(res)
}
