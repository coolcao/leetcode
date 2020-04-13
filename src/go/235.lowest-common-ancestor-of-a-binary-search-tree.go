/*
 * @lc app=leetcode id=235 lang=golang
 *
 * [235] Lowest Common Ancestor of a Binary Search Tree
 *
 * https://leetcode.com/problems/lowest-common-ancestor-of-a-binary-search-tree/description/
 *
 * algorithms
 * Easy (48.47%)
 * Likes:    1730
 * Dislikes: 95
 * Total Accepted:    371.9K
 * Total Submissions: 766.7K
 * Testcase Example:  '[6,2,8,0,4,7,9,null,null,3,5]\n2\n8'
 *
 * Given a binary search tree (BST), find the lowest common ancestor (LCA) of
 * two given nodes in the BST.
 *
 * According to the definition of LCA on Wikipedia: “The lowest common ancestor
 * is defined between two nodes p and q as the lowest node in T that has both p
 * and q as descendants (where we allow a node to be a descendant of itself).”
 *
 * Given binary search tree:  root = [6,2,8,0,4,7,9,null,null,3,5]
 *
 *
 *
 * Example 1:
 *
 *
 * Input: root = [6,2,8,0,4,7,9,null,null,3,5], p = 2, q = 8
 * Output: 6
 * Explanation: The LCA of nodes 2 and 8 is 6.
 *
 *
 * Example 2:
 *
 *
 * Input: root = [6,2,8,0,4,7,9,null,null,3,5], p = 2, q = 4
 * Output: 2
 * Explanation: The LCA of nodes 2 and 4 is 2, since a node can be a descendant
 * of itself according to the LCA definition.
 *
 *
 *
 *
 * Note:
 *
 *
 * All of the nodes' values will be unique.
 * p and q are different and both values will exist in the BST.
 *
 *
 */

// @lc code=start
package main

// TreeNode 树节点
type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

// 创建数节点
func createNode(num interface{}) *TreeNode {
	if num == nil {
		return nil
	}
	node := new(TreeNode)
	node.Val = num.(int)
	node.Left, node.Right = nil, nil
	return node
}

// 由数组构建二叉树
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

func lowestCommonAncestor(root, p, q *TreeNode) *TreeNode {
	if root == nil {
		return nil
	}
	if root == p || root == q {
		return root
	}
	// 一般二叉树的题目都优先考虑使用递归进行解决
	// 这里我们要求两个节点的最低公共祖先节点，我们可依次在root的左子树合右子树上查找
	// 如果能在左子树或右子树上找到，那么最低公共祖先节点就在左子树或右子树
	// 否则便是根节点
	left, right := lowestCommonAncestor(root.Left, p, q), lowestCommonAncestor(root.Right, p, q)
	if left == nil {
		return right
	}
	if right == nil {
		return left
	}
	return root

}

// @lc code=end
