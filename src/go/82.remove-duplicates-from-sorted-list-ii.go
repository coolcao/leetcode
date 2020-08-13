/*
 * @lc app=leetcode id=82 lang=golang
 *
 * [82] Remove Duplicates from Sorted List II
 *
 * https://leetcode.com/problems/remove-duplicates-from-sorted-list-ii/description/
 *
 * algorithms
 * Medium (35.66%)
 * Likes:    1760
 * Dislikes: 109
 * Total Accepted:    254.4K
 * Total Submissions: 693K
 * Testcase Example:  '[1,2,3,3,4,4,5]'
 *
 * Given a sorted linked list, delete all nodes that have duplicate numbers,
 * leaving only distinct numbers from the original list.
 *
 * Return the linked list sorted as well.
 *
 * Example 1:
 *
 *
 * Input: 1->2->3->3->4->4->5
 * Output: 1->2->5
 *
 *
 * Example 2:
 *
 *
 * Input: 1->1->1->2->3
 * Output: 2->3
 *
 *
 */

package main

import "fmt"

func main() {
	nums := []int{}
	head := genList(nums, 0)
	printList(head)
	head = deleteDuplicates(head)
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

func deleteDuplicates(head *ListNode) *ListNode {
	if head == nil || head.Next == nil {
		return head
	}
	// 先统计每个节点值出现的次数
	hmap := map[int]int{}
	current := head
	for current != nil {
		hmap[current.Val]++
		current = current.Next
	}

	current = head
	head = nil
	var pre *ListNode
	// 再次遍历，删除重复节点
	for current != nil {
		// 节点重复
		if hmap[current.Val] == 1 {
			// 这里需要处理如果head节点重复，重新指头节点
			if head == nil && pre == nil {
				head = current
			}
			pre = current
		} else if pre != nil {
			// 删除当前已重复的节点，但需要注意如果pre节点为空，那么此时不需要处理
			pre.Next = current.Next
		}

		current = current.Next
	}
	return head
}

// @lc code=end
