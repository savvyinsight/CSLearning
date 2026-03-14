package main

// 104. Maximum Depth of Binary Tree
// https://leetcode.com/problems/maximum-depth-of-binary-tree/

/*
Thinking process:
1. Recursive dfs
    - compute the maximum depth of a binary tree by exploring every node.
	- the depth of a tree = 1 + max(depth of left subtree, depth of right subtree)
	- base case: if the node is null, return 0
	- TC:O(n),SC:O(n) in the worst case (degenerate tree), O(log n) in the best case (balanced tree)
2. bfs
	- compute the maximum depth of a binary tree by level order traversal.
	- use a queue to keep track of nodes at each level.
	- for each level, increment the depth by 1.
	- TC:O(n),SC:O(n) in the worst case (degenerate tree), O(log n) in the best case (balanced tree)
*/

// Definition for a binary tree node.
type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

// recursive dfs
func maxDepth_1(root *TreeNode) int {
	if root == nil {
		return 0
	}
	lDepth := maxDepth_1(root.Left)
	rDepth := maxDepth_1(root.Right)
	return 1 + max(lDepth, rDepth)
}

// bfs
func maxDepth_2(root *TreeNode) int {
	if root == nil {
		return 0
	}
	q := []*TreeNode{root}
	depth := 0
	for len(q) > 0 {
		size := len(q)
		for i := 0; i < size; i++ {
			node := q[0]
			q = q[1:]
			if node.Left != nil {
				q = append(q, node.Left)
			}
			if node.Right != nil {
				q = append(q, node.Right)
			}
		}
		depth++
	}
	return depth
}

// iterative dfs
func maxDepth_3(root *TreeNode) int {
	if root == nil {
		return 0
	}
	stack := []struct {
		node  *TreeNode
		depth int
	}{{root, 1}}
	maxDepth := 0
	for len(stack) > 0 {
		curr := stack[len(stack)-1]
		stack = stack[:len(stack)-1]
		if curr.node != nil {
			maxDepth = max(maxDepth, curr.depth)
			stack = append(stack, struct {
				node  *TreeNode
				depth int
			}{curr.node.Left, curr.depth + 1})
			stack = append(stack, struct {
				node  *TreeNode
				depth int
			}{curr.node.Right, curr.depth + 1})
		}
	}
	return maxDepth
}

func main() {
	// [1,2,3,null,null,4,5]
	n1 := &TreeNode{Val: 1}
	n2 := &TreeNode{Val: 2}
	n3 := &TreeNode{Val: 3}
	n4 := &TreeNode{Val: 4}
	n5 := &TreeNode{Val: 5}
	n1.Left = n2
	n1.Right = n3
	n3.Left = n4
	n3.Right = n5

	println(maxDepth_1(n1)) // 3
	println(maxDepth_2(n1)) // 3
	println(maxDepth_3(n1))
}
