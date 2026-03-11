package main

// 125. Valid Palindrome
// https://leetcode.com/problems/valid-palindrome/

/*
Thinking process:
1. Two pointers
	- Time complexity: O(n)
	- Space complexity: O(1)
2. Reverse string and compare
	- Time complexity: O(n)
	- Space complexity: O(n)
*/

// approach 1: two pointers
func isPalindrome_1(s string) bool {
	l, r := 0, len(s)-1

	for l < r {
		for l < r && !isAlphaNum(s[l]) {
			l++
		}
		for l < r && !isAlphaNum(s[r]) {
			r--
		}
		if toLower(s[l]) != toLower(s[r]) {
			return false
		}
		l++
		r--
	}
	return true
}

func isAlphaNum(c byte) bool {
	return (c >= 'a' && c <= 'z') ||
		(c >= 'A' && c <= 'Z') ||
		(c >= '0' && c <= '9')
}

func toLower(c byte) byte {
	if c >= 'A' && c <= 'Z' {
		return c + 'a' - 'A'
	}
	return c
}

// approach 2: reverse string and compare
func isPalindrome_2(s string) bool {
	var filtered []byte
	for i := 0; i < len(s); i++ {
		if isAlphaNum(s[i]) {
			filtered = append(filtered, toLower(s[i]))
		}
	}

	n := len(filtered)
	for i := 0; i < n/2; i++ {
		if filtered[i] != filtered[n-1-i] {
			return false
		}
	}
	return true
}

func main() {
	s := "A man, a plan, a canal: Panama"
	println(isPalindrome_1(s)) // true
	println(isPalindrome_2(s)) // true
}
