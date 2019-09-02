package main

import "fmt"

/*
Determine whether an integer is a palindrome. An integer is a palindrome when it reads the same backward as forward.

Example 1:

Input: 121
Output: true
Example 2:

Input: -121
Output: false
Explanation: From left to right, it reads -121. From right to left, it becomes 121-. Therefore it is not a palindrome.
Example 3:

Input: 10
Output: false
Explanation: Reads 01 from right to left. Therefore it is not a palindrome.
Follow up:

Coud you solve it without converting the integer to a string?
*/

func isPalindrome(x int) bool {
	// 如果是负数，肯定不是回文
	if x < 0 {
		return false
	}

	nums := make([]int, 0)

	for x > 0 {
		m := x % 10
		x /= 10
		nums = append(nums, m)
	}

	fmt.Printf("%v\n", nums)
	length := len(nums)
	for i := 0; i < length; i++ {
		if nums[i] != nums[length-1-i] {
			return false
		}
	}

	return true
}

func main() {
	num := 11111111111
	is := isPalindrome(num)
	fmt.Printf("%v\n", is)
}
