/*
 * @lc app=leetcode id=25 lang=golang
 *
 * [25] Reverse Nodes in k-Group
 *
 * https://leetcode.com/problems/reverse-nodes-in-k-group/description/
 *
 * algorithms
 * Hard (40.16%)
 * Likes:    2341
 * Dislikes: 359
 * Total Accepted:    275K
 * Total Submissions: 655.5K
 * Testcase Example:  '[1,2,3,4,5]\n2'
 *
 * Given a linked list, reverse the nodes of a linked list k at a time and
 * return its modified list.
 *
 * k is a positive integer and is less than or equal to the length of the
 * linked list. If the number of nodes is not a multiple of k then left-out
 * nodes in the end should remain as it is.
 *
 *
 *
 *
 * Example:
 *
 * Given this linked list: 1->2->3->4->5
 *
 * For k = 2, you should return: 2->1->4->3->5
 *
 * For k = 3, you should return: 3->2->1->4->5
 *
 * Note:
 *
 *
 * Only constant extra memory is allowed.
 * You may not alter the values in the list's nodes, only nodes itself may be
 * changed.
 *
 *
 */

package main

import "fmt"

func main() {
	nums := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12}
	head := genList(nums, 0)

	printList(head)
	head = reverseKGroup(head, 3)
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
// type NodeGroup struct {
// 	Count int       `json:"count"`
// 	Head  *ListNode `json:"head"`
// }

// func split(head *ListNode, k int) []*NodeGroup {
// 	if head == nil {
// 		return nil
// 	}
// 	result := []*NodeGroup{}
// 	c := 0
// 	current := head
// 	for c < k {
// 		c++
// 		next := current.Next
// 		if next == nil {
// 			result = append(result, &NodeGroup{
// 				Count: c,
// 				Head:  head,
// 			})
// 			current = current.Next
// 			return result
// 		}
// 		if c == k {
// 			current.Next = nil
// 			result = append(result, &NodeGroup{
// 				Count: k,
// 				Head:  head,
// 			})
// 			result = append(result, split(next, k)...)
// 		} else {
// 			current = current.Next
// 		}
// 	}
// 	return result
// }
// func reverse(head *ListNode) *ListNode {
// 	if head == nil || head.Next == nil {
// 		return head
// 	}
// 	next := head.Next
// 	node := reverse(next)
// 	next.Next = head
// 	head.Next = nil
// 	return node
// }
// func reverseKGroup(head *ListNode, k int) *ListNode {
// 	if head == nil {
// 		return nil
// 	}
// 	nodeGroups := split(head, k)
// 	result := &ListNode{}
// 	current := result
// 	for _, ng := range nodeGroups {
// 		h := ng.Head
// 		if ng.Count == k {
// 			current.Next = reverse(ng.Head)
// 			current = h
// 		} else {
// 			current.Next = ng.Head
// 		}
// 	}

// 	return result.Next
// }

func reverse(head *ListNode) *ListNode {
	if head == nil || head.Next == nil {
		return head
	}

	next := head.Next
	node := reverse(next)
	next.Next = head
	head.Next = nil

	return node
}
func reverseKGroup(head *ListNode, k int) *ListNode {
	if head == nil {
		return nil
	}

	count := 0
	// 先计算链表长度
	current := head
	for current != nil {
		count++
		current = current.Next
	}

	start, end := head, head
	next := end.Next
	preHead := head
	used := 0

	for end != nil {
		next = end.Next
		used++
		if used%k == 0 {
			fmt.Printf("start: %v, end: %v\n", start.Val, end.Val)
			end.Next = nil
			node := reverse(start)
			if used == k {
				head = node
			}
			preHead.Next = node
			preHead = start
			start.Next = next
			start = next
		}
		end = next
	}

	return head
}

// @lc code=end
