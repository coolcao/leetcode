package main

import "fmt"

/*
Given a non-negative index k where k ≤ 33, return the kth index row of the Pascal's triangle.

Note that the row index starts from 0.


In Pascal's triangle, each number is the sum of the two numbers directly above it.

Example:

Input: 3
Output: [1,3,3,1]

**********************
生成杨辉三角的第n行
杨辉三角：
[
     [1],
    [1,1],
   [1,2,1],
  [1,3,3,1],
 [1,4,6,4,1]
]
*/

func getRow(rowIndex int) []int {
	if rowIndex == 0 {
		return []int{1}
	}
	if rowIndex == 1 {
		return []int{1, 1}
	}
	numRows := rowIndex + 1
	result := make([][]int, 0)
	result = append(result, []int{1})
	result = append(result, []int{1, 1})
	for row := 3; row <= numRows; row++ {
		tmp := make([]int, row)
		tmp[0], tmp[row-1] = 1, 1
		for i := 1; i < row-1; i++ {
			pre := result[row-2]
			tmp[i] = pre[i] + pre[i-1]
		}
		result = append(result, tmp)
	}
	return result[rowIndex]
}

func getRow2(rowIndex int) []int {
	pre := make([]int, rowIndex+1)
	pre[0], pre[rowIndex] = 1, 1
	if rowIndex == 0 {
		return pre
	}
	if rowIndex == 1 {
		pre[1] = 1
		return pre
	}
	pre[1] = 1
	current := make([]int, rowIndex+1)
	current[0], current[rowIndex] = 1, 1
	for row := 2; row <= rowIndex; row++ {
		for i := 1; i <= row; i++ {
			if i == row {
				current[i] = 1
			} else {
				current[i] = pre[i] + pre[i-1]
			}
		}
		pre, current = current, pre
	}
	return pre
}

func getRow3(rowIndex int) []int {
	length := rowIndex + 1
	nums := make([]int, length*2)
	pre := nums[:length]

	pre[0], pre[rowIndex] = 1, 1
	if rowIndex == 0 || rowIndex == 1 {
		return pre
	}

	pre[1] = 1
	current := nums[length:]
	current[0], current[rowIndex] = 1, 1
	for row := 2; row <= rowIndex; row++ {
		for i := 1; i <= row; i++ {
			if i == row {
				current[i] = 1
			} else {
				current[i] = pre[i] + pre[i-1]
			}
		}
		pre, current = current, pre
	}

	return pre
}

func main() {
	rowIndex := 9
	result2 := getRow2(rowIndex)
	result3 := getRow3(rowIndex)
	fmt.Printf("result2: %v\n", result2)
	fmt.Printf("result3: %v\n", result3)
}
