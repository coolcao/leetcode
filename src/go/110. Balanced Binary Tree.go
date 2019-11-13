package main

import "fmt"

/*
Given a binary tree, determine if it is height-balanced.

For this problem, a height-balanced binary tree is defined as:

a binary tree in which the depth of the two subtrees of every node never differ by more than 1.

Example 1:

Given the following tree [3,9,20,null,null,15,7]:

    3
   / \
  9  20
    /  \
   15   7
Return true.

Example 2:

Given the following tree [1,2,2,3,3,null,null,4,4]:

       1
      / \
     2   2
    / \
   3   3
  / \
 4   4
Return false.
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

// 判断某个节点是否平衡，并记录其节点深度
func balanced(root *TreeNode, depth *int) bool {
	if root == nil {
		*depth = 0
		return true
	}
	left, right := 0, 0
	leftBalanced, rightBalanced := balanced(root.Left, &left), balanced(root.Right, &right)

	if leftBalanced && rightBalanced {
		gap := left - right
		if gap <= 1 && gap >= -1 {
			if left > right {
				*depth = left + 1
			} else {
				*depth = right + 1
			}
			return true
		}
	}
	return false
}
func isBalanced(root *TreeNode) bool {
	if root == nil {
		return true
	}
	depth := 0
	return balanced(root, &depth)
}

func main() {
	nums := []int{1, 2, 2, 3, 0, 0, 3, 4, 0, 0, 4}
	root := createTreeFromArray(nums)
	is := isBalanced(root)
	fmt.Printf("%v\n", is)
}
