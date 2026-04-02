package main

// 42. Trapping Rain Water
// https://leetcode.com/problems/trapping-rain-water/

/*
1. Brute Force:
	- For each position, find the maximum height of the left and right sides, and calculate the trapped water.
	- For i , trapped water = min(lMax, rMax) - height[i]
	- TC: O(n^2), SC: O(1)
2.Two Pointers:
	- Use two pointers, left and right, to traverse the height array from both ends.
	- Keep track of the maximum height on both sides (leftMax and rightMax).
	- Move the pointer with the smaller height towards the center, and calculate the trapped water at each step.
	- TC: O(n), SC: O(1)
*/

func trap_brute(height []int) int {
	n := len(height)
	if n == 0 {
		return 0
	}
	res := 0
	for i := 0; i < n; i++ {
		lMax, rMax := height[i], height[i]
		for j := 0; j < i; j++ {
			lMax = max(lMax, height[j])
		}
		for j := i + 1; j < n; j++ {
			rMax = max(rMax, height[j])
		}
		res += min(lMax, rMax) - height[i]
	}
	return res
}

func trap(height []int) int {
	n := len(height)
	if n == 0 {
		return 0
	}
	l, r := 0, n-1
	lMax, rMax := height[l], height[r]
	res := 0
	for l < r {
		if lMax < rMax {
			l++
			lMax = max(lMax, height[l])
			res += lMax - height[l]
		} else {
			r--
			rMax = max(rMax, height[r])
			res += rMax - height[r]
		}
	}
	return res
}
func main() {
	height := []int{0, 1, 0, 2, 1, 0, 1, 3, 2, 1, 2, 1}
	println(trap(height)) // 6
}
