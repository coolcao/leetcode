/*
给你 n 个非负整数 a1，a2，...，an，每个数代表坐标中的一个点 (i, ai) 。
在坐标内画 n 条垂直线，垂直线 i 的两个端点分别为 (i, ai) 和 (i, 0)。
找出其中的两条线，使得它们与 x 轴共同构成的容器可以容纳最多的水。

说明：你不能倾斜容器，且 n 的值至少为 2。

![](https://img001-10042971.cos.ap-shanghai.myqcloud.com/leetcode/question_11.jpg)

图中垂直线代表输入数组 [1,8,6,2,5,4,8,3,7]。在此情况下，容器能够容纳水（表示为蓝色部分）的最大值为 49。

示例：
输入：[1,8,6,2,5,4,8,3,7]
输出：49
*/
package main

import "fmt"

// max 求最大值或最小值。默认求最大值，若min为true则返回最小值
func max(a, b int, min bool) int {
	if a > b {
		if min {
			return b
		}
		return a
	}
	if min {
		return a
	}
	return b
}
func maxArea(height []int) int {
	length := len(height)
	start, end := 0, length-1
	pre, next := start, end
	maxArea := max(height[0], height[length-1], true) * (length - 1)

	for start < end {
		hstart, hend := height[start], height[end]
		if hstart < height[pre] {
			pre = start
			start++
			continue
		}
		if hend < height[next] {
			next = end
			end--
			continue
		}
		h := max(hstart, hend, true)
		l := end - start
		area := h * l

		maxArea = max(maxArea, area, false)

		// 这个地方很有意思，移动start还是end看height[start]大还是height[end]大
		// 这里肯定移动小的，因为这样“损失”更小
		if hstart > hend {
			end--
		} else {
			start++
		}
	}

	return maxArea
}

func main() {
	nums := []int{1, 8, 6, 2, 5, 4, 8, 3, 7}
	area := maxArea(nums)
	fmt.Printf("maxArea: %d\n", area)
}

// @lc code=end
