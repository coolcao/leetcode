package main

/*
Given a binary tree, find its maximum depth.

The maximum depth is the number of nodes along the longest path from the root node down to the farthest leaf node.

Note: A leaf is a node with no children.

Example:

Given binary tree [3,9,20,null,null,15,7],

    3
   / \
  9  20
    /  \
   15   7
return its depth = 3.
*/

// TreeNode 节点
type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func visit(root *TreeNode, depth int) int {
	if root == nil {
		return depth
	}

	left := visit(root.Left, depth+1)
	right := visit(root.Right, depth+1)
	if left > right {
		return left
	}
	return right

}

func maxDepth(root *TreeNode) int {
	if root == nil {
		return 0
	}
	return visit(root, 0)
}
