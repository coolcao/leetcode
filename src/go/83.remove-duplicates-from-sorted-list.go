/*
 * @lc app=leetcode id=83 lang=golang
 *
 * [83] Remove Duplicates from Sorted List
 *
 * https://leetcode.com/problems/remove-duplicates-from-sorted-list/description/
 *
 * algorithms
 * Easy (44.57%)
 * Likes:    1611
 * Dislikes: 114
 * Total Accepted:    472.7K
 * Total Submissions: 1M
 * Testcase Example:  '[1,1,2]'
 *
 * Given a sorted linked list, delete all duplicates such that each element
 * appear only once.
 *
 * Example 1:
 *
 *
 * Input: 1->1->2
 * Output: 1->2
 *
 *
 * Example 2:
 *
 *
 * Input: 1->1->2->3->3
 * Output: 1->2->3
 *
 *
 */
package main

import "fmt"

func main() {
	nums := []int{1, 1, 1, 1}
	head := genList(nums, 0)

	head = deleteDuplicates(head)
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

func deleteDuplicates(head *ListNode) *ListNode {
	if head == nil {
		return nil
	}

	h, current := head, head
	diffVal := head.Val

	for current.Next != nil {
		if current.Next.Val != diffVal {
			current = current.Next
			diffVal = current.Val
			continue
		}

		next := current.Next.Next
		current.Next = next
	}

	return h
}

// @lc code=end
