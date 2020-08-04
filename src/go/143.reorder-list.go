/*
 * @lc app=leetcode id=143 lang=golang
 *
 * [143] Reorder List
 *
 * https://leetcode.com/problems/reorder-list/description/
 *
 * algorithms
 * Medium (35.03%)
 * Likes:    1986
 * Dislikes: 118
 * Total Accepted:    238.2K
 * Total Submissions: 644.7K
 * Testcase Example:  '[1,2,3,4]'
 *
 * Given a singly linked list L: L0→L1→…→Ln-1→Ln,
 * reorder it to: L0→Ln→L1→Ln-1→L2→Ln-2→…
 *
 * You may not modify the values in the list's nodes, only nodes itself may be
 * changed.
 *
 * Example 1:
 *
 *
 * Given 1->2->3->4, reorder it to 1->4->2->3.
 *
 * Example 2:
 *
 *
 * Given 1->2->3->4->5, reorder it to 1->5->2->4->3.
 *
 *
 */
package main

import "fmt"

func main() {
	nums := []int{1, 2, 3, 4, 5, 6}
	head := genList(nums, 0)
	printList(head)
	reorderList(head)
	printList(head)
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
// 先将每个节点缓存到map，然后依次拿出前后节点进行拼装
func reorderList(head *ListNode) {
	if head == nil {
		return
	}
	hmap := map[int]*ListNode{}
	c := 0
	current := head
	for current != nil {
		c++
		hmap[c] = current
		current = current.Next
	}
	i, j := 1, c
	current = head
	for i <= j {
		current.Next = hmap[i]
		current = current.Next
		i++
		current.Next = hmap[j]
		current = current.Next
		j--
	}
	current.Next = nil
}

// @lc code=end
