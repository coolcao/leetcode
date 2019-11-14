/*
 * @lc app=leetcode id=160 lang=golang
 *
 * [160] Intersection of Two Linked Lists
 */

package main

import "fmt"

// ListNode 链表节点
type ListNode struct {
	Val  int
	Next *ListNode
}

func lengthOfListNode(head *ListNode) int {
	length := 0
	node := head
	for {
		if node != nil {
			length++
			node = node.Next
		} else {
			return length
		}
	}
}

func getIntersectionNode(headA, headB *ListNode) *ListNode {
	if headA == headB {
		return headA
	}
	lengthA := lengthOfListNode(headA)
	lengthB := lengthOfListNode(headB)
	// fmt.Printf("LA:%d\n", lengthA)
	// fmt.Printf("LB:%d\n", lengthB)

	diff := 0
	a, b := headA, headB
	length := 0
	if lengthA > lengthB {
		length = lengthB
		diff = lengthA - lengthB
		for i := 0; i < diff; i++ {
			a = a.Next
		}
	} else {
		length = lengthA
		diff = lengthB - lengthA
		for i := 0; i < diff; i++ {
			b = b.Next
		}
	}
	// fmt.Printf("a: %v\n", a)
	// fmt.Printf("b: %v\n", b)
	for i := 0; i < length; i++ {
		if a == b {
			return a
		}
		a = a.Next
		b = b.Next
	}
	return nil
}

func main() {
	node1 := &ListNode{Val: 1}
	node2 := &ListNode{Val: 2}

	node2.Next = node1

	node := getIntersectionNode(node1, node2)
	fmt.Printf("%v\n", node)

}

// @lc code=end
