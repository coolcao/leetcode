/*
 * @lc app=leetcode id=142 lang=golang
 *
 * [142] Linked List Cycle II
 *
 * https://leetcode.com/problems/linked-list-cycle-ii/description/
 *
 * algorithms
 * Medium (35.76%)
 * Likes:    2878
 * Dislikes: 228
 * Total Accepted:    336.6K
 * Total Submissions: 902K
 * Testcase Example:  '[3,2,0,-4]\n1'
 *
 * Given a linked list, return the node where the cycle begins. If there is no
 * cycle, return null.
 *
 * To represent a cycle in the given linked list, we use an integer pos which
 * represents the position (0-indexed) in the linked list where tail connects
 * to. If pos is -1, then there is no cycle in the linked list.
 *
 * Note: Do not modify the linked list.
 *
 *
 *
 * Example 1:
 *
 *
 * Input: head = [3,2,0,-4], pos = 1
 * Output: tail connects to node index 1
 * Explanation: There is a cycle in the linked list, where tail connects to the
 * second node.
 *
 *
 *
 *
 * Example 2:
 *
 *
 * Input: head = [1,2], pos = 0
 * Output: tail connects to node index 0
 * Explanation: There is a cycle in the linked list, where tail connects to the
 * first node.
 *
 *
 *
 *
 * Example 3:
 *
 *
 * Input: head = [1], pos = -1
 * Output: no cycle
 * Explanation: There is no cycle in the linked list.
 *
 *
 *
 *
 *
 *
 * Follow-up:
 * Can you solve it without using extra space?
 *
 */
package main

import (
	"fmt"
)

func main() {
	n1 := &ListNode{Val: 3}
	n2 := &ListNode{Val: 2}
	n3 := &ListNode{Val: 0}
	n4 := &ListNode{Val: -4}
	// n5 := &ListNode{Val: 19}
	// n6 := &ListNode{Val: 6}
	// n7 := &ListNode{Val: -9}
	// n8 := &ListNode{Val: -5}
	// n9 := &ListNode{Val: -2}
	// n10 := &ListNode{Val: -5}
	n1.Next = n2
	n2.Next = n3
	n3.Next = n4
	// n4.Next = n5
	// n5.Next = n6
	// n6.Next = n7
	// n7.Next = n8
	// n8.Next = n9
	// n9.Next = n10
	n4.Next = n2
	n := detectCycle(n1)
	fmt.Printf("n: %v\n", n.Val)

}

type ListNode struct {
	Next *ListNode
	Val  int
}

// @lc code=start
// 使用哈希表
// func detectCycle(head *ListNode) *ListNode {
// 	if head == nil || head.Next == nil {
// 		return nil
// 	}
// 	hmap := map[*ListNode]bool{}
// 	current := head
// 	for current != nil {
// 		if hmap[current] {
// 			return current
// 		} else {
// 			hmap[current] = true
// 		}
// 		current = current.Next
// 	}
// 	return nil
// }

// 快慢双指针解决
func detectCycle(head *ListNode) *ListNode {
	if head == nil || head.Next == nil {
		return nil
	}
	slow, fast := head, head
	for fast != nil && fast.Next != nil {
		slow = slow.Next
		fast = fast.Next.Next
		// 第一次相遇，说明有环
		if slow == fast {
			break
		}
	}

	if fast != nil && fast.Next != nil {
		slow = head
		for slow != fast {
			slow = slow.Next
			fast = fast.Next
		}
		return slow
	}

	return nil
}

// @lc code=end
