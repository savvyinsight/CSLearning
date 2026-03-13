package main

import "fmt"

// 21. Merge Two Sorted Lists
// https://leetcode.com/problems/merge-two-sorted-lists/

/*	Thinking process:
1. Iterative approach:
   - Use two pointers to traverse both lists and merge them into a new list.
   - Time complexity: O(n + m), where n and m are the lengths of the two lists.
   - Space complexity: O(1) for the new list (excluding the input lists).

2. Recursive approach:
   - Recursively merge the two lists by comparing their head nodes.
   - Time complexity: O(n + m)
   - Space complexity: O(n + m) due to recursive call stack
*/

// Definition for singly-linked list.
type ListNode struct {
	Val  int
	Next *ListNode
}

// approach 1: Iterative approach
func mergetTwoLists(l1 *ListNode, l2 *ListNode) *ListNode {
	dummy := &ListNode{}
	curr := dummy

	for l1 != nil && l2 != nil {
		if l1.Val < l2.Val {
			curr.Next = l1
			l1 = l1.Next
		} else {
			curr.Next = l2
			l2 = l2.Next
		}
		curr = curr.Next
	}

	if l1 != nil {
		curr.Next = l1
	} else {
		curr.Next = l2
	}

	return dummy.Next
}

// approach 2: Recursive approach
func mergeTwoListsRecursive(l1 *ListNode, l2 *ListNode) *ListNode {
	if l1 == nil {
		return l2
	}
	if l2 == nil {
		return l1
	}

	if l1.Val < l2.Val {
		l1.Next = mergeTwoListsRecursive(l1.Next, l2)
		return l1
	} else {
		l2.Next = mergeTwoListsRecursive(l1, l2.Next)
		return l2
	}
}

func main() {
	// 1->2>4, 1->3->4
	n1 := &ListNode{Val: 1}
	n2 := &ListNode{Val: 2}
	n3 := &ListNode{Val: 4}
	n1.Next = n2
	n2.Next = n3

	m1 := &ListNode{Val: 1}
	m2 := &ListNode{Val: 3}
	m3 := &ListNode{Val: 4}
	m1.Next = m2
	m2.Next = m3

	merged := mergeTwoListsRecursive(n1, m1)
	for merged != nil {
		fmt.Print(merged.Val, " ")
		merged = merged.Next
	}
}
