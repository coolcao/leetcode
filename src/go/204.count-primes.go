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

import "fmt"

// coutPrimes 计算n之内的素数个数
// 采用埃拉托色尼筛选法。具体参见百科
// 示意图动画：https://img001-10042971.cos.ap-shanghai.myqcloud.com/leetcode/20160703235154990.gif
func countPrimes(n int) int {
	if n <= 1 {
		return 0
	}
	array := make([]bool, n, n)
	for i := 2; i <= n; i++ {
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

func main() {
	num := 2
	c := countPrimes(num)
	fmt.Println(c)
}
