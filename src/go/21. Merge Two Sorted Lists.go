/*
Merge two sorted linked lists and return it as a new list. The new list should be made by splicing together the nodes of the first two lists.

Example:

Input: 1->2->4, 1->3->4
Output: 1->1->2->3->4->4
*/

package main

import "fmt"

type ListNode struct {
	Val  int
	Next *ListNode
}

func mergeTwoLists(l1 *ListNode, l2 *ListNode) *ListNode {
	var result, current *ListNode
	p1, p2 := l1, l2
	for p1 != nil && p2 != nil {
		node := new(ListNode)
		if p1.Val > p2.Val {
			node.Val = p2.Val
			p2 = p2.Next
		} else {
			node.Val = p1.Val
			p1 = p1.Next
		}
		node.Next = nil
		if current == nil {
			current = node
			result = current
		} else {
			current.Next = node
		}
		current = node
	}
	for p1 != nil {
		node := new(ListNode)
		node.Val = p1.Val
		node.Next = nil
		if current == nil {
			current = node
			result = current
		} else {
			current.Next = node
		}
		p1 = p1.Next
		current = current.Next
	}
	for p2 != nil {
		node := new(ListNode)
		node.Val = p2.Val
		node.Next = nil
		if current == nil {
			current = node
			result = current
		} else {
			current.Next = node
		}
		p2 = p2.Next
		current = current.Next
	}

	return result
}

func main() {
	n1 := &ListNode{
		Val: 1,
		Next: &ListNode{
			Val: 2,
			Next: &ListNode{
				Val:  4,
				Next: nil,
			},
		},
	}
	n2 := &ListNode{
		Val: 1,
		Next: &ListNode{
			Val: 3,
			Next: &ListNode{
				Val:  4,
				Next: nil,
			},
		},
	}
	result := mergeTwoLists(n1, n2)
	fmt.Printf("%v\n", result)
}
