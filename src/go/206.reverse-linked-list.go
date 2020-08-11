/*
 * @lc app=leetcode id=206 lang=golang
 *
 * [206] Reverse Linked List
 *
 * https://leetcode.com/problems/reverse-linked-list/description/
 *
 * algorithms
 * Easy (60.39%)
 * Likes:    4776
 * Dislikes: 91
 * Total Accepted:    1M
 * Total Submissions: 1.7M
 * Testcase Example:  '[1,2,3,4,5]'
 *
 * Reverse a singly linked list.
 *
 * Example:
 *
 *
 * Input: 1->2->3->4->5->NULL
 * Output: 5->4->3->2->1->NULL
 *
 *
 * Follow up:
 *
 * A linked list can be reversed either iteratively or recursively. Could you
 * implement both?
 *
 */
package main

import "fmt"

func main() {
	nums := []int{1, 2, 3, 4, 5}
	head := genList(nums, 0)
	head = reverseList(head)
	printList(head)
}

type ListNode struct {
	Val  int
	Next *ListNode
}

func printList(head *ListNode) {
	result := []int{}
	for head != nil {
		result = append(result, head.Val)
		head = head.Next
	}
	fmt.Printf("%v\n", result)
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

// 递归
// func reverseList(head *ListNode) *ListNode {
// 	if head == nil || head.Next == nil {
// 		return head
// 	}
// 	next := head.Next
// 	reversed := reverseList(next)
// 	next.Next = head
// 	head.Next = nil
// 	return reversed
// }

// 迭代
func reverseList(head *ListNode) *ListNode {
	if head == nil || head.Next == nil {
		return head
	}
	pre, current, next := head, head.Next, head.Next.Next
	pre.Next = nil
	for current != nil {
		next = current.Next
		current.Next = pre

		pre = current
		current = next
	}

	return pre
}

// @lc code=end
