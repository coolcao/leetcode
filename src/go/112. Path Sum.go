package main

import "fmt"

/*
Given a binary tree and a sum, determine if the tree has a root-to-leaf path such that adding up all the values along the path equals the given sum.

Note: A leaf is a node with no children.

Example:

Given the below binary tree and sum = 22,

      5
     / \
    4   8
   /   / \
  11  13  4
 /  \      \
7    2      1
return true, as there exist a root-to-leaf path 5->4->11->2 which sum is 22.
*/

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func createNode(num int) *TreeNode {
	node := new(TreeNode)
	node.Val = num
	return node
}

func createTreeFromArray(nums []int) *TreeNode {
	if len(nums) == 0 {
		return nil
	}

	nodes := []*TreeNode{}
	root := createNode(nums[0])
	nodes = append(nodes, root)
	current := root
	for i := 1; i < len(nums); i++ {
		var node *TreeNode
		if nums[i] != 0 {
			node = createNode(nums[i])
		}
		if i%2 == 1 {
			current.Left = node
		} else {
			current.Right = node
		}
		if node != nil {
			nodes = append(nodes, node)
		}
		current = nodes[i/2]
	}

	return root
}

func visit(node *TreeNode, pathLength int, sum int) bool {
	if node == nil {
		return false
	}
	if node.Left == nil && node.Right == nil {
		pathLength += node.Val
		if pathLength == sum {
			return true
		}
	}
	if node.Left == nil {
		return visit(node.Right, pathLength+node.Val, sum)
	}
	if node.Right == nil {
		return visit(node.Left, pathLength+node.Val, sum)
	}

	return visit(node.Left, pathLength+node.Val, sum) || visit(node.Right, pathLength+node.Val, sum)
}

func hasPathSum(root *TreeNode, sum int) bool {
	return visit(root, 0, sum)
}
func main() {
	nums := []int{1, 2, 3, 4, 5}
	root := createTreeFromArray(nums)
	sum := 8
	b := hasPathSum(root, sum)
	fmt.Printf("has:%v\n", b)
}
