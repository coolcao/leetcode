/*
 * @lc app=leetcode id=725 lang=golang
 *
 * [725] Split Linked List in Parts
 *
 * https://leetcode.com/problems/split-linked-list-in-parts/description/
 *
 * algorithms
 * Medium (51.13%)
 * Likes:    671
 * Dislikes: 133
 * Total Accepted:    49.2K
 * Total Submissions: 94.4K
 * Testcase Example:  '[1,2,3,4]\n5'
 *
 * Given a (singly) linked list with head node root, write a function to split
 * the linked list into k consecutive linked list "parts".
 *
 * The length of each part should be as equal as possible: no two parts should
 * have a size differing by more than 1.  This may lead to some parts being
 * null.
 *
 * The parts should be in order of occurrence in the input list, and parts
 * occurring earlier should always have a size greater than or equal parts
 * occurring later.
 *
 * Return a List of ListNode's representing the linked list parts that are
 * formed.
 *
 *
 * Examples
 * 1->2->3->4, k = 5 // 5 equal parts
 * [ [1],
 * [2],
 * [3],
 * [4],
 * null ]
 *
 * Example 1:
 *
 * Input:
 * root = [1, 2, 3], k = 5
 * Output: [[1],[2],[3],[],[]]
 * Explanation:
 * The input and each element of the output are ListNodes, not arrays.
 * For example, the input root has root.val = 1, root.next.val = 2,
 * \root.next.next.val = 3, and root.next.next.next = null.
 * The first element output[0] has output[0].val = 1, output[0].next = null.
 * The last element output[4] is null, but it's string representation as a
 * ListNode is [].
 *
 *
 *
 * Example 2:
 *
 * Input:
 * root = [1, 2, 3, 4, 5, 6, 7, 8, 9, 10], k = 3
 * Output: [[1, 2, 3, 4], [5, 6, 7], [8, 9, 10]]
 * Explanation:
 * The input has been split into consecutive parts with size difference at most
 * 1, and earlier parts are a larger size than the later parts.
 *
 *
 *
 * Note:
 * The length of root will be in the range [0, 1000].
 * Each value of a node in the input will be an integer in the range [0, 999].
 * k will be an integer in the range [1, 50].
 *
 */

package main

import "fmt"

func main() {
	nums := []int{4, 5, 2, 3, 98, 7, 1, 3, 2}
	result := splitListToParts(genList(nums, 0), 3)
	for _, node := range result {
		printList(node)
	}
}

func printList(head *ListNode) {
	nums := []int{}
	for head != nil {
		nums = append(nums, head.Val)
		head = head.Next
	}
	fmt.Printf("%v\n", nums)
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
func split(head *ListNode, count int, k int) []*ListNode {
	result := []*ListNode{}
	if k == 0 {
		return result
	}
	if head == nil || count == 0 {
		result = append(result, nil)
		result = append(result, split(head, count, k-1)...)
		return result
	}
	limit := 1
	if count%k == 0 {
		limit = count / k
	} else {
		limit = count/k + 1
	}
	node := head
	for i := 1; i < limit; i++ {
		node = node.Next
	}
	next := node.Next
	node.Next = nil
	result = append(result, head)
	result = append(result, split(next, count-limit, k-1)...)
	return result
}
func splitListToParts(root *ListNode, k int) []*ListNode {
	count, head := 0, root
	for head != nil {
		count++
		head = head.Next
	}

	return split(root, count, k)
}

// @lc code=end
