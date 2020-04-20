/*
 * @lc app=leetcode id=1351 lang=golang
 *
 * [1351] Count Negative Numbers in a Sorted Matrix
 *
 * https://leetcode.com/problems/count-negative-numbers-in-a-sorted-matrix/description/
 *
 * algorithms
 * Easy (77.78%)
 * Likes:    214
 * Dislikes: 16
 * Total Accepted:    31.1K
 * Total Submissions: 40.1K
 * Testcase Example:  '[[4,3,2,-1],[3,2,1,-1],[1,1,-1,-2],[-1,-1,-2,-3]]'
 *
 * Given a m * n matrix grid which is sorted in non-increasing order both
 * row-wise and column-wise.
 *
 * Return the number of negative numbers in grid.
 *
 *
 * Example 1:
 *
 *
 * Input: grid = [[4,3,2,-1],[3,2,1,-1],[1,1,-1,-2],[-1,-1,-2,-3]]
 * Output: 8
 * Explanation: There are 8 negatives number in the matrix.
 *
 *
 * Example 2:
 *
 *
 * Input: grid = [[3,2],[1,0]]
 * Output: 0
 *
 *
 * Example 3:
 *
 *
 * Input: grid = [[1,-1],[-1,-1]]
 * Output: 3
 *
 *
 * Example 4:
 *
 *
 * Input: grid = [[-1]]
 * Output: 1
 *
 *
 *
 * Constraints:
 *
 *
 * m == grid.length
 * n == grid[i].length
 * 1 <= m, n <= 100
 * -100 <= grid[i][j] <= 100
 *
 */
package main

import "fmt"

func main() {
	grid := [][]int{[]int{4, 3, 2, -1}, []int{3, 2, -1, -1}, []int{1, 1, -1, -2}, []int{-1, -1, -2, -3}}
	c := countNegatives(grid)
	fmt.Printf("count: %d\n", c)
}

// @lc code=start
func countNegatives(grid [][]int) int {
	m, n := len(grid), len(grid[0])
	row, col := m, n
	count := 0
L1:
	for i := 0; i < row; i++ {
		for j := 0; j < n; j++ {
			if grid[i][j] < 0 {
				col = j
				if col == 0 {
					row = i
				}
				continue L1
			}
			count++
		}
	}

	return m*n - count
}

// @lc code=end
