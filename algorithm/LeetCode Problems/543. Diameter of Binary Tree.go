package main

/*
Thinking process:
1.Brute Force
   - For each node, calculate the depth of left and right subtree, and update the diameter.
   - TC:O(n^2),SC:O(n) for recursion stack
2.DFS
   - Compute the diameter and height in a single DFS traversal. Use a global variable to
   Keep track of the maximum diameter found so far.
   - TC:O(n),SC:O(n)
*/

// Definition for a binary tree node.
type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func diameterOfBinaryTree(root *TreeNode) int {
	if root == nil {
		return 0
	}
	lHeight := maxHeight(root.Left)
	rHeight := maxHeight(root.Right)
	diameter := lHeight + rHeight
	leftDiameter := diameterOfBinaryTree(root.Left)
	rightDiameter := diameterOfBinaryTree(root.Right)

	return max(diameter, max(leftDiameter, rightDiameter))
}

func maxHeight(node *TreeNode) int {
	if node == nil {
		return 0
	}
	leftHeight := maxHeight(node.Left)
	rightHeight := maxHeight(node.Right)
	return 1 + max(leftHeight, rightHeight)
}

func diameterOfBinaryTree_dfs(root *TreeNode) int {
	var maxDiameter int
	var dfs func(node *TreeNode) int

	dfs = func(node *TreeNode) int {
		if node == nil {
			return 0
		}
		leftHeight := dfs(node.Left)
		rightHeight := dfs(node.Right)

		// Update the diameter at this node
		maxDiameter = max(maxDiameter, leftHeight+rightHeight)

		// Return the height of the current node
		return 1 + max(leftHeight, rightHeight)
	}

	dfs(root)
	return maxDiameter
}

func main() {
	root := &TreeNode{Val: 1}
	root.Left = &TreeNode{Val: 2}
	root.Right = &TreeNode{Val: 3}
	root.Left.Left = &TreeNode{Val: 4}
	root.Left.Right = &TreeNode{Val: 5}

	result := diameterOfBinaryTree(root)
	println(result) // Output: 3

}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
