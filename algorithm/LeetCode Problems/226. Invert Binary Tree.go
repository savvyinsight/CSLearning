package main

/*
Thinking process:
1. DFS
	- swap the left and right child of the current node
	- recursively call the function on the left and right child
	TC :O(n) where n is the number of nodes in the tree
	SC :O(h)
2. BFS
	- use a queue to traverse the tree level by level
	- swap the left and right child of the current node
	TC:O(n),SC:O(n) in the worst case when the tree unbalanced
*/

// Definition for a binary tree node.
type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func invertTree_dfs(root *TreeNode) *TreeNode {
	if root == nil {
		return nil
	}
	root.Left, root.Right = root.Right, root.Left
	invertTree_dfs(root.Left)
	invertTree_dfs(root.Right)
	return root
}

func invertTree_bfs(root *TreeNode) *TreeNode {
	if root == nil {
		return nil
	}
	q := []*TreeNode{root}
	for len(q) > 0 {
		node := q[0]
		q = q[1:]
		node.Left, node.Right = node.Right, node.Left
		if node.Left != nil {
			q = append(q, node.Left)
		}
		if node.Right != nil {
			q = append(q, node.Right)
		}
	}
	return root
}

func verify(root *TreeNode) {
	if root == nil {
		return
	}
	println(root.Val)
	verify(root.Left)
	verify(root.Right)
}

func main() {
	root := &TreeNode{Val: 4}
	root.Left = &TreeNode{Val: 2}
	root.Right = &TreeNode{Val: 7}
	root.Left.Left = &TreeNode{Val: 1}
	root.Left.Right = &TreeNode{Val: 3}
	root.Right.Left = &TreeNode{Val: 6}
	root.Right.Right = &TreeNode{Val: 9}

	invertTree_bfs(root)
	// traverse the tree to verify the result
	verify(root)
}
