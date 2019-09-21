package main

import (
	"fmt"
)

/*
Given an array where elements are sorted in ascending order, convert it to a height balanced BST.

For this problem, a height-balanced binary tree is defined as a binary tree in which the depth of the two subtrees of every node never differ by more than 1.

Example:

Given the sorted array: [-10,-3,0,5,9],

One possible answer is: [0,-3,9,-10,null,5], which represents the following height balanced BST:

      0
     / \
   -3   9
   /   /
 -10  5
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
func sortedArrayToBST(nums []int) *TreeNode {
	length := len(nums)
	if length == 0 {
		return nil
	}
	head := createNode(nums[len(nums)/2])
	head.Left = sortedArrayToBST(nums[:len(nums)/2])
	head.Right = sortedArrayToBST(nums[len(nums)/2+1:])
	return head
}

func main() {
	nums := []int{-10, -3, 0, 5, 9, 10}
	// root := sortedArrayToBST(nums)
	// s, _ := json.Marshal(root)
	// fmt.Printf("%v\n", string(s))
	length := len(nums)
	fmt.Printf("%v\n", nums[length/2])
	fmt.Printf("%v\n", nums[:length/2])
	fmt.Printf("%v\n", nums[length/2+1:])
}
