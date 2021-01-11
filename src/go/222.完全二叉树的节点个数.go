/*
 * @lc app=leetcode.cn id=222 lang=golang
 *
 * [222] 完全二叉树的节点个数
 *
 * https://leetcode-cn.com/problems/count-complete-tree-nodes/description/
 *
 * algorithms
 * Medium (76.55%)
 * Likes:    414
 * Dislikes: 0
 * Total Accepted:    74.3K
 * Total Submissions: 97K
 * Testcase Example:  '[1,2,3,4,5,6]'
 *
 * 给出一个完全二叉树，求出该树的节点个数。
 *
 * 说明：
 * 完全二叉树的定义如下：在完全二叉树中，除了最底层节点可能没填满外，其余每层节点数都达到最大值，并且最下面一层的节点都集中在该层最左边的若干位置。若最底层为第
 * h 层，则该层包含 1~ 2^h 个节点。
 *
 * 示例:

 * 输入:
 * ⁠   1
 * ⁠  / \
 * ⁠ 2   3
 * ⁠/ \  /
 * 4  5 6
 *
 * 输出: 6
 *
 */
package main

import (
	"fmt"
	"sort"
)

func main() {
	root := createTreeFromArray([]int{1, 2, 3, 4, 5, 6})
	sum := countNodes(root)
	fmt.Println(sum)
}

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
		if current == nil {
			current = nodes[i/2+1]
		}
	}

	return root
}

// @lc code=start
// 递归方式
// func countHelper(node *TreeNode) int {
// 	if node == nil {
// 		return 0
// 	}
// 	return 1 + countHelper(node.Left) + countHelper(node.Right)
// }
// func countNodes(root *TreeNode) int {
// 	if root == nil {
// 		return 0
// 	}
// 	return countHelper(root)
// }

func countNodes(root *TreeNode) int {
	if root == nil {
		return 0
	}
	level := 0
	for node := root; node.Left != nil; node = node.Left {
		level++
	}
	return sort.Search(1<<(level+1), func(k int) bool {
		if k <= 1<<level {
			return false
		}
		bits := 1 << (level - 1)
		node := root
		for node != nil && bits > 0 {
			if bits&k == 0 {
				node = node.Left
			} else {
				node = node.Right
			}
			bits >>= 1
		}
		return node == nil
	}) - 1
}

// @lc code=end
