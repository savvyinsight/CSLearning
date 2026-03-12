package main

// 20. Valid Parentheses
// https://leetcode.com/problems/valid-parentheses/

/*
Thinking process:
- Note:Open brackets must be closed in the correct order. "()[]{}" is valid but "(]" and "([)]" are not.
1. Brute Force
   - Iterate through the string and remove pairs of valid parentheses until the string is empty or no more pairs can be removed.
   - Time Complexity: O(n^2) in the worst case, as we may need to scan the string multiple times.
   - Space Complexity: O(n) for the string manipulation.

2. Stack
   - Use a stack to keep track of opening parentheses. For each closing parenthesis, check if it matches the top of the stack.
   - Time Complexity: O(n) as we traverse the string once.
   - Space Complexity: O(n) in the worst case if all characters are opening parentheses.
*/

// approach 1:Brute force
func isValid_1(s string) bool {
	for {
		newS := removePair(s)
		if newS == s { // if there are no more pairs to remove, we can break the loop
			break
		}
		s = newS
	}
	return s == ""
}

func removePair(s string) string {
	for i := 0; i < len(s)-1; i++ {
		if s[i:i+2] == "()" || s[i:i+2] == "[]" || s[i:i+2] == "{}" {
			return s[:i] + s[i+2:]
		}
	}
	return s
}

// approach 2: Stack
func isValid_2(s string) bool {
	stack := []byte{}
	pairs := map[byte]byte{
		')': '(',
		']': '[',
		'}': '{',
	}

	for i := 0; i < len(s); i++ {
		char := s[i]
		if _, exists := pairs[char]; exists {
			if len(stack) == 0 || stack[len(stack)-1] != pairs[char] {
				return false
			}
			stack = stack[:len(stack)-1]
		} else {
			stack = append(stack, char)
		}
	}
	return len(stack) == 0
}

func main() {
	println(isValid_2("()"))
	println(isValid_2("(][]{}"))
	println(isValid_2("(]"))
	println(isValid_2("([)]"))
	println(isValid_2("{[]}"))
	println(isValid_2("([)]")) // false
}
