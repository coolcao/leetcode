package main

import (
	"fmt"
)

/*
Given a non-empty array of digits representing a non-negative integer, plus one to the integer.

The digits are stored such that the most significant digit is at the head of the list, and each element in the array contain a single digit.

You may assume the integer does not contain any leading zero, except the number 0 itself.

Example 1:

Input: [1,2,3]
Output: [1,2,4]
Explanation: The array represents the integer 123.
Example 2:

Input: [4,3,2,1]
Output: [4,3,2,2]
Explanation: The array represents the integer 4321.
*/

func plusOne(digits []int) []int {
	length := len(digits)
	digits[length-1]++

	result := make([]int, 0)
	for i := length - 1; i >= 0; i-- {
		if digits[i] > 9 {
			if i != 0 {
				digits[i-1]++
				result = append(result, digits[i]%10)
			} else {
				result = append(result, digits[i]%10)
				result = append(result, 1)
			}
		} else {
			result = append(result, digits[i])
		}
	}

	length = len(result)
	for i := 0; i < length/2; i++ {
		result[i], result[length-1-i] = result[length-1-i], result[i]
	}

	return result
}

func main() {
	nums := []int{1, 9, 9, 9, 7}
	result := plusOne(nums)
	fmt.Printf("%v\n", result)
}
