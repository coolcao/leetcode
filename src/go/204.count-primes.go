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
	"time"
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

// ------------------------------采用链表形式进行筛选--------------------
// 使用链表并不能提高筛选速度，相反，使用链表的方式速度是非常低的，就本题而言。
// 因为判断素数是有数理逻辑的，链表由于没有下标，只能遍历，当数据量很大时，效率会非常慢。甚至比一般判断素数方法还慢。
// 上面方法，使用数组，虽然数组的删除操作会很低效，但是可以利用其数理逻辑与下标的对应关系，跳过很多不必要的判断。
// 你品，你细品。
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

// --------------------------- 最原始的判断方法------------------------
func isPrimes(n int) bool {
	if n < 2 {
		return false
	}
	for i := 2; i*i <= n; i++ {
		if n%i == 0 {
			return false
		}
	}
	return true
}
func countPrimes3(n int) int {
	if n <= 2 {
		return 0
	}
	count := 0
	for i := 2; i < n; i++ {
		if isPrimes(i) {
			count++
		}
	}
	return count
}
func main() {
	n := 138331
	fmt.Println(time.Now().UnixNano())
	countPrimes(n)
	fmt.Println(time.Now().UnixNano())
	countPrimes2(n)
	fmt.Println(time.Now().UnixNano())
	countPrimes3(n)
	fmt.Println(time.Now().UnixNano())

}
