package main

import "fmt"

// 206. Reverse Linked List
// https://leetcode.com/problems/reverse-linked-list/

/*	Thinking process:
1. Iterative approach:
   - Use three pointers: prev, curr, and next.
   - Iterate through the linked list, reversing the pointers at each step.
   - Time complexity: O(n)
   - Space complexity: O(1)

2. Recursive approach:
   - Recursively reverse the linked list and adjust the pointers accordingly.
   - Time complexity: O(n)
   - Space complexity: O(n) due to recursive call stack
*/

// Definition for singly-linked list.
type ListNode struct {
	Val  int
	Next *ListNode
}

// approach 1: Iterative approach
func reverseList_1(head *ListNode) *ListNode {
	var prev *ListNode
	curr := head

	for curr != nil {
		next := curr.Next
		curr.Next = prev
		prev = curr
		curr = next
	}
	return prev
}

// approach 2: Recursive approach
func reverseList_2(head *ListNode) *ListNode {
	if head == nil || head.Next == nil {
		return head
	}

	newHead := reverseList_2(head.Next)
	head.Next.Next = head
	head.Next = nil

	return newHead
}

func main() {
	// [0,1,2,3]
	n1 := &ListNode{Val: 0}
	n2 := &ListNode{Val: 1}
	n3 := &ListNode{Val: 2}
	n4 := &ListNode{Val: 3}

	n1.Next = n2
	n2.Next = n3
	n3.Next = n4

	reversed := reverseList_2(n1)
	for reversed != nil {
		fmt.Print(reversed.Val, " ")
		reversed = reversed.Next
	}
}
