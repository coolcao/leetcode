/*
 * @lc app=leetcode id=1380 lang=golang
 *
 * [1380] Lucky Numbers in a Matrix
 *
 * https://leetcode.com/problems/lucky-numbers-in-a-matrix/description/
 *
 * algorithms
 * Easy (72.32%)
 * Likes:    122
 * Dislikes: 10
 * Total Accepted:    14.9K
 * Total Submissions: 20.6K
 * Testcase Example:  '[[3,7,8],[9,11,13],[15,16,17]]'
 *
 * Given a m * n matrix of distinct numbers, return all lucky numbers in the
 * matrix in any order.
 *
 * A lucky number is an element of the matrix such that it is the minimum
 * element in its row and maximum in its column.
 *
 *
 * Example 1:
 *
 *
 * Input: matrix = [[3,7,8],[9,11,13],[15,16,17]]
 * Output: [15]
 * Explanation: 15 is the only lucky number since it is the minimum in its row
 * and the maximum in its column
 *
 *
 * Example 2:
 *
 *
 * Input: matrix = [[1,10,4,2],[9,3,8,7],[15,16,17,12]]
 * Output: [12]
 * Explanation: 12 is the only lucky number since it is the minimum in its row
 * and the maximum in its column.
 *
 *
 * Example 3:
 *
 *
 * Input: matrix = [[7,8],[1,2]]
 * Output: [7]
 *
 *
 *
 * Constraints:
 *
 *
 * m == mat.length
 * n == mat[i].length
 * 1 <= n, m <= 50
 * 1 <= matrix[i][j] <= 10^5.
 * All elements in the matrix are distinct.
 *
 */
package main

import "fmt"

func main() {
	matrix := [][]int{[]int{3, 2, 8}, []int{9, 7, 1}, []int{11, 10, 19}, []int{89, 39, 27}, []int{54, 21, 90}, []int{47, 38, 28}}
	res := luckyNumbers(matrix)
	fmt.Printf("lucy nums: %v\n", res)
}

// @lc code=start
func luckyNumbers(matrix [][]int) []int {
	m, n := len(matrix), len(matrix[0])
	max, min := make([]int, n), make([]int, m)
	for i := 0; i < m; i++ {
		min[i] = 100001
	}
	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			if matrix[i][j] > max[j] {
				max[j] = matrix[i][j]
			}
			if matrix[i][j] < min[i] {
				min[i] = matrix[i][j]
			}
		}
	}
	// 交点
	result := []int{}
	for i := 0; i < n; i++ {
		for j := 0; j < m; j++ {
			if max[i] == min[j] {
				result = append(result, max[i])
			}
		}
	}
	return result
}

// @lc code=end
