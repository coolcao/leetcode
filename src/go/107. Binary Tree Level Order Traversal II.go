package main

import (
	"fmt"
)

/*
Given a binary tree, return the bottom-up level order traversal of its nodes' values. (ie, from left to right, level by level from leaf to root).

For example:
Given binary tree [3,9,20,null,null,15,7],
    3
   / \
  9  20
    /  \
   15   7
return its bottom-up level order traversal as:
[
  [15,7],
  [9,20],
  [3]
]
*/

// "encoding/json"

// TreeNode 🌲️节点
type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func maxDepth(root *TreeNode, depth int) int {
	if root == nil {
		return depth
	}

	left := maxDepth(root.Left, depth+1)
	right := maxDepth(root.Right, depth+1)
	if left > right {
		return left
	}
	return right
}

func levelOrderBottom(root *TreeNode) [][]int {
	if root == nil {
		return nil
	}
	type QueueNode struct {
		level int
		node  *TreeNode
	}
	result := make([][]int, 0)
	queue := make([]*QueueNode, 0)
	queue = append(queue, &QueueNode{level: 0, node: root})
	for len(queue) > 0 {
		qnode := queue[0]
		node := qnode.node
		queue = queue[1:]
		if node != nil {
			//if result[qnode.level] != nil {
			if len(result) > qnode.level && result[qnode.level] != nil {
				result[qnode.level] = append(result[qnode.level], node.Val)
			} else {
				result = append(result, []int{node.Val})
			}
			queue = append(queue, &QueueNode{node: node.Left, level: qnode.level + 1})
			queue = append(queue, &QueueNode{node: node.Right, level: qnode.level + 1})
		}
	}
	tmp := make([][]int, 0)
	length := len(result)
	for idx, _ := range result {
		tmp = append(tmp, result[length-1-idx])
	}
	return tmp
}

func visit(node *TreeNode, level int, result [][]int) {
	if node == nil {
		return
	}
	length := len(result)
	sub := result[length-1-level]
	sub = append(sub, node.Val)
	result[length-1-level] = sub
	visit(node.Left, level+1, result)
	visit(node.Right, level+1, result)
}

func levelOrderBottom2(root *TreeNode) [][]int {
	// 获取最大深度
	max := maxDepth(root, 0)
	result := make([][]int, max)
	visit(root, 0, result)
	return result
}

func createNode(num int) *TreeNode {
	node := new(TreeNode)
	node.Val = num
	node.Left, node.Right = nil, nil
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
		node := createNode(nums[i])
		if i%2 == 1 {
			current.Left = node
		} else {
			current.Right = node
		}
		nodes = append(nodes, node)
		current = nodes[i/2]
	}

	return root
}

func main() {
	nums := []int{1, 2, 3, 0, 4, 5}
	root := createTreeFromArray(nums)
	result := levelOrderBottom2(root)
	fmt.Printf("%v\n", result)
	// nums = append(nums, 9)
	// fmt.Printf("%v\n", nums)
	// nums = nums[:len(nums)-1]
	// fmt.Printf("%v\n", nums)
	// nums = append(nums, 8)
	// fmt.Printf("%v\n", nums)
}
