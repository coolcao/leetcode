/*
 * @lc app=leetcode id=437 lang=golang
 *
 * [437] Path Sum III
 *
 * https://leetcode.com/problems/path-sum-iii/description/
 *
 * algorithms
 * Easy (45.15%)
 * Likes:    2931
 * Dislikes: 227
 * Total Accepted:    161.7K
 * Total Submissions: 357.8K
 * Testcase Example:  '[10,5,-3,3,2,null,11,3,-2,null,1]\n8'
 *
 * You are given a binary tree in which each node contains an integer value.
 *
 * Find the number of paths that sum to a given value.
 *
 * The path does not need to start or end at the root or a leaf, but it must go
 * downwards
 * (traveling only from parent nodes to child nodes).
 *
 * The tree has no more than 1,000 nodes and the values are in the range
 * -1,000,000 to 1,000,000.
 *
 * Example:
 *
 * root = [10,5,-3,3,2,null,11,3,-2,null,1], sum = 8
 *
 * ⁠     10
 * ⁠    /  \
 * ⁠   5   -3
 * ⁠  / \    \
 * ⁠ 3   2   11
 * ⁠/ \   \
 * 3  -2   1
 *
 * Return 3. The paths that sum to 8 are:
 *
 * 1.  5 -> 3
 * 2.  5 -> 2 -> 1
 * 3. -3 -> 11
 *
 *
 */

package main

import "fmt"

// @lc code=start
type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func createNode(num interface{}) *TreeNode {
	if num == nil {
		return nil
	}
	node := new(TreeNode)
	node.Val = num.(int)
	node.Left, node.Right = nil, nil
	return node
}
func createTreeFromArray(nums []interface{}) *TreeNode {
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

func helper(node *TreeNode, sum int, count *int) {
	if node == nil {
		return
	}
	if node.Val == sum {
		*count++
	}
	helper(node.Left, sum-node.Val, count)
	helper(node.Right, sum-node.Val, count)
}
func path(root *TreeNode, sum int, count *int) {
	helper(root, sum, count)
	if root.Left != nil {
		path(root.Left, sum, count)
	}
	if root.Right != nil {
		path(root.Right, sum, count)
	}
}
func pathSum(root *TreeNode, sum int) int {
	count := 0
	if root == nil {
		return count
	}

	path(root, sum, &count)

	return count
}

func main() {
	nums := []interface{}{10, 5, -3, 3, 2, nil, 11, 3, -2, nil, 1}
	sum := 6
	root := createTreeFromArray(nums)
	c := pathSum(root, sum)
	fmt.Printf("%d\n", c)
}

// @lc code=end
