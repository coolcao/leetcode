/*
 * @lc app=leetcode id=1290 lang=golang
 *
 * [1290] Convert Binary Number in a Linked List to Integer
 *
 * https://leetcode.com/problems/convert-binary-number-in-a-linked-list-to-integer/description/
 *
 * algorithms
 * Easy (79.63%)
 * Likes:    413
 * Dislikes: 35
 * Total Accepted:    66.9K
 * Total Submissions: 83.5K
 * Testcase Example:  '[1,0,1]'
 *
 * Given head which is a reference node to a singly-linked list. The value of
 * each node in the linked list is either 0 or 1. The linked list holds the
 * binary representation of a number.
 *
 * Return the decimal value of the number in the linked list.
 *
 *
 * Example 1:
 *
 *
 * Input: head = [1,0,1]
 * Output: 5
 * Explanation: (101) in base 2 = (5) in base 10
 *
 *
 * Example 2:
 *
 *
 * Input: head = [0]
 * Output: 0
 *
 *
 * Example 3:
 *
 *
 * Input: head = [1]
 * Output: 1
 *
 *
 * Example 4:
 *
 *
 * Input: head = [1,0,0,1,0,0,1,1,1,0,0,0,0,0,0]
 * Output: 18880
 *
 *
 * Example 5:
 *
 *
 * Input: head = [0,0]
 * Output: 0
 *
 *
 *
 * Constraints:
 *
 *
 * The Linked List is not empty.
 * Number of nodes will not exceed 30.
 * Each node's value is either 0 or 1.
 *
 */
package main

import (
	"fmt"
)

func main() {
	node1 := &ListNode{Val: 1, Next: nil}
	node2 := &ListNode{Val: 0, Next: nil}
	node3 := &ListNode{Val: 1, Next: nil}
	node4 := &ListNode{Val: 1, Next: nil}
	node5 := &ListNode{Val: 0, Next: nil}

	node1.Next = node2
	node2.Next = node3
	node3.Next = node4
	node4.Next = node5

	fmt.Printf("%v\n", getDecimalValue(node1))

}

type ListNode struct {
	Val  int
	Next *ListNode
}

// @lc code=start

func getDecimalValue(head *ListNode) int {
	sum := 0
	for head != nil {
		sum = sum*2 + head.Val
		head = head.Next
	}
	return sum
}

// @lc code=end
