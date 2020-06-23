/*
 * @lc app=leetcode id=687 lang=golang
 *
 * [687] Longest Univalue Path
 *
 * https://leetcode.com/problems/longest-univalue-path/description/
 *
 * algorithms
 * Easy (35.39%)
 * Likes:    1610
 * Dislikes: 449
 * Total Accepted:    89.2K
 * Total Submissions: 248.4K
 * Testcase Example:  '[5,4,5,1,1,5]'
 *
 * Given a binary tree, find the length of the longest path where each node in
 * the path has the same value. This path may or may not pass through the
 * root.
 *
 * The length of path between two nodes is represented by the number of edges
 * between them.
 *
 *
 *
 * Example 1:
 *
 * Input:
 *
 *
 * ⁠             5
 * ⁠            / \
 * ⁠           4   5
 * ⁠          / \   \
 * ⁠         1   1   5
 *
 *
 * Output: 2
 *
 *
 *
 * Example 2:
 *
 * Input:
 *
 *
 * ⁠             1
 * ⁠            / \
 * ⁠           4   5
 * ⁠          / \   \
 * ⁠         4   4   5
 *
 *
 * Output: 2
 *
 *
 *
 * Note: The given binary tree has not more than 10000 nodes. The height of the
 * tree is not more than 1000.
 *
 */
package main

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

// @lc code=start

func pathLength(node *TreeNode, length *int) int {
	if node == nil {
		return 0
	}
	left, right := 0, 0
	if node.Left != nil {
		d := pathLength(node.Left, length)
		if node.Left.Val == node.Val {
			left = d + 1
		}
	}
	if node.Right != nil {
		d := pathLength(node.Right, length)
		if node.Right.Val == node.Val {
			right = d + 1
		}
	}
	if left+right > *length {
		*length = left + right
	}
	if left > right {
		return left
	}
	return right
}
func longestUnivaluePath(root *TreeNode) int {
	if root == nil {
		return 0
	}
	length := 0
	pathLength(root, &length)
	return length
}

// @lc code=end
