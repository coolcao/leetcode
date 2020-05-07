/*
 * @lc app=leetcode id=239 lang=golang
 *
 * [239] Sliding Window Maximum
 *
 * https://leetcode.com/problems/sliding-window-maximum/description/
 *
 * algorithms
 * Hard (41.31%)
 * Likes:    2986
 * Dislikes: 164
 * Total Accepted:    247.2K
 * Total Submissions: 589.7K
 * Testcase Example:  '[1,3,-1,-3,5,3,6,7]\n3'
 *
 * Given an array nums, there is a sliding window of size k which is moving
 * from the very left of the array to the very right. You can only see the k
 * numbers in the window. Each time the sliding window moves right by one
 * position. Return the max sliding window.
 *
 * Follow up:
 * Could you solve it in linear time?
 *
 * Example:
 *
 *
 * Input: nums = [1,3,-1,-3,5,3,6,7], and k = 3
 * Output: [3,3,5,5,6,7]
 * Explanation:
 *
 * Window position                Max
 * ---------------               -----
 * [1  3  -1] -3  5  3  6  7      3
 * ⁠1 [3  -1  -3] 5  3  6  7       3
 * ⁠1  3 [-1  -3  5] 3  6  7       5
 * ⁠1  3  -1 [-3  5  3] 6  7       5
 * ⁠1  3  -1  -3 [5  3  6] 7       6
 * ⁠1  3  -1  -3  5 [3  6  7]      7
 *
 *
 *
 * Constraints:
 *
 *
 * 1 <= nums.length <= 10^5
 * -10^4 <= nums[i] <= 10^4
 * 1 <= k <= nums.length
 *
 *
 */
package main

import (
	"fmt"
)

func main() {
	nums := []int{1, 3, -1, -3, 5, 3, 6, 7, 1, 4, 9, 4, 7, 2, 6, 5, 3, -1, -4, 6}
	k := 6
	res := maxSlidingWindow(nums, k)
	fmt.Printf("%v\n", res)
	// dq := new(DoubleQueue)
	// fmt.Printf("size: %d\n", dq.Size)
	// dq.LeftPush(1)
	// dq.LeftPush(2)
	// Print(dq.Left)
	// fmt.Printf("size: %d\n", dq.Size)
	// dq.LeftPop()
	// Print(dq.Left)
	// PrintI(dq.Right)
	// fmt.Printf("size: %d\n", dq.Size)
	// dq.LeftPop()
	// Print(dq.Left)
	// PrintI(dq.Right)
	// fmt.Printf("size: %d\n", dq.Size)
}

// @lc code=start
// QueueNode 队列节点
type QueueNode struct {
	Val  int
	Next *QueueNode
	Pre  *QueueNode
}

// DoubleQueue 双端队列
type DoubleQueue struct {
	Left  *QueueNode
	Right *QueueNode
	Size  int
}

func Print(qn *QueueNode) {
	if qn == nil {
		return
	}
	fmt.Printf("l:%d\t", qn.Val)
	Print(qn.Next)
}
func PrintI(qn *QueueNode) {
	if qn == nil {
		return
	}
	fmt.Printf("r:%d\t", qn.Val)
	PrintI(qn.Pre)
}
func (dq *DoubleQueue) LeftPeek() int {
	return dq.Left.Val
}

func (dq *DoubleQueue) LeftPush(num int) {
	node := new(QueueNode)
	node.Val = num
	if dq.Left == nil && dq.Right == nil {
		dq.Left = node
		dq.Right = node
	} else {
		current := dq.Left
		current.Pre = node
		node.Next = current
		dq.Left = node
	}

	dq.Size++
}

func (dq *DoubleQueue) LeftPop() int {
	current := dq.Left
	dq.Left = current.Next
	if dq.Left != nil {
		dq.Left.Pre = nil
	}
	dq.Size--
	if dq.Size == 0 {
		dq.Right = nil
	}
	return current.Val
}
func (dq *DoubleQueue) RightPeek() int {
	return dq.Right.Val
}
func (dq *DoubleQueue) RightPush(num int) {
	node := new(QueueNode)
	node.Val = num
	if dq.Left == nil && dq.Right == nil {
		dq.Left = node
		dq.Right = node
	} else {
		current := dq.Right
		current.Next = node
		node.Pre = current
		dq.Right = node
	}

	dq.Size++
}
func (dq *DoubleQueue) RightPop() int {
	current := dq.Right
	dq.Right = current.Pre
	if dq.Right != nil {
		dq.Right.Next = nil
	}
	dq.Size--
	if dq.Size == 0 {
		dq.Left = nil
	}
	return current.Val
}

func maxSlidingWindow(nums []int, k int) []int {
	length := len(nums)
	res := []int{}

	// 初始化一个双端队列，用于存储窗口内的最大值
	dq := new(DoubleQueue)

	left, right := 0, 0
	for right < length {
		rightNum := nums[right]
		if dq.Size == 0 {
			dq.RightPush(rightNum)
		} else {
			for dq.Size > 0 && rightNum > dq.RightPeek() {
				dq.RightPop()
			}
			dq.RightPush(rightNum)
		}
		right++
		if right-left == k {
			res = append(res, dq.LeftPeek())
			if dq.Size > 0 && nums[left] == dq.LeftPeek() {
				dq.LeftPop()
			}
			left++
		}
	}

	return res
}

// @lc code=end
