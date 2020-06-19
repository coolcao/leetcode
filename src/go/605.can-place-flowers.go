/*
 * @lc app=leetcode id=605 lang=golang
 *
 * [605] Can Place Flowers
 *
 * https://leetcode.com/problems/can-place-flowers/description/
 *
 * algorithms
 * Easy (31.63%)
 * Likes:    772
 * Dislikes: 354
 * Total Accepted:    99.5K
 * Total Submissions: 315.2K
 * Testcase Example:  '[1,0,0,0,1]\n1'
 *
 * Suppose you have a long flowerbed in which some of the plots are planted and
 * some are not. However, flowers cannot be planted in adjacent plots - they
 * would compete for water and both would die.
 *
 * Given a flowerbed (represented as an array containing 0 and 1, where 0 means
 * empty and 1 means not empty), and a number n, return if n new flowers can be
 * planted in it without violating the no-adjacent-flowers rule.
 *
 * Example 1:
 * Input: flowerbed = [1,0,0,0,1], n = 1
 * Output: True
 *
 * Example 2:
 * Input: flowerbed = [1,0,0,0,1], n = 2
 * Output: False
 *
 *
 * Note:
 *
 * The input array won't violate no-adjacent-flowers rule.
 * The input array size is in the range of [1, 20000].
 * n is a non-negative integer which won't exceed the input array size.
 */

package main

import "fmt"

func main() {
	flowerbed := []int{0}
	n := 1
	res := canPlaceFlowers(flowerbed, n)
	fmt.Printf("%v\n", res)
}

// @lc code=start
// 是否可以在位置i放置
// 条件是位置i未放置花，且i+1和i-1位置都未放置花
func canPlace(flowerbed []int, i int) bool {
	can := false
	if flowerbed[i] == 1 {
		return can
	}
	length := len(flowerbed)

	if i+1 == length && i-1 == -1 {
		// 判断左右两边是否超边界
		can = true
	} else if i-1 == -1 {
		// 左边超边界
		can = flowerbed[i+1] == 0
	} else if i+1 == length {
		// 右边超边界
		can = flowerbed[i-1] == 0
	} else {
		// 左右两边都没超边界
		can = flowerbed[i-1] == 0 && flowerbed[i+1] == 0
	}
	if can {
		flowerbed[i] = 1
	}

	return can
}

func canPlaceFlowers(flowerbed []int, n int) bool {
	length := len(flowerbed)
	if n > (length+1)/2 {
		return false
	}
	for i := 0; i < length && n > 0; i++ {
		if canPlace(flowerbed, i) {
			n--
		}
	}
	return n <= 0
}

// @lc code=end
