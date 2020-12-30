/*
* 间隔反转链表
* 给定一个链表，每隔k个节点一组进行反转。
* 例如，给定链表 1->2->3->4->5，k=3，反转后得到 3->2->1->4->5
**/
package main

import "fmt"

func main() {

	nums := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 13, 14, 15}
	head := genList(nums, 0)

	printList(head)
	head = reverseK(head, 6)
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

func reverse(head *ListNode) *ListNode {
	if head == nil || head.Next == nil {
		return head
	}

	next := head.Next
	head.Next = reverse(next.Next)
	next.Next = head

	return next
}

func reverseK(head *ListNode, k int) *ListNode {
	if head == nil || head.Next == nil {
		return head
	}

	var pre, current, next *ListNode

	count := 0
	current = head
	for current != nil {
		count++
		current = current.Next
		if count == k {
			break
		}
	}
	// 如果链表长度小于k则不进行反转
	if count < k {
		return head
	}

	pre, current, next = head, head.Next, head.Next.Next
	pre.Next = nil
	idx := 2
	for current != nil && idx <= k {
		idx++
		next = current.Next
		current.Next = pre

		pre = current
		current = next
	}

	head.Next = reverseK(next, k)

	return pre

}
