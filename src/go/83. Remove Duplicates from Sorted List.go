/*
Given a sorted linked list, delete all duplicates such that each element appear only once.

Example 1:

Input: 1->1->2
Output: 1->2
Example 2:

Input: 1->1->2->3->3
Output: 1->2->3
*/

package main

import "fmt"

// ListNode 链表节点
type ListNode struct {
	Val  int
	Next *ListNode
}

func deleteDuplicates(head *ListNode) *ListNode {
	if head == nil {
		return nil
	}

	current := head
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

	return head
}

func display(head *ListNode) {
	current := head
	for current != nil {
		fmt.Printf("%d", current.Val)
		if current.Next != nil {
			fmt.Printf("->")
		}
		current = current.Next
	}
	fmt.Println("\n")
}

func createNode(num int) *ListNode {
	node := new(ListNode)
	node.Val = num
	node.Next = nil
	return node
}
func createListFromArray(arr []int) *ListNode {
	if len(arr) == 0 {
		return nil
	}
	i := 0
	var pre *ListNode
	var head *ListNode
	for i < len(arr) {
		node := createNode(arr[i])
		if pre != nil {
			pre.Next = node
		}
		if i == 0 {
			head = node
		}
		i++
		pre = node
	}
	return head
}

func main() {
	head := createListFromArray([]int{1, 2, 2, 3, 3, 3, 3, 4, 4, 5})

	display(head)
	deleteDuplicates(head)
	display(head)
}
