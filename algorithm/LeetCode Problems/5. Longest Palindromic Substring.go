package main

// 5. Longest Palindromic Substring
// https://leetcode.com/problems/longest-palindromic-substring/

/*
Thinking process:
1. Brute Force:
   - Generate all possible substrings of the input string.
   - Check if each substring is a palindrome.
   - Keep track of the longest palindromic substring found.
   - TC:O(n^3) (O(n^2) for generating substrings and O(n) for checking palindrome)

2. two pointers:
   - Expand around the center of the palindrome.
   - For each character in the string, consider it as the center of a palindrome and expand outwards to check for palindromic substrings.
   - This can be done for both odd-length and even-length palindromes.
   - TC: O(n^2) (O(n) for each center and O(n) for expansion)
*/

func longestPalindrome_brute_force(s string) string {
	n := len(s)
	if n == 0 {
		return ""
	}

	res := ""
	for i := 0; i < n; i++ {
		for j := i; j < n; j++ {
			substr := s[i : j+1]
			if isPalindrome(substr) && len(substr) > len(res) {
				res = substr
			}
		}
	}
	return res
}

func isPalindrome(s string) bool {
	left, right := 0, len(s)-1
	for left < right {
		if s[left] != s[right] {
			return false
		}
		left++
		right--
	}
	return true
}

func longestPalindrome(s string) string {
	n := len(s)
	if n == 0 {
		return ""
	}

	res := ""
	for i := 0; i < n; i++ {
		// Odd length palindromes
		left, right := i, i
		for left >= 0 && right < n && s[left] == s[right] {
			if right-left+1 > len(res) {
				res = s[left : right+1]
			}
			left--
			right++
		}

		// Even length palindromes
		left, right = i, i+1
		for left >= 0 && right < n && s[left] == s[right] {
			if right-left+1 > len(res) {
				res = s[left : right+1]
			}
			left--
			right++
		}
	}
	return res
}

func main() {
	s := "babab"
	result := longestPalindrome(s)
	println(result) // Output: "b"
}
