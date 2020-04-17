/*
 * @lc app=leetcode id=501 lang=golang
 *
 * [501] Find Mode in Binary Search Tree
 *
 * https://leetcode.com/problems/find-mode-in-binary-search-tree/description/
 *
 * algorithms
 * Easy (41.33%)
 * Likes:    772
 * Dislikes: 296
 * Total Accepted:    77.4K
 * Total Submissions: 187K
 * Testcase Example:  '[1,null,2,2]'
 *
 * Given a binary search tree (BST) with duplicates, find all the mode(s) (the
 * most frequently occurred element) in the given BST.
 *
 * Assume a BST is defined as follows:
 *
 *
 * The left subtree of a node contains only nodes with keys less than or equal
 * to the node's key.
 * The right subtree of a node contains only nodes with keys greater than or
 * equal to the node's key.
 * Both the left and right subtrees must also be binary search trees.
 *
 *
 *
 *
 * For example:
 * Given BST [1,null,2,2],
 *
 *
 * ⁠  1
 * ⁠   \
 * ⁠    2
 * ⁠   /
 * ⁠  2
 *
 *
 *
 *
 * return [2].
 *
 * Note: If a tree has more than one mode, you can return them in any order.
 *
 * Follow up: Could you do that without using any extra space? (Assume that the
 * implicit stack space incurred due to recursion does not count).
 *
 */

// @lc code=start
package main

// type TreeNode struct {
// 	Val   int
// 	Left  *TreeNode
// 	Right *TreeNode
// }

// func createNode(num interface{}) *TreeNode {
// 	if num == nil {
// 		return nil
// 	}
// 	node := new(TreeNode)
// 	node.Val = num.(int)
// 	node.Left, node.Right = nil, nil
// 	return node
// }
// func createTreeFromArray(nums []interface{}) *TreeNode {
// 	if len(nums) == 0 {
// 		return nil
// 	}

// 	nodes := []*TreeNode{}
// 	root := createNode(nums[0])
// 	nodes = append(nodes, root)
// 	current := root
// 	for i := 1; i < len(nums); i++ {
// 		node := createNode(nums[i])
// 		if i%2 == 1 {
// 			current.Left = node
// 		} else {
// 			current.Right = node
// 		}
// 		nodes = append(nodes, node)
// 		current = nodes[i/2]
// 		if current == nil {
// 			current = nodes[i/2+1]
// 		}
// 	}

// 	return root
// }
// func main() {
// 	// nums := []interface{}{6, 3, 9, 1, 4, 7, 9, nil, 1, 4, 4}
// 	nums := []interface{}{5, 3, 8, 2, 4, 6, 8, nil, nil, 4}
// 	root := createTreeFromArray(nums)
// 	res := findMode(root)
// 	fmt.Printf("%v\n", res)
// }

var max int       //相同节点数计数最大数
var count int     // 相同节点数计数
var result []int  // 存放结果集
var pre *TreeNode // 保存前一个节点

func visit(node *TreeNode) {
	if node == nil {
		return
	}
	if node.Left != nil {
		visit(node.Left)
	}

	// 节点值有变化
	if node.Val > pre.Val {
		count = 1
	} else {
		count++
	}
	pre = node

	if count > max {
		max = count
		result = []int{node.Val}
	} else if count == max {
		result = append(result, node.Val)
	}

	if node.Right != nil {
		visit(node.Right)
	}
}

func findMode(root *TreeNode) []int {
	max = 0
	count = 0
	result = []int{}
	if root == nil {
		return result
	}
	pre = root
	visit(root)
	return result
}

// @lc code=end
