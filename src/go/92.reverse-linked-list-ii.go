/*
 * @lc app=leetcode id=92 lang=golang
 *
 * [92] Reverse Linked List II
 *
 * https://leetcode.com/problems/reverse-linked-list-ii/description/
 *
 * algorithms
 * Medium (37.63%)
 * Likes:    2463
 * Dislikes: 145
 * Total Accepted:    277.6K
 * Total Submissions: 715.8K
 * Testcase Example:  '[1,2,3,4,5]\n2\n4'
 *
 * Reverse a linked list from position m to n. Do it in one-pass.
 *
 * Note: 1 ≤ m ≤ n ≤ length of list.
 *
 * Example:
 *
 *
 * Input: 1->2->3->4->5->NULL, m = 2, n = 4
 * Output: 1->4->3->2->5->NULL
 *
 *
 */
package main

import "fmt"

func main() {
	nums := []int{1, 2, 3, 4, 5}
	head := genList(nums, 0)
	m, n := 1, 5
	head = reverseBetween(head, m, n)
	printList(head)
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

type ListNode struct {
	Val  int
	Next *ListNode
}

// @lc code=start
func reverseBetween(head *ListNode, m int, n int) *ListNode {
	if head == nil || head.Next == nil {
		return head
	}
	if m >= n {
		return head
	}
	idx := 1
	var pre, current, next *ListNode
	var nodeM, mPre *ListNode

	current = head

	// 先找到第m个节点
	for current != nil && idx < m {
		pre = current
		current = current.Next
		idx++
	}

	nodeM = current
	mPre = pre

	pre = current
	current = current.Next
	next = current.Next
	pre.Next = nil

	// 开始反转区间链表
	for current != nil && idx < n {
		idx++
		next = current.Next
		current.Next = pre

		pre = current
		current = next
	}

	if mPre != nil {
		mPre.Next = pre
	} else {
		head = pre
	}
	nodeM.Next = current

	return head
}

// @lc code=end
