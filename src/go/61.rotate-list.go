/*
 * @lc app=leetcode id=61 lang=golang
 *
 * [61] Rotate List
 *
 * https://leetcode.com/problems/rotate-list/description/
 *
 * algorithms
 * Medium (29.13%)
 * Likes:    1295
 * Dislikes: 1014
 * Total Accepted:    278.9K
 * Total Submissions: 931K
 * Testcase Example:  '[1,2,3,4,5]\n2'
 *
 * Given a linked list, rotate the list to the right by k places, where k is
 * non-negative.
 *
 * Example 1:
 *
 *
 * Input: 1->2->3->4->5->NULL, k = 2
 * Output: 4->5->1->2->3->NULL
 * Explanation:
 * rotate 1 steps to the right: 5->1->2->3->4->NULL
 * rotate 2 steps to the right: 4->5->1->2->3->NULL
 *
 *
 * Example 2:
 *
 *
 * Input: 0->1->2->NULL, k = 4
 * Output: 2->0->1->NULL
 * Explanation:
 * rotate 1 steps to the right: 2->0->1->NULL
 * rotate 2 steps to the right: 1->2->0->NULL
 * rotate 3 steps to the right: 0->1->2->NULL
 * rotate 4 steps to the right: 2->0->1->NULL
 *
 */
package main

import "fmt"

func main() {
	nums := []int{1, 2}
	head := genList(nums, 0)
	node := rotateRight(head, 2)
	printList(node)

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
func count(head *ListNode) (int, map[int]*ListNode) {
	if head == nil {
		return 0, nil
	}
	hmap := map[int]*ListNode{}
	c := 0
	for head != nil {
		c++
		hmap[c] = head
		head = head.Next
	}
	return c, hmap
}
func rotateRight(head *ListNode, k int) *ListNode {
	c, hmap := count(head)
	if c == 0 || c == 1 {
		return head
	}
	if k >= c {
		k = k % c
	}
	if k == 0 {
		return head
	}
	tail := hmap[c]
	last := hmap[c-k]
	first := hmap[c-k+1]
	tail.Next = head
	last.Next = nil
	return first
}

// @lc code=end
