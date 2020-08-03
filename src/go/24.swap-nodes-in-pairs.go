/*
 * @lc app=leetcode id=24 lang=golang
 *
 * [24] Swap Nodes in Pairs
 *
 * https://leetcode.com/problems/swap-nodes-in-pairs/description/
 *
 * algorithms
 * Medium (48.83%)
 * Likes:    2378
 * Dislikes: 172
 * Total Accepted:    479K
 * Total Submissions: 952.7K
 * Testcase Example:  '[1,2,3,4]'
 *
 * Given a linked list, swap every two adjacent nodes and return its head.
 *
 * You may not modify the values in the list's nodes, only nodes itself may be
 * changed.
 *
 *
 *
 * Example:
 *
 *
 * Given 1->2->3->4, you should return the list as 2->1->4->3.
 *
 *
 */

package main

import "fmt"

func main() {
	nums := []int{1, 2, 3, 4, 5}
	head := genList(nums, 0)
	printList(head)
	printList(swapPairs(head))
}

type ListNode struct {
	Val  int
	Next *ListNode
}

func printList(head *ListNode) {
	for head != nil {
		fmt.Printf("%v\t", head.Val)
		head = head.Next
	}
	fmt.Println("")
}

func genList(nums []int, idx int) *ListNode {
	if len(nums) <= idx {
		return nil
	}

	head := &ListNode{
		Val:  nums[idx],
		Next: genList(nums, idx+1),
	}

	return head
}

// @lc code=start

func swapPairs(head *ListNode) *ListNode {
	if head == nil {
		return nil
	}
	next := head.Next
	if next == nil {
		return head
	}
	// 取节点对的下一个节点
	nextNext := next.Next
	// 反转
	next.Next = head
	head.Next = swapPairs(nextNext)
	return next
}

// @lc code=end
