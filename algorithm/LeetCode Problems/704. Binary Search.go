package main

// 704. Binary Search
// https://leetcode.com/problems/binary-search/

/*
Thinking process:
1. Binary Search
   - Use two pointers to keep track of the left and right boundaries of the search space. Calculate the middle index and compare the target with the middle element. Adjust the search space accordingly.
   - Time Complexity: O(log n) as we are halving the search space in each step.
   - Space Complexity: O(1) for iterative approach, O(log n) for recursive approach due to call stack.
*/

func binarySearch(nums []int, target int) int {
	left, right := 0, len(nums)-1

	for left <= right {
		mid := left + (right-left)/2

		if nums[mid] == target {
			return mid
		} else if nums[mid] < target {
			left = mid + 1
		} else {
			right = mid - 1
		}
	}

	return -1 // target not found
}

func main() {
	println(binarySearch([]int{-1, 0, 3, 5, 9, 12}, 9))  // Output: 4
	println(binarySearch([]int{-1, 0, 3, 5, 9, 12}, 2))  // Output: -1
	println(binarySearch([]int{-1, 0, 3, 5, 9, 12}, -1)) // Output: 0
}
