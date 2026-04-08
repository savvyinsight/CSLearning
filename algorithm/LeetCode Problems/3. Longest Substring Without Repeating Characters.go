package main

import "fmt"

// 3. Longest Substring Without Repeating Characters
// https://leetcode.com/problems/longest-substring-without-repeating-characters/

/*
Thinking:
1. two pointers:
- left and right
- move right until we find a duplicate character
- move left until we remove the duplicate character
- update the max length
TC:O(n) SC:O(n)

2. brute force:
- generate all substrings
- check if they have duplicate characters
- update the max length
- TC:O(n^3) SC:O(n)
*/

func lengthOfLongestSubstring_two_pointer(s string) int {
	charIndex := map[byte]int{}
	maxLen := 0
	l := 0
	for r := 0; r < len(s); r++ {
		if index, exists := charIndex[s[r]]; exists && index >= l {
			l = charIndex[s[r]] + 1
		}
		charIndex[s[r]] = r
		maxLen = max(maxLen, r-l+1)
	}
	return maxLen
}

func check(str string) bool {
	set := map[byte]bool{} // map overhead
	for i := 0; i < len(str); i++ {
		if _, ok := set[str[i]]; ok {
			return false
		}
		set[str[i]] = true
	}
	return true
}

func lengthOfLongestSubstring_brute_force(s string) int {
	maxLen := 0
	n := len(s)
	for i := 0; i < n; i++ {
		// seen := map[byte]bool{} // map creates overhead
		seen := [128]bool{} // array is faster for ASCII characters
		for j := i; j < n; j++ {
			if seen[s[j]] {
				break
			}
			seen[s[j]] = true
			maxLen = max(maxLen, j-i+1)
			// str := s[i : j+1] // O(k) for copying the substring
			// if check(str) {   // O(n)
			// 	maxLen = max(maxLen, len(str))
			// } else {
			// 	break
			// }
		}
	}
	return maxLen
}

func main() {
	s := "abcabcbb"
	fmt.Print(lengthOfLongestSubstring_brute_force(s))
}
