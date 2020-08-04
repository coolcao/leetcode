/*
 * @lc app=leetcode id=86 lang=golang
 *
 * [86] Partition List
 *
 * https://leetcode.com/problems/partition-list/description/
 *
 * algorithms
 * Medium (40.13%)
 * Likes:    1337
 * Dislikes: 303
 * Total Accepted:    218.7K
 * Total Submissions: 528.3K
 * Testcase Example:  '[1,4,3,2,5,2]\n3'
 *
 * Given a linked list and a value x, partition it such that all nodes less
 * than x come before nodes greater than or equal to x.
 *
 * You should preserve the original relative order of the nodes in each of the
 * two partitions.
 *
 * Example:
 *
 *
 * Input: head = 1->4->3->2->5->2, x = 3
 * Output: 1->2->2->4->3->5
 *
 *
 */
package main

import "fmt"

func main() {
	nums := []int{1, 4, 3, 2, 5, 2}
	x := 3
	head := genList(nums, 0)
	head = partition(head, x)
	printList(head)

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

type ListNode struct {
	Val  int
	Next *ListNode
}

// @lc code=start
func partition(head *ListNode, x int) *ListNode {
	less, greater := &ListNode{}, &ListNode{}
	current := head
	lc, gc := less, greater
	for current != nil {
		if current.Val < x {
			lc.Next = current
			lc = lc.Next
			current = current.Next
			lc.Next = nil
		} else {
			gc.Next = current
			gc = gc.Next
			current = current.Next
			gc.Next = nil
		}
	}

	lc.Next = greater.Next

	return less.Next
}

// @lc code=end
