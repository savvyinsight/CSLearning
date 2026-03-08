package main

import "sort"

// 242. Valid Anagram
// https://leetcode.com/problems/valid-anagram/

/*
Thinking process:
1. Sorting
   - We can sort both strings and compare them. If they are the same, then they are anagrams.
   - Time complexity: O(n log n) due to sorting
   - Space complexity: O(1) if we sort in place, otherwise O(n)
2. Hash map
   - We can use a hash map to count the frequency of each character in the first string and then decrement the counts based on the second string. If all counts are zero at the end, they are anagrams.
   - Time complexity: O(n)
   - Space complexity: O(1) if we assume a fixed character set (e.g., ASCII), otherwise O(n)
*/

// approach 1: sorting
func isAnagram_1(s string, t string) bool {
	if len(s) != len(t) {
		return false
	}
	sBytes := []byte(s)
	tBytes := []byte(t)
	sort.Slice(sBytes, func(i, j int) bool {
		return sBytes[i] < sBytes[j]
	})
	sort.Slice(tBytes, func(i, j int) bool {
		return tBytes[i] < tBytes[j]
	})
	for i := 0; i < len(sBytes); i++ {
		if sBytes[i] != tBytes[i] {
			return false
		}
	}
	return true
}

// approach 2: hash map
func isAnagram_2(s string, t string) bool {
	if len(s) != len(t) {
		return false
	}
	count := make(map[rune]int)
	for _, char := range s {
		count[char]++
	}
	for _, char := range t {
		count[char]--
		if count[char] < 0 {
			return false
		}
	}
	return true
}

func main() {
	s := "anagram"
	t := "nagaram"
	println(isAnagram_1(s, t)) // true
	println(isAnagram_2(s, t)) // true

	s = "rat"
	t = "car"
	println(isAnagram_1(s, t)) // false
	println(isAnagram_2(s, t)) // false
}
