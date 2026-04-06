package main

// 19. Remove Nth Node From End of List
// https://leetcode.com/problems/remove-nth-node-from-end-of-list/
/*
Thinking process:
1. Two passes:
	- First pass to calculate the length of the list.
	- Second pass to find the (L-n)th node and remove the nth node from the end.
    - TC: O(N), SC: O(1)
2. One pass with two pointers:
    - Two pointers, fast and slow, both start at the head.
	- Move the fast pointer n steps ahead.
	- Then move both pointers one step at a time until the fast pointer reaches the end of the list.
    - TC: O(N), SC: O(1)
3. One pass with dummy node:
	- Create a dummy node that points to the head of the list.
	- Use two pointers, fast and slow, both start at the dummy node.
	- Move the fast pointer n+1 steps ahead.
	- Then move both pointers one step at a time until the fast pointer reaches the end of the list.
	- TC: O(N), SC: O(1)

4. Recursive approach:
    - Define a recursive function that returns the index of the current node from the end of the list.
	- When the index equals n, remove the current node.
	- TC: O(N), SC: O(N) due to recursion stack.
*/

// Definition for singly-linked list.
type ListNode struct {
	Val  int
	Next *ListNode
}

// two passes
func removeNthFromEndTwoPasses(head *ListNode, n int) *ListNode {
	if head == nil {
		return nil
	}

	cnt := 0
	for curr := head; curr != nil; curr = curr.Next {
		cnt++
	}

	if cnt == n {
		return head.Next
	}

	curr := head
	for i := 0; i < cnt-n-1; i++ {
		curr = curr.Next
	}
	curr.Next = curr.Next.Next

	return head
}

// one pass with two pointers
func removeNthFromEnd(head *ListNode, n int) *ListNode {
	if head == nil {
		return nil
	}

	fast, slow := head, head

	// Move the fast pointer n steps ahead
	for i := 0; i < n; i++ {
		fast = fast.Next
	}

	// If fast is nil, it means we need to remove the head node
	if fast == nil {
		return head.Next
	}

	// Move both pointers until fast reaches the end of the list
	for fast.Next != nil {
		fast = fast.Next
		slow = slow.Next
	}

	// Remove the nth node from the end
	slow.Next = slow.Next.Next

	return head
}

// one pass with dummy node
func removeNthFromEndWithDummy(head *ListNode, n int) *ListNode {
	dummy := &ListNode{Next: head}
	fast, slow := head, dummy

	for i := 0; i < n; i++ {
		fast = fast.Next
	}

	for fast != nil {
		fast = fast.Next
		slow = slow.Next
	}

	slow.Next = slow.Next.Next

	return dummy.Next
}

// recursive approach
func removeNthFromEndRecursive(head *ListNode, n int) *ListNode {
	var remove func(node *ListNode) int
	remove = func(node *ListNode) int {
		if node == nil {
			return 0
		}

		index := remove(node.Next) + 1

		if index == n+1 {
			node.Next = node.Next.Next
		}

		return index
	}

	dummy := &ListNode{Next: head}
	remove(dummy)
	return dummy.Next
}

func main() {
	// 1 -> 2 -> 3 -> 4 -> 5, n = 2
	head := &ListNode{Val: 1}
	head.Next = &ListNode{Val: 2}
	head.Next.Next = &ListNode{Val: 3}
	head.Next.Next.Next = &ListNode{Val: 4}
	head.Next.Next.Next.Next = &ListNode{Val: 5}

	n := 2
	newHead := removeNthFromEnd(head, n)

	// Print the modified list
	current := newHead
	for current != nil {
		print(current.Val, " ")
		current = current.Next
	}

}
