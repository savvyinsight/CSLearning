package main

import (
	"fmt"
	"strconv"
)

// 150. Evaluate Reverse Polish Notation
// https://leetcode.com/problems/evaluate-reverse-polish-notation

/*
Thinking process:
1. stack
	- using stack for data structure
	- when input number we can add to stack. when meet operands we can pop all numbers
	then do operations.
	- TC:O(n),SC:O(n)

2. Brute Force
	- First find the operand, then take two element before, compute result, replace three tokens
	with the result.
	- Second need to repeatedly do the same operation until remain only one token ->result.
	-TC:O(n),SC:O(n)

*/

func evalRPN(tokens []string) int {
	res := 0 // maybe only one number, this will not update.
	// we added it to stack, ultimate there left only one element in stack.
	stk := []int{}
	for _, v := range tokens {
		if v != "+" && v != "-" && v != "/" && v != "*" {
			v_int, _ := strconv.Atoi(v)
			stk = append(stk, v_int)
		} else {
			right := stk[len(stk)-1]
			stk = stk[:len(stk)-1]
			left := stk[len(stk)-1]
			stk = stk[:len(stk)-1]

			switch v {
			case "+":
				res = left + right
			case "-":
				res = left - right
			case "/":
				res = left / right
			case "*":
				res = left * right
			}
			stk = append(stk, res)
		}
	}
	return stk[0]
}

func evalRPN_brute(tokens []string) int {
	for i := 0; i < len(tokens); i++ {
		v := tokens[i]
		if v == "+" || v == "-" || v == "*" || v == "/" {
			r, _ := strconv.Atoi(tokens[i-1])
			l, _ := strconv.Atoi(tokens[i-2])
			tokens_update := []string{}
			tokens_update = append(tokens_update, tokens[:i-2]...)
			switch v {
			case "+":
				// delete or flag three tokens, then replace with l+r, same as others.
				// but in Go, we can't change string, so not flaggable.
				// can we delete, if delete, needs move all elements to forward.
				// what should I do?

				// -->create new slice
				tokens_update = append(tokens_update, strconv.Itoa(l+r))
			case "-":
				tokens_update = append(tokens_update, strconv.Itoa(l-r))
			case "*":
				tokens_update = append(tokens_update, strconv.Itoa(l*r))
			case "/":
				tokens_update = append(tokens_update, strconv.Itoa(l/r))
			}

			tokens_update = append(tokens_update, tokens[i+1:]...)
			tokens = tokens_update
			i = i - 3
		}
	}
	res, _ := strconv.Atoi(tokens[0])
	return res // will remain only one token
}

func main() {
	// tokens := []string{"10", "6", "9", "3", "+", "-11", "*", "/", "*", "17", "+", "5", "+"}
	tokens := []string{"2", "1", "+", "3", "*"}
	res := evalRPN_brute(tokens)
	fmt.Println(res)
}
