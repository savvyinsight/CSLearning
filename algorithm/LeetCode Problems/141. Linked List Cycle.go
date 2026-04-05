package main

// 141. Linked List Cycle
// https://leetcode.com/problems/linked-list-cycle/

/*
Thinking process:
1. Fast and slow pointer approach:
   - slow and fast poninters start at the head of the linked list.
   - slow pointer moves one step, fast pointer moves two steps.
   - If there is a cycle, the fast pointer will eventually meet the slow pointer.
   - If there is no cycle, the fast pointer will reach the end of the list.
   - TC:O(n), SC:O(1)
2. Hash set(O(1) lookup):
   - Traverse the linked list and store each visited node in a hash set.
   - If we encounter a node that is already in  the hash set, it means there is a cycle. otherwise, if we encounter null pointer,
   it means there is no cycle.
   - TC:O(n), SC:O(n)
*/

// Definition for singly-linked list.
type ListNode struct {
	Val  int
	Next *ListNode
}

func hasCycle(head *ListNode) bool {
	if head == nil {
		return false
	}

	slow, fast := head, head

	for fast != nil && fast.Next != nil {
		slow = slow.Next
		fast = fast.Next.Next

		if slow == fast {
			return true
		}
	}

	return false
}

func hasCycleHashSet(head *ListNode) bool {
	visited := make(map[*ListNode]bool)
	for current := head; current != nil; current = current.Next {
		if visited[current] {
			return true
		}
		visited[current] = true
	}
	return false
}

func main() {
	n1 := &ListNode{Val: 3}
	n2 := &ListNode{Val: 2}
	n3 := &ListNode{Val: 0}
	n4 := &ListNode{Val: -4}

	n1.Next = n2
	n2.Next = n3
	n3.Next = n4
	n4.Next = n2

	println(hasCycleHashSet(n1)) // Output: true
}
