/*
 * @lc app=leetcode id=19 lang=golang
 *
 * [19] Remove Nth Node From End of List
 *
 * https://leetcode.com/problems/remove-nth-node-from-end-of-list/description/
 *
 * algorithms
 * Medium (34.87%)
 * Likes:    3475
 * Dislikes: 234
 * Total Accepted:    643.9K
 * Total Submissions: 1.8M
 * Testcase Example:  '[1,2,3,4,5]\n2'
 *
 * Given a linked list, remove the n-th node from the end of list and return
 * its head.
 *
 * Example:
 *
 *
 * Given linked list: 1->2->3->4->5, and n = 2.
 *
 * After removing the second node from the end, the linked list becomes
 * 1->2->3->5.
 *
 *
 * Note:
 *
 * Given n will always be valid.
 *
 * Follow up:
 *
 * Could you do this in one pass?
 *
 */

package main

import "fmt"

func main() {
	head := &ListNode{Val: 1}
	// n2 := &ListNode{Val: 2}
	// n3 := &ListNode{Val: 3}
	// n4 := &ListNode{Val: 4}
	// n5 := &ListNode{Val: 5}
	// n6 := &ListNode{Val: 6}

	// head.Next = n2
	// n2.Next = n3
	// n3.Next = n4
	// n4.Next = n5
	// n5.Next = n6
	printList(head)
	printList(removeNthFromEnd(head, 1))
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

// @lc code=start

func removeNthFromEnd(head *ListNode, n int) *ListNode {
	hmap := map[int]*ListNode{}

	current := head
	count := 1

	for current != nil {
		hmap[count] = current
		count++
		current = current.Next
	}

	delNode := hmap[count-n]
	preNode := hmap[count-n-1]

	if preNode == nil {
		head = delNode.Next
	} else {
		preNode.Next = delNode.Next
	}
	return head
}

// @lc code=end
