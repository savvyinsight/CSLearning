package main

import "sort"

// 1. Two Sum
// https://leetcode.com/problems/two-sum/

/*
Thinking process:
1. Brute force
   - Iterate through all pairs of numbers in the array and check if their sum equals the target.
   - Time complexity: O(n^2)
   - Space complexity: O(1)

2. Hash map
   - Use a hash map to store the numbers and their indices as we iterate through the array.
   - For each number, calculate its complement (target - current number) and check if it exists in the hash map.
   - If it exists, return the indices of the current number and the complement.
   - If it doesn't exist, add the current number and its index to the hash map.
   - Time complexity: O(n)
   - Space complexity: O(n)

3. Sorting and two pointers
   - Sort the array while keeping track of the original indices.
   - Use two pointers, one at the beginning and one at the end of the sorted array.
   - Move the pointers towards each other based on the sum of the numbers they point to.
   - If the sum equals the target, return the original indices.
   - If the sum is less than the target, move the left pointer to the right.
   - If the sum is greater than the target, move the right pointer to the left.
   - Time complexity: O(n log n) due to sorting.
   - Space complexity: O(1) if we sort in place.
*/

// approach 1: brute force
func twoSum_1(nums []int, target int) []int {
	for i := 0; i < len(nums); i++ {
		for j := i + 1; j < len(nums); j++ {
			if nums[i]+nums[j] == target {
				return []int{i, j}
			}
		}
	}
	return []int{}
}

// approach 2: hash map
func twoSum_2(nums []int, target int) []int {
	hashMap := make(map[int]int)
	for i, num := range nums {
		complement := target - num
		if index, found := hashMap[complement]; found {
			return []int{index, i}
		}
		hashMap[num] = i
	}
	return []int{}
}

// approach 3: sorting and two pointers
func twoSum_3(nums []int, target int) []int {
	type pair struct {
		value int
		index int
	}
	pairs := make([]pair, len(nums))
	for i, num := range nums {
		pairs[i] = pair{value: num, index: i}
	}
	// Sort pairs based on value
	sort.Slice(pairs, func(i, j int) bool {
		return pairs[i].value < pairs[j].value
	})
	left, right := 0, len(pairs)-1
	for left < right {
		sum := pairs[left].value + pairs[right].value
		if sum == target {
			// output the original indices, not the sorted indices
			// [2,4] not [4,2]
			// so we need to sort the indices before returning
			if pairs[left].index < pairs[right].index {
				return []int{pairs[left].index, pairs[right].index}
			} else {
				return []int{pairs[right].index, pairs[left].index}
			}
		} else if sum < target {
			left++
		} else {
			right--
		}
	}
	return []int{}
}

/*
nums=[-1,-2,-3,-4,-5]
target=-8
output=[2,4]
*/

func main() {
	// nums := []int{2, 7, 11, 15}
	// target := 9
	nums := []int{-1, -2, -3, -4, -5}
	target := -8
	result := twoSum_3(nums, target)
	println(result[0], result[1]) // Output: 2, 4
}
