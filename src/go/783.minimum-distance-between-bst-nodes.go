/*
 * @lc app=leetcode id=783 lang=golang
 *
 * [783] Minimum Distance Between BST Nodes
 *
 * https://leetcode.com/problems/minimum-distance-between-bst-nodes/description/
 *
 * algorithms
 * Easy (51.76%)
 * Likes:    659
 * Dislikes: 190
 * Total Accepted:    61K
 * Total Submissions: 116.5K
 * Testcase Example:  '[4,2,6,1,3,null,null]'
 *
 * Given a Binary Search Tree (BST) with the root node root, return the minimum
 * difference between the values of any two different nodes in the tree.
 *
 * Example :
 *
 *
 * Input: root = [4,2,6,1,3,null,null]
 * Output: 1
 * Explanation:
 * Note that root is a TreeNode object, not an array.
 *
 * The given tree [4,2,6,1,3,null,null] is represented by the following
 * diagram:
 *
 * ⁠         4
 * ⁠       /   \
 * ⁠     2      6
 * ⁠    / \
 * ⁠   1   3
 *
 * while the minimum difference in this tree is 1, it occurs between node 1 and
 * node 2, also between node 3 and node 2.
 *
 *
 * Note:
 *
 *
 * The size of the BST will be between 2 and 100.
 * The BST is always valid, each node's value is an integer, and each node's
 * value is different.
 * This question is the same as 530:
 * https://leetcode.com/problems/minimum-absolute-difference-in-bst/
 *
 *
 */
package main

import "fmt"

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func main() {
	root := new(TreeNode)
	root.Val = 4
	node1 := new(TreeNode)
	node1.Val = 2
	node2 := new(TreeNode)
	node2.Val = 6
	node3 := new(TreeNode)
	node3.Val = 1
	node4 := new(TreeNode)
	node4.Val = 3

	root.Left = node1
	root.Right = node2

	node1.Left = node3
	node1.Right = node4

	diff := minDiffInBST(root)
	fmt.Printf("min diff : %v\n", diff)
}

// @lc code=start

var values []int

func visit(root *TreeNode) {
	if root != nil {
		visit(root.Left)
		values = append(values, root.Val)
		visit(root.Right)
	}
}

func minDiffInBST(root *TreeNode) int {
	values = []int{}
	visit(root)
	min := values[1] - values[0]

	for i := 2; i < len(values); i++ {
		diff := values[i] - values[i-1]
		if diff < min {
			min = diff
		}
	}

	return min
}

// @lc code=end
