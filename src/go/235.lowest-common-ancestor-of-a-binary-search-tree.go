/*
 * @lc app=leetcode id=235 lang=golang
 *
 * [235] Lowest Common Ancestor of a Binary Search Tree
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
