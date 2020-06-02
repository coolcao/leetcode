package main

import "fmt"

func main() {
	str := "abcbadc"
	l := lengthOfLongestSubstring(str)
	fmt.Printf("%v\n", l)
}

/*
给定一个字符串str，找出其中不含有重复字符的最长子串的长度。
例如，str="abcabcdd"，最长不重复子串"abcd"的长度为4。
**/
func lengthOfLongestSubstring(str string) int {
	length := len(str)
	if length == 0 || length == 1 {
		return length
	}

	left, right := 0, 0
	max := right - left

	// 窗口，使用map保存在窗口中的子串
	windowMap := map[byte]bool{}

	for right < length {
		// 窗口右侧边界是否在窗口内
		if !windowMap[str[right]] {
			// 不在窗口内，右侧边界向右移动一格
			windowMap[str[right]] = true
			right++
			// 判断当前窗口长度是否最大
			if right-left > max {
				max = right - left
			}
		} else {
			// 如果在窗口内，遇到重复的，窗口左侧边界移动到重复字符位置
			for left < right {
				// 将左侧边界到重复位置的子串移出窗口
				windowMap[str[left]] = false
				if windowMap[str[left]] == windowMap[str[right]] {
					left++
					break
				}
				left++
			}
		}
	}
	return max
}
