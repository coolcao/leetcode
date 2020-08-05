/*
 * @lc app=leetcode id=328 lang=golang
 *
 * [328] Odd Even Linked List
 *
 * https://leetcode.com/problems/odd-even-linked-list/description/
 *
 * algorithms
 * Medium (52.65%)
 * Likes:    1999
 * Dislikes: 303
 * Total Accepted:    287.3K
 * Total Submissions: 516.4K
 * Testcase Example:  '[1,2,3,4,5]'
 *
 * Given a singly linked list, group all odd nodes together followed by the
 * even nodes. Please note here we are talking about the node number and not
 * the value in the nodes.
 *
 * You should try to do it in place. The program should run in O(1) space
 * complexity and O(nodes) time complexity.
 *
 * Example 1:
 *
 *
 * Input: 1->2->3->4->5->NULL
 * Output: 1->3->5->2->4->NULL
 *
 *
 * Example 2:
 *
 *
 * Input: 2->1->3->5->6->4->7->NULL
 * Output: 2->3->6->7->1->5->4->NULL
 *
 *
 *
 * Constraints:
 *
 *
 * The relative order inside both the even and odd groups should remain as it
 * was in the input.
 * The first node is considered odd, the second node even and so on ...
 * The length of the linked list is between [0, 10^4].
 *
 *
 */
package main

import "fmt"

func main() {
	nums := []int{1, 2, 3, 4, 5}
	head := genList(nums, 0)
	printList(head)
	head = oddEvenList(head)
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

// func oddEvenList(head *ListNode) *ListNode {
// 	// 空链表，或者只有一个或两个节点的链表不做任何操作
// 	if head == nil || head.Next == nil || head.Next.Next == nil {
// 		return head
// 	}

// 	c := 2

// 	odd, even := head, head.Next

// 	current := even.Next

// 	for current != nil {
// 		c++
// 		if c%2 == 1 {
// 			// 奇数，调到前面
// 			oddNext := odd.Next
// 			odd.Next = current

// 			currentNext := current.Next
// 			current.Next = oddNext
// 			current = currentNext

// 			odd = odd.Next
// 			even.Next = nil
// 		} else {
// 			// 偶数，拼到后面
// 			currentNext := current.Next

// 			even.Next = current
// 			even = even.Next
// 			even.Next = nil

// 			current = currentNext
// 		}
// 	}

// 	return head
// }
func oddEvenList(head *ListNode) *ListNode {
	if head == nil {
		return head
	}
	odd, even := head, head.Next
	evenHead := even

	// 取出偶数位置的节点，组成一个新链表
	for even != nil && even.Next != nil {
		odd.Next = even.Next
		odd = odd.Next

		even.Next = odd.Next
		even = even.Next
	}
	odd.Next = evenHead
	return head
}

// @lc code=end
