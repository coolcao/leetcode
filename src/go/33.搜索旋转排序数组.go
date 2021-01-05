/*
 * @lc app=leetcode.cn id=33 lang=golang
 *
 * [33] 搜索旋转排序数组
 *
 * https://leetcode-cn.com/problems/search-in-rotated-sorted-array/description/
 *
 * algorithms
 * Medium (40.01%)
 * Likes:    1122
 * Dislikes: 0
 * Total Accepted:    205.4K
 * Total Submissions: 513K
 * Testcase Example:  '[4,5,6,7,0,1,2]\n0'
 *
 * 升序排列的整数数组 nums 在预先未知的某个点上进行了旋转（例如， [0,1,2,4,5,6,7] 经旋转后可能变为 [4,5,6,7,0,1,2]
 * ）。
 *
 * 请你在数组中搜索 target ，如果数组中存在这个目标值，则返回它的索引，否则返回 -1 。
 *
 * 示例 1：
 * 输入：nums = [4,5,6,7,0,1,2], target = 0
 * 输出：4
 *
 * 示例 2：
 * 输入：nums = [4,5,6,7,0,1,2], target = 3
 * 输出：-1
 *
 * 示例 3：
 * 输入：nums = [1], target = 0
 * 输出：-1
 *
 * 提示：
 * 1.  1 <= nums.length <= 5000
 * 2.  -10^4 <= nums[i] <= 10^4
 * 3.  nums 中的每个值都 独一无二
 * 4.  nums 肯定会在某个点上旋转
 * 5.  -10^4 <= target <= 10^4
 */

package main

import (
	"fmt"
)

func main() {
	nums := []int{4, 5, 6, 7, 0, 1, 2}
	target := 0
	idx := search(nums, target)
	fmt.Printf("%d\n", idx)
}

// @lc code=start
// binarySearch 二分查找
func binarySearch(nums []int, start, end int, target int) int {
	if start > end {
		return -1
	}
	mid := (start + end) / 2
	if nums[mid] == target {
		return mid
	}
	if nums[mid] > target {
		return binarySearch(nums, start, mid-1, target)
	}
	if nums[mid] < target {
		return binarySearch(nums, mid+1, end, target)
	}
	return -1
}

// _search 查找
func _search(nums []int, start, end int, target int) int {
	if start >= end {
		if nums[start] == target {
			return start
		}
		return -1
	}
	mid := (start + end) / 2
	// 升序区间start, end
	ascStart, ascEnd := start, mid
	// 剩余区间start, end
	nStart, nEnd := mid+1, end
	if nums[mid] < nums[end] {
		ascStart = mid
		ascEnd = end
		nStart = start
		nEnd = mid - 1
	}

	// 判断target是否在升序区间内，如果在，直接二分
	if target >= nums[ascStart] && target <= nums[ascEnd] {
		return binarySearch(nums, ascStart, ascEnd, target)
	}
	// 如果不在，则继续在剩余非升序区间查找
	return _search(nums, nStart, nEnd, target)

}
func search(nums []int, target int) int {
	length := len(nums)
	if length == 0 {
		return -1
	}
	return _search(nums, 0, length-1, target)
}

// @lc code=end
