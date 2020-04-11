/*
 * @lc app=leetcode id=204 lang=golang
 *
 * [204] Count Primes
Count the number of prime numbers less than a non-negative number, n.

Example:

Input: 10
Output: 4
Explanation: There are 4 prime numbers less than 10, they are 2, 3, 5, 7.
*/
package main

import (
	"fmt"
)

// coutPrimes 计算n之内的素数个数
// 采用埃拉托色尼筛选法。具体参见百科
// 示意图动画：https://img001-10042971.cos.ap-shanghai.myqcloud.com/leetcode/20160703235154990.gif
func countPrimes(n int) int {
	if n <= 2 {
		return 0
	}
	array := make([]bool, n, n)
	for i := 2; i < n; i++ {
		array[i] = true
	}
	for i := 0; i < n; i++ {
		if array[i] {
			for j := i * i; j < n; j += i {
				array[j] = false
			}
		}
	}
	count := 0
	for i := 2; i < n; i++ {
		if array[i] {
			count++
		}
	}
	return count
}

// 采用链表形式进行筛选，减少删除操作开销
type ListNode struct {
	Val  int       `json:"val"`
	Next *ListNode `json:"next"`
}

func createNode(n int) *ListNode {
	return &ListNode{
		Val:  n,
		Next: nil,
	}
}
func countPrimes2(n int) int {
	if n <= 2 {
		return 0
	}
	// 构造初始链表
	nodes := make([]*ListNode, n-2)
	nodes[0] = createNode(2)
	for i := 3; i < n; i++ {
		nodes[i-2] = createNode(i)
		nodes[i-3].Next = nodes[i-2]
	}

	// 做筛选
	count := 0
	head := nodes[0]
	for head != nil {
		pre, current := head, head.Next
		for current != nil {
			// 合数，做删除
			if current.Val%head.Val == 0 {
				pre.Next = current.Next
			} else {
				pre = current
			}
			current = current.Next
		}
		count++
		head = head.Next
	}
	return count
}
func main() {
	n := 499979
	fmt.Println(countPrimes(n))
	fmt.Println(countPrimes2(n))
}
