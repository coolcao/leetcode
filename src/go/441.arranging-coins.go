/*
 * @lc app=leetcode id=441 lang=golang
 *
 * [441] Arranging Coins
 *
 * https://leetcode.com/problems/arranging-coins/description/
 *
 * algorithms
 * Easy (39.49%)
 * Likes:    283
 * Dislikes: 586
 * Total Accepted:    92K
 * Total Submissions: 232.9K
 * Testcase Example:  '5'
 *
 * You have a total of n coins that you want to form in a staircase shape,
 * where every k-th row must have exactly k coins.
 * ⁠
 * Given n, find the total number of full staircase rows that can be formed.
 *
 * n is a non-negative integer and fits within the range of a 32-bit signed
 * integer.
 *
 * Example 1:
 *
 * n = 5
 *
 * The coins can form the following rows:
 * ¤
 * ¤ ¤
 * ¤ ¤
 *
 * Because the 3rd row is incomplete, we return 2.
 *
 *
 *
 * Example 2:
 *
 * n = 8
 *
 * The coins can form the following rows:
 * ¤
 * ¤ ¤
 * ¤ ¤ ¤
 * ¤ ¤
 *
 * Because the 4th row is incomplete, we return 3.
 *
 *
 */
package main

import (
	"fmt"
	"math"
)

func main() {
	for i := 0; i < 1000000; i++ {
		a1 := arrangeCoins(i)
		a2 := arrangeCoins2(i)
		if a1 != a2 {
			fmt.Printf("i:%d;a1:%d;a2:%d\n", i, a1, a2)
		}
	}
	fmt.Println("ok")
}

// @lc code=start
// 常规的方法，一行一行的检查
// 遇到不符合规范的，返回
func arrangeCoins(n int) int {
	if n == 0 || n == 1 {
		return n
	}
	i := 0
	for n > i {
		n -= i
		i++
	}
	if n == i {
		return i
	}
	return i - 1
}

// 根据其数理逻辑求解
// 前n行总数： i*(i+1)/2
// 因此我们只需要从 Math.sqrt(n*2) 开始即可
func arrangeCoins2(n int) int {
	if n == 0 || n == 1 {
		return n
	}
	end := int(math.Sqrt(float64(n * 2)))
	if end*(end+1)/2 > n {
		return end - 1
	}
	return end
}

// @lc code=end
