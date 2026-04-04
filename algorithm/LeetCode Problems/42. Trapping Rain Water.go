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

3. prefix and suffix arrays:
	- Create two arrays, lMax and rMax, to store the maximum height to the left and right of each position.
	- First traverse the height array from left to right to fill lMax , samely traverse from right to left to fill rMax.
	- Finally, calculate the trapped water at each position using the formula: trapped water = min(lMax[i], rMax[i]) - height[i]
	- TC: O(n), SC: O(n)
4. Stack:
	- Use stack to keep track of the indices of the bars.
	- Traverse the height array, and for each bar, pop the stack until the current bar is shorter than the top of the stack.
	- For each popped bar, calculate the trapped water using the distance between current index and the index of the new top of
	the stack, and the height difference between the current bar and the popped bar.
	- TC: O(n), SC: O(n)
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

func trap_two_pointers(height []int) int {
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

func trap_prefix_suffix(height []int) int {
	n := len(height)
	if n == 0 {
		return 0
	}
	lMax := make([]int, n)
	rMax := make([]int, n)

	lMax[0] = height[0]
	for i := 1; i < n; i++ {
		lMax[i] = max(lMax[i-1], height[i])
	}

	rMax[n-1] = height[n-1]
	for i := n - 2; i >= 0; i-- {
		rMax[i] = max(rMax[i+1], height[i])
	}

	res := 0
	for i := 0; i < n; i++ {
		res += min(lMax[i], rMax[i]) - height[i]
	}
	return res
}

func trap_stack(height []int) int {
	n := len(height)
	if n == 0 {
		return 0
	}
	type pair struct {
		height, index int
	}
	stack := make([]pair, 0)
	res := 0

	for i := 0; i < n; i++ {
		for len(stack) > 0 && height[i] > stack[len(stack)-1].height {
			popped := stack[len(stack)-1]
			stack = stack[:len(stack)-1]
			if len(stack) > 0 {
				left := stack[len(stack)-1]
				width := i - left.index - 1
				h := min(height[i], left.height) - popped.height
				res += width * h
			}
		}
		stack = append(stack, pair{height: height[i], index: i})
	}
	return res
}

func main() {
	height := []int{0, 1, 0, 2, 1, 0, 1, 3, 2, 1, 2, 1}
	println(trap_stack(height)) // 6
}
