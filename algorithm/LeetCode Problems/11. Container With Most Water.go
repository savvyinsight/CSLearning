package main

// 11. Container With Most Water
// https://leetcode.com/problems/container-with-most-water/

/*
Thinking process:
1.brute force: O(n^2)
2.two pointers: O(n)
*/
func maxArea(height []int) int {
	res := 0
	for i := 0; i < len(height); i++ {
		for j := i + 1; j < len(height); j++ {
			res = max(res, min(height[i], height[j])*(j-i))
		}
	}
	return res
}

func maxArea_twopointer(height []int) int {
	res, l, r := 0, 0, len(height)-1
	for l < r {
		res = max(res, min(height[l], height[r])*(r-l))
		if height[l] < height[r] {
			l++
		} else {
			r--
		}
	}
	return res
}

func main() {
	height := []int{1, 7, 2, 5, 4, 7, 3, 6}
	println(maxArea(height)) // Output: 36
}
