package main

/*
Thinking process:
1. DFS
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

func main() {

}
