/*
 * @lc app=leetcode id=21 lang=golang
 *
 * [21] Merge Two Sorted Lists
 *
 * https://leetcode.com/problems/merge-two-sorted-lists/description/
 *
 * algorithms
 * Easy (52.84%)
 * Likes:    4463
 * Dislikes: 606
 * Total Accepted:    1M
 * Total Submissions: 1.9M
 * Testcase Example:  '[1,2,4]\n[1,3,4]'
 *
 * Merge two sorted linked lists and return it as a new sorted list. The new
 * list should be made by splicing together the nodes of the first two lists.
 *
 * Example:
 *
 *
 * Input: 1->2->4, 1->3->4
 * Output: 1->1->2->3->4->4
 *
 *
 */
package main

import "fmt"

func main() {

	l1 := &ListNode{Val: 1}
	n11 := &ListNode{Val: 2}
	n12 := &ListNode{Val: 3}
	n13 := &ListNode{Val: 4}
	l1.Next = n11
	n11.Next = n12
	n12.Next = n13

	l2 := &ListNode{Val: 2}
	n21 := &ListNode{Val: 3}
	n22 := &ListNode{Val: 3}
	n23 := &ListNode{Val: 5}
	l2.Next = n21
	n21.Next = n22
	n22.Next = n23

	printList(l1)
	printList(l2)
	head := mergeTwoLists(l1, l2)
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

// @lc code=start
// func mergeTwoLists(l1 *ListNode, l2 *ListNode) *ListNode {
// 	var head *ListNode
// 	var current *ListNode

// 	start := &ListNode{}
// 	current = start

// 	n1, n2 := l1, l2

// 	for n1 != nil && n2 != nil {
// 		if n1.Val < n2.Val {
// 			current.Next = n1
// 			n1 = n1.Next
// 		} else {
// 			current.Next = n2
// 			n2 = n2.Next
// 		}
// 		current = current.Next
// 	}

// 	for n1 != nil {
// 		current.Next = n1
// 		current = current.Next
// 		n1 = n1.Next
// 	}

// 	for n2 != nil {
// 		current.Next = n2
// 		current = current.Next
// 		n2 = n2.Next
// 	}

// 	head = start.Next

// 	return head
// }

func mergeTwoLists(l1 *ListNode, l2 *ListNode) *ListNode {
	if l1 == nil {
		return l2
	}
	if l2 == nil {
		return l1
	}

	var head *ListNode

	if l1.Val < l2.Val {
		head = l1
		l1 = l1.Next
	} else {
		head = l2
		l2 = l2.Next
	}

	head.Next = mergeTwoLists(l1, l2)

	return head

}

// @lc code=end
