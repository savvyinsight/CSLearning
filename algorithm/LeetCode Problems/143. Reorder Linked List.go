package main

import "fmt"

// 143. Reorder List
// https://leetcode.com/problems/reorder-list/

/*
thinking process:
1. reverse and merge:
	- Find middle using fast and slow pointers.
	- Reverse the second half of the list.
	- Merge the two halves together.
	- TC:O(n), SC:O(1)
2. stack:
	- Push all nodes onto a stack.
	- Pop nodes from the stack and interleave them with original nodes.
	- TC:O(n), SC:O(n)

3. recursive:
	-
*/

// Definition for singly-linked list.
type ListNode struct {
	Val  int
	Next *ListNode
}

func reorderList(head *ListNode) {
	if head == nil || head.Next == nil {
		return
	}
	// find mid
	fast, slow := head, head
	for fast != nil && fast.Next != nil {
		fast = fast.Next.Next
		slow = slow.Next
	}

	// reverse second half
	var prev *ListNode
	curr := slow.Next
	slow.Next = nil
	for curr != nil {
		next := curr.Next
		curr.Next = prev
		prev = curr
		curr = next
	}

	// merge two halves
	first, second := head, prev
	for second != nil {
		t1 := first.Next
		t2 := second.Next

		first.Next = second
		second.Next = t1

		first = t1
		second = t2
	}
}

func reorderListStack(head *ListNode) {
	if head == nil || head.Next == nil {
		return
	}
	stack := []*ListNode{}
	curr := head
	for curr != nil {
		stack = append(stack, curr)
		curr = curr.Next
	}

	curr = head
	for i := len(stack) - 1; i >= len(stack)/2; i-- {
		node := stack[i]
		node.Next = curr.Next
		curr.Next = node
		curr = node.Next
	}
	curr.Next = nil
}

func main() {
	// 1 -> 2 -> 3 -> 4 -> 5
	head := &ListNode{Val: 1}
	head.Next = &ListNode{Val: 2}
	head.Next.Next = &ListNode{Val: 3}
	head.Next.Next.Next = &ListNode{Val: 4}
	head.Next.Next.Next.Next = &ListNode{Val: 5}

	reorderListStack(head)

	for head != nil {
		fmt.Print(head.Val, " ")
		head = head.Next
	}
}
