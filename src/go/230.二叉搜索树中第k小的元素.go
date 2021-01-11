/*
 * @lc app=leetcode.cn id=230 lang=golang
 *
 * [230] 二叉搜索树中第K小的元素
 *
 * https://leetcode-cn.com/problems/kth-smallest-element-in-a-bst/description/
 *
 * algorithms
 * Medium (72.57%)
 * Likes:    335
 * Dislikes: 0
 * Total Accepted:    84.5K
 * Total Submissions: 116.4K
 * Testcase Example:  '[3,1,4,null,2]\n1'
 *
 * 给定一个二叉搜索树，编写一个函数 kthSmallest 来查找其中第 k 个最小的元素。
 *
 * 说明：
 * 你可以假设 k 总是有效的，1 ≤ k ≤ 二叉搜索树元素个数。
 *
 * 示例 1:
 *
 * 输入: root = [3,1,4,null,2], k = 1
 * ⁠  3
 * ⁠ / \
 * ⁠1   4
 * ⁠ \
 * 2
 * 输出: 1
 *
 * 示例 2:
 *
 * 输入: root = [5,3,6,2,4,null,null,1], k = 3
 * ⁠      5
 * ⁠     / \
 * ⁠    3   6
 * ⁠   / \
 * ⁠  2   4
 * ⁠ /
 * ⁠1
 * 输出: 3
 *
 * 进阶：
 * 如果二叉搜索树经常被修改（插入/删除操作）并且你需要频繁地查找第 k 小的值，你将如何优化 kthSmallest 函数？
 *
 */

package main

import (
	"fmt"
)

func main() {
	nums := []int{5, 3, 6, 2, 4, 0, 0, 1}

	root := createTreeFromArray(nums)
	k := 3
	num := kthSmallest(root, k)
	fmt.Println(num)
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

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

// @lc code=start
func visit(node *TreeNode, nums *[]int) {
	if node == nil {
		return
	}
	visit(node.Left, nums)
	*nums = append(*nums, node.Val)
	visit(node.Right, nums)
}
func kthSmallest(root *TreeNode, k int) int {
	nums := []int{}
	visit(root, &nums)
	return nums[k-1]
}

// @lc code=end
