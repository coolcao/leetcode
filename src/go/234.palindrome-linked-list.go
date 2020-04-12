/*
 * @lc app=leetcode id=234 lang=golang
 *
 * [234] Palindrome Linked List
 *
 * https://leetcode.com/problems/palindrome-linked-list/description/
 *
 * algorithms
 * Easy (38.28%)
 * Likes:    2624
 * Dislikes: 332
 * Total Accepted:    374K
 * Total Submissions: 976.5K
 * Testcase Example:  '[1,2]'
 *
 * Given a singly linked list, determine if it is a palindrome.
 *
 * Example 1:
 *
 *
 * Input: 1->2
 * Output: false
 *
 * Example 2:
 *
 *
 * Input: 1->2->2->1
 * Output: true
 *
 * Follow up:
 * Could you do it in O(n) time and O(1) space?
 *
 */

package main

// @lc code=start
// ListNode Definition for singly-linked list.
// type ListNode struct {
// 	Val  int
// 	Next *ListNode
// }

// O(n) time, O(n) space
func isPalindrome(head *ListNode) bool {
	if head == nil {
		return true
	}
	length := 0
	data := []int{}
	for head != nil {
		data = append(data, head.Val)
		length++
		head = head.Next
	}
	i, j := 0, length-1
	for i <= j {
		if data[i] != data[j] {
			return false
		}
		i++
		j--
	}
	return true
}

// func main() {
// 	n1 := &ListNode{Val: 1, Next: nil}
// 	n2 := &ListNode{Val: 2, Next: nil}
// 	n3 := &ListNode{Val: 1, Next: nil}
// 	n4 := &ListNode{Val: 1, Next: nil}

// 	n1.Next = n2
// 	n2.Next = n3
// 	n3.Next = n4

// 	fmt.Println(isPalindrome(n1))
// }

// @lc code=end
