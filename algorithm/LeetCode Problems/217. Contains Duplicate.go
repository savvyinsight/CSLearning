package main

import "sort"

// 217. Contains Duplicate
// Given an integer array nums, return true if any value appears at least twice in the array, and return false if every element is distinct.

/*
Thinking process:
1. brute force
   - We can use two nested loops to compare each element with every other element in the array.
   - Time complexity: O(n^2)
   - Space complexity: O(1)

2. sorting
   - We can sort the array and then check for adjacent elements that are the same.
   - Time complexity: O(n log n) due to sorting
   - Space complexity: O(1) if we sort in place, otherwise O(n)

3. hash set
   - We can use a hash set to keep track of the elements we have seen so far.
   - For each element, we check if it is already in the set. If it is, we return true. If not, we add it to the set.
   - Time complexity: O(n) on average
   - Space complexity: O(n) in the worst case if all elements are distinct
*/

// approach 1: brute force
func hasDuplicate_1(nums []int) bool {
	for i := 0; i < len(nums); i++ {
		for j := i + 1; j < len(nums); j++ {
			if nums[i] == nums[j] {
				return true
			}
		}
	}
	return false
}

// approach 2: sorting

func hasDuplicate_2(nums []int) bool {
	sort.Ints(nums)
	for i := 1; i < len(nums); i++ {
		if nums[i] == nums[i-1] {
			return true
		}
	}
	return false
}

// approach 3: hash set
func hasDuplicate_3(nums []int) bool {
	seen := make(map[int]bool)
	for _, num := range nums {
		if seen[num] {
			return true
		}
		seen[num] = true
	}
	return false
}

func hasDuplicate_4(nums []int) bool {
	seen := make(map[int]struct{})
	for _, num := range nums {
		if _, exists := seen[num]; exists {
			return true
		}
		seen[num] = struct{}{}
	}
	return false
}

/*
Example 1:
Input: nums = [1,2,3,1]
Output: true

Example 2:
Input: nums = [1,2,3,4]
Output: false

Example 3:
Input: nums = [1,1,1,3,3,4,3,2,4,2]
Output: true
*/

func main() {
	nums1 := []int{1, 2, 3, 1}
	nums2 := []int{1, 2, 3, 4}
	nums3 := []int{1, 1, 1, 3, 3, 4, 3, 2, 4, 2}

	println(hasDuplicate_1(nums1)) // true
	println(hasDuplicate_1(nums2)) // false
	println(hasDuplicate_1(nums3)) // true

	println(hasDuplicate_2(nums1)) // true
	println(hasDuplicate_2(nums2)) // false
	println(hasDuplicate_2(nums3)) // true

	println(hasDuplicate_3(nums1)) // true
	println(hasDuplicate_3(nums2)) // false
	println(hasDuplicate_3(nums3)) // true
}
