/*
 * @lc app=leetcode id=1304 lang=golang
 *
 * [1304] Find N Unique Integers Sum up to Zero
 *
 * https://leetcode.com/problems/find-n-unique-integers-sum-up-to-zero/description/
 *
 * algorithms
 * Easy (76.23%)
 * Likes:    266
 * Dislikes: 169
 * Total Accepted:    42.8K
 * Total Submissions: 56.3K
 * Testcase Example:  '5'
 *
 * Given an integer n, return any array containing n unique integers such that
 * they add up to 0.
 *
 *
 * Example 1:
 *
 *
 * Input: n = 5
 * Output: [-7,-1,1,3,4]
 * Explanation: These arrays also are accepted [-5,-1,1,2,3] ,
 * [-3,-1,2,-2,4].
 *
 *
 * Example 2:
 *
 *
 * Input: n = 3
 * Output: [-1,0,1]
 *
 *
 * Example 3:
 *
 *
 * Input: n = 1
 * Output: [0]
 *
 *
 *
 * Constraints:
 *
 *
 * 1 <= n <= 1000
 *
 */
package main

import "fmt"

func main() {
	n := 5
	fmt.Printf("%v\n", sumZero(n))
}

// @lc code=start
func sumZero(n int) []int {
	ans := []int{}
	for i := 1; i < n; i += 2 {
		ans = append(ans, i, -1*i)
	}
	if n%2 == 1 {
		ans = append(ans, 0)
	}
	return ans
}

// @lc code=end
