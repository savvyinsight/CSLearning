package main

import "fmt"

// 167. Two Sum II - Input Array Is Sorted
// https://leetcode.com/problems/two-sum-ii-input-array-is-sorted/

/*
Thinking process:
1.Two Pointer
	- use two pointers to track from start and end, bacause the array is sorted
	- TC:O(n),SC(1)

2.Brute force
	- checks every possible pair
	- TC:O(n**2), SC:O(1)

3.Binary Search
	- for every element we can find complement number, than birary search the
	complemety element O(log n) as array sorted
	- TC:O(n log n), SC:O(1)

4.Hash Map
	- For each number compute needed complement. and check it on Hash Map.
	- after compute complete number for each number we check hash map to has already been
	stored in the hash. if not , we add current number to hash map. iterate until find.
	- TC:O(n),SC:O(n)
*/

// Two pointer
func twoSum_1(numbers []int, target int) []int {
	l, r := 0, len(numbers)-1

	for l < r {
		sum := numbers[l] + numbers[r]
		if sum == target {
			return []int{l + 1, r + 1}
		} else if sum > target {
			r--
		} else {
			l++
		}
	}
	return []int{}
}

// Brute force
func twoSum_2(numbers []int, target int) []int {
	for i := 0; i < len(numbers); i++ {
		for j := i + 1; j < len(numbers); j++ {
			if target == numbers[i]+numbers[j] {
				return []int{i + 1, j + 1}
			}
		}
	}
	return []int{}
}

// Binary Search
func twoSum_3(numbers []int, target int) []int {
	for i := 0; i < len(numbers); i++ {
		l, r := i+1, len(numbers)-1
		complement := target - numbers[i]
		for l <= r {
			mid := l + (r-l)/2
			if numbers[mid] == complement {
				return []int{i + 1, mid + 1}
			} else if numbers[mid] < complement {
				l = mid + 1
			} else {
				r = mid - 1
			}
		}
	}
	return []int{}
}

// Hash Map
func twoSum(numbers []int, target int) []int {
	hashMap := make(map[int]int)
	for i := 0; i < len(numbers); i++ {
		complement := target - numbers[i]
		if v, ok := hashMap[complement]; ok {
			return []int{v + 1, i + 1}
		}
		hashMap[numbers[i]] = i
	}
	return []int{}
}
func main() {
	nums := []int{2, 7, 11, 15}
	res := twoSum(nums, 9)
	for _, v := range res {
		fmt.Print(v, " ")
	}
}
