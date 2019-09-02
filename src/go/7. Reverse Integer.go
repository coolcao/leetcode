/**
Given a 32-bit signed integer, reverse digits of an integer.

Example 1:

Input: 123
Output: 321
Example 2:

Input: -123
Output: -321
Example 3:

Input: 120
Output: 21
Note:
Assume we are dealing with an environment which could only store integers within the 32-bit signed integer range: [−231,  231 − 1]. For the purpose of this problem, assume that your function returns 0 when the reversed integer overflows.
*/

package main

import (
	"fmt"
)

func reverse(x int) int {
	tag := 1000000000
	min := [10]int{2, 1, 4, 7, 4, 8, 3, 6, 4, 8}
	max := [10]int{2, 1, 4, 7, 4, 8, 3, 6, 4, 7}
	currentArr := &max

	resultArr := make([]int, 0)

	// 标志位，根据标志位设置正负号， 以及使用最大值最小值进行比较
	flag := 1
	if x < 0 {
		currentArr = &min
		flag = -1
		x = -x
	}

	for tag > x {
		m := x / tag
		tag = tag / 10
		resultArr = append(resultArr, m)

	}
	for x > 0 {
		m := x % 10
		x /= 10
		resultArr = append(resultArr, m)
	}

	if len(resultArr) > 10 {
		fmt.Println("位数已溢出")
		return 0
	}

	// 判断溢出，并计算逆转后的结果
	current := 0
	result := 0
	for i := 0; i < 10; i++ {
		if currentArr[i] == resultArr[i] {
			current++
		}
		if current == i && resultArr[i] > currentArr[i] {
			return 0
		}
		result = result*10 + resultArr[i]

	}

	return flag * result
}

func main() {
	num := -23876378623423
	fmt.Printf("%d\n", num)
	result := reverse(num)
	fmt.Printf("%d\n", result)
}
