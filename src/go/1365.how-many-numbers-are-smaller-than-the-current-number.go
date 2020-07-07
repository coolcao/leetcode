/*
 * @lc app=leetcode id=1365 lang=golang
 *
 * [1365] How Many Numbers Are Smaller Than the Current Number
 *
 * https://leetcode.com/problems/how-many-numbers-are-smaller-than-the-current-number/description/
 *
 * algorithms
 * Easy (84.60%)
 * Likes:    636
 * Dislikes: 20
 * Total Accepted:    82K
 * Total Submissions: 96.1K
 * Testcase Example:  '[8,1,2,2,3]'
 *
 * Given the array nums, for each nums[i] find out how many numbers in the
 * array are smaller than it. That is, for each nums[i] you have to count the
 * number of valid j's such that j != i and nums[j] < nums[i].
 *
 * Return the answer in an array.
 *
 *
 * Example 1:
 *
 *
 * Input: nums = [8,1,2,2,3]
 * Output: [4,0,1,1,3]
 * Explanation:
 * For nums[0]=8 there exist four smaller numbers than it (1, 2, 2 and 3).
 * For nums[1]=1 does not exist any smaller number than it.
 * For nums[2]=2 there exist one smaller number than it (1).
 * For nums[3]=2 there exist one smaller number than it (1).
 * For nums[4]=3 there exist three smaller numbers than it (1, 2 and 2).
 *
 *
 * Example 2:
 *
 *
 * Input: nums = [6,5,4,8]
 * Output: [2,1,0,3]
 *
 *
 * Example 3:
 *
 *
 * Input: nums = [7,7,7,7]
 * Output: [0,0,0,0]
 *
 *
 *
 * Constraints:
 *
 *
 * 2 <= nums.length <= 500
 * 0 <= nums[i] <= 100
 *
 */
package main

import "fmt"

func main() {
	nums := []int{8, 1, 2, 2, 3}
	ans := smallerNumbersThanCurrent(nums)
	fmt.Printf("%v\n", ans)
}

// @lc code=start
func smallerNumbersThanCurrent(nums []int) []int {
	used := [101]int{}
	min, max := 100, 0
	for _, n := range nums {
		used[n]++
		if n > max {
			max = n
		}
		if n < min {
			min = n
		}
	}
	ans := []int{}
	for _, n := range nums {
		count := 0
		for i := min; i < n; i++ {
			count += used[i]
		}
		ans = append(ans, count)
	}
	return ans
}

// @lc code=end
