package main

import (
	"fmt"
	"math"
)

/*
Given a binary tree, find its minimum depth.

The minimum depth is the number of nodes along the shortest path from the root node down to the nearest leaf node.

Note: A leaf is a node with no children.

Example:

Given binary tree [3,9,20,null,null,15,7],

    3
   / \
  9  20
    /  \
   15   7
return its minimum depth = 2.
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

func visit(root *TreeNode, level int, min *int) {
	if root.Left == nil && root.Right == nil {
		if *min > level {
			*min = level
		}
		return
	}
	if root.Left == nil {
		visit(root.Right, level+1, min)
		return
	}
	if root.Right == nil {
		visit(root.Left, level+1, min)
		return
	}

	visit(root.Left, level+1, min)
	visit(root.Right, level+1, min)

}

// 遍历二叉树，遇到叶子节点，判断当前节点的路径是否最小
func minDepth(root *TreeNode) int {
	if root == nil {
		return 0
	}
	min := math.MaxInt32
	visit(root, 1, &min)
	return min
}

type QueueNode struct {
	node  *TreeNode
	level int
}

// 采用层次遍历，如果遇到叶子节点，直接返回该叶子节点的路径长度。
func minDepth2(root *TreeNode) int {
	if root == nil {
		return 0
	}
	queue := []*QueueNode{&QueueNode{
		node:  root,
		level: 1,
	}}
	for len(queue) > 0 {
		current := queue[0]
		queue = queue[1:]
		if current.node.Left == nil && current.node.Right == nil {
			return current.level
		}
		if current.node.Right != nil {
			queue = append(queue, &QueueNode{
				level: current.level + 1,
				node:  current.node.Right,
			})
		}
		if current.node.Left != nil {
			queue = append(queue, &QueueNode{
				level: current.level + 1,
				node:  current.node.Left,
			})
		}
	}
	return 0
}

func main() {
	nums := []int{3, 9, 2, 0, 0, 15, 7, 0, 7, 0, 8, 9, 4, 2, 3}
	root := createTreeFromArray(nums)
	min := minDepth(root)
	min2 := minDepth2(root)
	fmt.Printf("min: %d\n", min)
	fmt.Printf("min2: %d\n", min2)

}
