package main

import "strconv"

// 271. Encode and Decode Strings
// https://leetcode.com/problems/encode-and-decode-strings/

/*
Thinking process:
1. Need to encode strings into a single string and decodde it back to the original array of strings.
2. Use a delimiter to separate the strings. However, we need to ensure that the delimiter does not appear in the original strings.
3. Use a length prefix to indicate the length of each string, followed by a special character (e.g.,'#') to separate the length
from the string itself. Then we can easily decode the string by reading the length prefix and extracting the corresponding substring.
4. For example, if we have the string "hello" , we can encode it ad "5#hello" where '5' is the length of the string
and '#' is the delimiter. if we have multiple strings, we can concatenate them together,e.g, "5#hello5#world" OR
"5,5#hello,world"(using ',' as a delimiter between different strings).

TC:O(n),SC:O(n+m) where each encode and decode functions calls.
*/

type Solution struct{}

func (s *Solution) Encode(strs []string) string {
	if len(strs) == 0 {
		return ""
	}
	res := ""
	for _, str := range strs {
		res += strconv.Itoa(len(str)) + "#" + str
	}
	return res
}

func (s *Solution) Decode(str string) []string {
	if len(str) == 0 {
		return []string{}
	}
	res := []string{}
	i := 0
	for i < len(str) {
		j := i
		for str[j] != '#' {
			j++
		}
		length, _ := strconv.Atoi(str[i:j])
		j++ // move past '#'
		res = append(res, str[j:j+length])
		i = j + length
	}
	return res
}

func main() {
	s := &Solution{}
	encoded := s.Encode([]string{"hello", "world"})
	println(encoded) // "5#hello5#world"

	decoded := s.Decode(encoded)
	println(decoded[0])
	println(decoded[1])
}
