/*
 * @lc app=leetcode id=766 lang=golang
 *
 * [766] Toeplitz Matrix
 *
 * https://leetcode.com/problems/toeplitz-matrix/description/
 *
 * algorithms
 * Easy (63.86%)
 * Likes:    988
 * Dislikes: 81
 * Total Accepted:    95.3K
 * Total Submissions: 147.2K
 * Testcase Example:  '[[1,2,3,4],[5,1,2,3],[9,5,1,2]]'
 *
 * A matrix is Toeplitz if every diagonal from top-left to bottom-right has the
 * same element.
 *
 * Now given an M x N matrix, return True if and only if the matrix is
 * Toeplitz.
 *
 *
 * Example 1:
 *
 *
 * Input:
 * matrix = [
 * [1,2,3,4],
 * [5,1,2,3],
 * [9,5,1,2]
 * ]
 * Output: True
 * Explanation:
 * In the above grid, the diagonals are:
 * "[9]", "[5, 5]", "[1, 1, 1]", "[2, 2, 2]", "[3, 3]", "[4]".
 * In each diagonal all elements are the same, so the answer is True.
 *
 *
 * Example 2:
 *
 *
 * Input:
 * matrix = [
 * [1,2],
 * [2,2]
 * ]
 * Output: False
 * Explanation:
 * The diagonal "[1, 2]" has different elements.
 *
 *
 *
 * Note:
 *
 *
 * matrix will be a 2D array of integers.
 * matrix will have a number of rows and columns in range [1, 20].
 * matrix[i][j] will be integers in range [0, 99].
 *
 *
 *
 * Follow up:
 *
 *
 * What if the matrix is stored on disk, and the memory is limited such that
 * you can only load at most one row of the matrix into the memory at once?
 * What if the matrix is so large that you can only load up a partial row into
 * the memory at once?
 *
 *
 */
package main

import "fmt"

func main() {
	matrix := [][]int{[]int{1, 2}, []int{2, 1}}
	b := isToeplitzMatrix(matrix)
	fmt.Printf("%v\n", b)
}

// @lc code=start
func isToeplitzMatrix(matrix [][]int) bool {
	rows, cols := len(matrix), len(matrix[0])
	if rows == 1 || cols == 1 {
		return true
	}

	for r := 1; r < rows; r++ {
		for c := 1; c < cols; c++ {
			if matrix[r][c] != matrix[r-1][c-1] {
				return false
			}
		}
	}

	return true
}

// @lc code=end
