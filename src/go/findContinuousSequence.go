/*
输入一个正整数 target ，输出所有和为 target 的连续正整数序列（至少含有两个数）。
序列内的数字由小到大排列，不同序列按照首个数字从小到大排列。

示例 1：
输入：target = 9
输出：[[2,3,4],[4,5]]

示例 2：
输入：target = 15
输出：[[1,2,3,4,5],[4,5,6],[7,8]]

限制：
1 <= target <= 10^5

**/

package main

import "fmt"

func main() {
	target := 54
	res := findContinuousSequence(target)
	res2 := findContinuousSequence2(target)
	fmt.Printf("%v = %v\n", res, res2)
}
func findContinuousSequence(target int) [][]int {
	// 由于最少两个元素，因此，我们只需要判断到其中间位置即可
	end := target/2 + 1
	left, right := 1, 1
	sum := 0
	res := [][]int{}
	for right <= end {
		// 入窗
		sum += right
		right++
		for left < right && sum >= target {
			if sum == target {
				tmp := []int{}
				for i := left; i < right; i++ {
					tmp = append(tmp, i)
				}
				res = append(res, tmp)
			}
			sum -= left
			left++
		}
	}
	return res
}
func findContinuousSequence2(target int) [][]int {
	// 记录窗口内元素和
	sum := 0
	left, right := 1, 3
	for i := 1; i < right; i++ {
		sum += i
	}

	result := [][]int{}
	for left <= target-1 {
		if sum == target {
			tmp := make([]int, right-left)
			for i := left; i < right; i++ {
				tmp[i-left] = i
			}
			result = append(result, tmp)
			// 窗口向右滑动
			left, right = left+1, left+3
			sum = (left + left + 1)
		} else if sum < target {
			sum += right
			right++
		} else if sum > target {
			sum -= left
			left++
		}

		if right-left == 2 && sum > target {
			break
		}

	}

	return result
}
