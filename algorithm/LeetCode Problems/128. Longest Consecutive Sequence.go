package main

import "fmt"

// 128. Longest Consecutive Sequence
// https://leetcode.com/problems/longest-consecutive-sequence/description/

/*
Thinking process:
1.Brute force
	- For each element x, check if x+1, exists, x+2 exists, etc. but that would be O(n^2) because
	for each x we may loop through many numbers.

	Optimizing with set:
		1.Put all numbers in a hash set to O(1) lookups.
		2.For each number in the set, check only if it's the start of a squence,->meaning num-1 is
		not in the set.
		3.Then start counting: current = num, streak = 1
		4.While current + 1 in the set, increment current and streak.
		5.Update maximum streak.

	This ensures each number is processed only once(as a start of a squence), and inside the while
	loop we visit each number in that squence only once overall, so total O(n).
*/

func longestConsecutive(nums []int) int {
	set := make(map[int]bool)
	for _, num := range nums {
		set[num] = true
	}

	MaxStreak := 0
	for num := range set {
		// Check if it's the start of a squence
		if !set[num-1] {
			currNum := num
			currStreak := 1

			for set[currNum+1] {
				currNum++
				currStreak++
			}
			if currStreak > MaxStreak {
				MaxStreak = currStreak
			}
		}
	}
	return MaxStreak
}

func main() {
	nums := []int{100, 4, 200, 1, 3, 2}
	res := longestConsecutive(nums)
	fmt.Print(res)
}
