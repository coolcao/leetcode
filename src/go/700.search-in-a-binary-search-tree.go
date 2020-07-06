/*
 * @lc app=leetcode id=700 lang=golang
 *
 * [700] Search in a Binary Search Tree
 *
 * https://leetcode.com/problems/search-in-a-binary-search-tree/description/
 *
 * algorithms
 * Easy (70.71%)
 * Likes:    785
 * Dislikes: 117
 * Total Accepted:    181.1K
 * Total Submissions: 248.3K
 * Testcase Example:  '[4,2,7,1,3]\n2'
 *
 * Given the root node of a binary search tree (BST) and a value. You need to
 * find the node in the BST that the node's value equals the given value.
 * Return the subtree rooted with that node. If such node doesn't exist, you
 * should return NULL.
 *
 * For example,
 *
 *
 * Given the tree:
 * ⁠       4
 * ⁠      / \
 * ⁠     2   7
 * ⁠    / \
 * ⁠   1   3
 *
 * And the value to search: 2
 *
 *
 * You should return this subtree:
 *
 *
 * ⁠     2
 * ⁠    / \
 * ⁠   1   3
 *
 *
 * In the example above, if we want to search the value 5, since there is no
 * node with value 5, we should return NULL.
 *
 * Note that an empty tree is represented by NULL, therefore you would see the
 * expected output (serialized tree format) as [], not null.
 *
 */
package main

import "fmt"

func main() {
	root := &TreeNode{Val: 4}
	node1 := &TreeNode{Val: 2}
	node2 := &TreeNode{Val: 7}
	node3 := &TreeNode{Val: 1}
	node4 := &TreeNode{Val: 3}

	node1.Left = node3
	node1.Right = node4

	root.Left = node1
	root.Right = node2

	node := searchBST(root, 5)

	fmt.Printf("%v\n", node)
}

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

// @lc code=start

func searchBST(root *TreeNode, val int) *TreeNode {
	if root == nil {
		return nil
	}
	nodes := []*TreeNode{root}
	for len(nodes) > 0 {
		node := nodes[0]
		nodes = nodes[1:]
		if node.Val == val {
			return node
		}
		if node.Left != nil {
			nodes = append(nodes, node.Left)
		}
		if node.Right != nil {
			nodes = append(nodes, node.Right)
		}
	}
	return nil
}

// @lc code=end
