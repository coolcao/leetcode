/*
 * @lc app=leetcode id=1399 lang=golang
 *
 * [1399] Count Largest Group
 *
 * https://leetcode.com/problems/count-largest-group/description/
 *
 * algorithms
 * Easy (62.91%)
 * Likes:    42
 * Dislikes: 69
 * Total Accepted:    7.4K
 * Total Submissions: 11.6K
 * Testcase Example:  '13\r'
 *
 * Given an integer n. Each number from 1 to n is grouped according to the sum
 * of its digits.
 *
 * Return how many groups have the largest size.
 *
 *
 * Example 1:
 *
 *
 * Input: n = 13
 * Output: 4
 * Explanation: There are 9 groups in total, they are grouped according sum of
 * its digits of numbers from 1 to 13:
 * [1,10], [2,11], [3,12], [4,13], [5], [6], [7], [8], [9]. There are 4 groups
 * with largest size.
 *
 *
 * Example 2:
 *
 *
 * Input: n = 2
 * Output: 2
 * Explanation: There are 2 groups [1], [2] of size 1.
 *
 *
 * Example 3:
 *
 *
 * Input: n = 15
 * Output: 6
 *
 *
 * Example 4:
 *
 *
 * Input: n = 24
 * Output: 5
 *
 *
 *
 * Constraints:
 *
 *
 * 1 <= n <= 10^4
 *
 *
 */
package main

import "fmt"

func main() {
	n := 13
	c := countLargestGroup(n)
	fmt.Printf("%v\n", c)
}

// @lc code=start
func getDigitSum(num int) int {
	sum := 0
	for num > 0 {
		mod := num % 10
		num = num / 10
		sum += mod
	}
	return sum
}

func countLargestGroup(n int) int {
	maxDigitCount := 0
	if n <= 9 {
		maxDigitCount = 9
	}
	if n <= 99 {
		maxDigitCount = 18
	}
	if n <= 999 {
		maxDigitCount = 27
	}
	if n <= 9999 {
		maxDigitCount = 36
	}

	tmp := make([]int, maxDigitCount+1)

	for i := 1; i <= n; i++ {
		digitSum := getDigitSum(i)
		tmp[digitSum]++
	}

	max := 0
	count := 0
	for i := 1; i <= maxDigitCount; i++ {
		if tmp[i] == max {
			count++
		}
		if tmp[i] > max {
			max = tmp[i]
			count = 1
		}
	}

	return count

}

// @lc code=end
