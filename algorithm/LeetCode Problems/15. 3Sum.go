package main

import (
	"fmt"
	"sort"
)

// 15. 3Sum
// https://leetcode.com/problems/3sum/
/*
Thinking process:
1.Brute force
	- try all combinations, use sort for avoid duplicates, use set to get unique triplets
	- TC:O(n^3),SC:O(1)
2.Hash map
	- sort first, then use hashmap to store frequency of each number.
	- sort for easy to avoid duplicates
	- hashmap for for quick look up and also avoid duplicates
	- TC:O(n^2),SC:O(n)
3.Two pointers

*/

func threeSum_brute(nums []int) [][]int {
	sort.Ints(nums)
	set := map[[3]int]struct{}{}

	for i := 0; i < len(nums); i++ {
		for j := i + 1; j < len(nums); j++ {
			for k := j + 1; k < len(nums); k++ {
				if nums[i]+nums[j]+nums[k] == 0 {
					set[[3]int{nums[i], nums[j], nums[k]}] = struct{}{}
				}
			}
		}
	}
	var res [][]int
	for trip := range set {
		res = append(res, []int{trip[0], trip[1], trip[2]})
	}
	return res
}

func threeSum_hashmap(nums []int) [][]int {
	sort.Ints(nums)
	count := map[int]int{}

	// for num := range nums {
	// 	count[num]++
	// }

	for i := 0; i < len(nums); i++ {
		count[nums[i]]++
	}

	res := [][]int{}
	for i := 0; i < len(nums); i++ {
		count[nums[i]]--
		if i > 0 && nums[i] == nums[i-1] {
			continue
		}
		for j := i + 1; j < len(nums); j++ {
			count[nums[j]]--
			if j > 0 && nums[j] == nums[j-1] {
				continue
			}
			target := -(nums[i] + nums[j])
			if count[target] > 0 {
				res = append(res, []int{nums[i], nums[j], target})
			}
		}
		for j := i + 1; j < len(nums); j++ {
			count[nums[j]]++
		}
	}
	return res
}

func threeSum_twopointer(nums []int) [][]int {
	sort.Ints(nums)
	var res [][]int

	for i := 0; i < len(nums); i++ {
		if nums[i] > 0 {
			break
		}
		if i > 0 && nums[i] == nums[i-1] {
			continue
		}
		l, r := i+1, len(nums)-1
		for l < r {
			sum := nums[i] + nums[l] + nums[r]
			if sum > 0 {
				r--
			} else if sum < 0 {
				l++
			} else {
				res = append(res, []int{nums[i], nums[l], nums[r]})
			}
			for l < r && nums[l] == nums[l+1] {
				l++
			}
			for l < r && nums[r] == nums[r-1] {
				r--
			}
			l++
			r--
		}
	}
	return res
}

func main() {
	// test cases
	fmt.Println(threeSum_twopointer([]int{-1, 0, 1, 2, -1, -4})) // [[-1,-1,2],[-1,0,1]]
}
