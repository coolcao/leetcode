package main

import "fmt"

/*
Given a non-negative integer numRows, generate the first numRows of Pascal's triangle.


In Pascal's triangle, each number is the sum of the two numbers directly above it.

Example:

Input: 5
Output:
[
     [1],
    [1,1],
   [1,2,1],
  [1,3,3,1],
 [1,4,6,4,1]
]
*********************
生成杨辉三角
给一个非负整数numRows，生成一个含有numRows行的杨辉三角。
*********************

*/

func generate(numRows int) [][]int {
	if numRows == 0 {
		return nil
	}
	result := make([][]int, 0)
	if numRows == 1 {
		return [][]int{[]int{1}}
	}
	if numRows == 2 {
		return [][]int{[]int{1}, []int{1, 1}}
	}
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
	return result
}

func main() {
	numRows := 2
	result := generate(numRows)
	fmt.Printf("%v\n", result)
}
