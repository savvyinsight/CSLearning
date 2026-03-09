package main

import (
	"fmt"
	"sort"
	"strings"
)

// 49. Group Anagrams
// https://leetcode.com/problems/group-anagrams/

/*
Thinking process:
1. Sorting
	- Sort each string and use the sorted string as a key in a map
	- Group all strings that have the same sorted version together
	- Time complexity: O(n * k log k), where n is the number of strings and k is the maximum length of a string
	- Space complexity: O(n * k), where n is the number of strings and k is the maximum length of a string
2. Counting characters
	- Count the frequency of each character in the string and use the count as a key in a map
	- Group all strings that have the same character count together
	- Time complexity: O(n * k)
	- Space complexity: O(n * k)

3. Counting characters with a more compact key
	- create a hash map where each key is a 26-length tuple representing the count of each character in the string
	- Time complexity: O(n * k)
	- Space complexity: O(n * k)
*/

// approach 1: sorting
func groupAnagrams_1(strs []string) [][]string {
	res := make(map[string][]string)
	for _, s := range strs {
		sorted := sortString(s)
		res[sorted] = append(res[sorted], s)
	}
	result := make([][]string, 0, len(res))
	for _, group := range res {
		result = append(result, group)
	}
	return result
}

func sortString(s string) string {
	runes := []rune(s)
	sort.Slice(runes, func(i, j int) bool {
		return runes[i] < runes[j]
	})
	return string(runes)
}

// approach 2: counting characters
func groupAnagrams_2(strs []string) [][]string {
	res := make(map[string][]string)
	for _, s := range strs {
		count := make([]int, 26)
		for _, char := range s {
			count[char-'a']++
		}
		key := strings.Trim(strings.Join(strings.Fields(fmt.Sprint(count)), ","), "[]")
		res[key] = append(res[key], s)
	}
	result := make([][]string, 0, len(res))
	for _, group := range res {
		result = append(result, group)
	}
	return result
}

// approach 2 : counting characters with a more compact key
func groupAnagrams_3(strs []string) [][]string {
	res := make(map[string][]string)
	for _, s := range strs {
		count := make([]int, 26)
		for _, char := range s {
			count[char-'a']++
		}
		key := fmt.Sprint(count)
		res[key] = append(res[key], s)
	}
	result := make([][]string, 0, len(res))
	for _, group := range res {
		result = append(result, group)
	}
	return result
}

func main() {
	strs := []string{"eat", "tea", "tan", "ate", "nat", "bat"}
	result := groupAnagrams_3(strs)
	for _, group := range result {
		println(strings.Join(group, ", "))
	}
}
