/*
 * @lc app=leetcode id=1337 lang=golang
 *
 * [1337] The K Weakest Rows in a Matrix
 *
 * https://leetcode.com/problems/the-k-weakest-rows-in-a-matrix/description/
 *
 * algorithms
 * Easy (68.64%)
 * Likes:    156
 * Dislikes: 7
 * Total Accepted:    17.5K
 * Total Submissions: 25.7K
 * Testcase Example:  '[[1,1,0,0,0],[1,1,1,1,0],[1,0,0,0,0],[1,1,0,0,0],[1,1,1,1,1]]\n3'
 *
 * Given a m * n matrix mat of ones (representing soldiers) and zeros
 * (representing civilians), return the indexes of the k weakest rows in the
 * matrix ordered from the weakest to the strongest.
 *
 * A row i is weaker than row j, if the number of soldiers in row i is less
 * than the number of soldiers in row j, or they have the same number of
 * soldiers but i is less than j. Soldiers are always stand in the frontier of
 * a row, that is, always ones may appear first and then zeros.
 *
 *
 * Example 1:
 *
 *
 * Input: mat =
 * [[1,1,0,0,0],
 * ⁠[1,1,1,1,0],
 * ⁠[1,0,0,0,0],
 * ⁠[1,1,0,0,0],
 * ⁠[1,1,1,1,1]],
 * k = 3
 * Output: [2,0,3]
 * Explanation:
 * The number of soldiers for each row is:
 * row 0 -> 2
 * row 1 -> 4
 * row 2 -> 1
 * row 3 -> 2
 * row 4 -> 5
 * Rows ordered from the weakest to the strongest are [2,0,3,1,4]
 *
 *
 * Example 2:
 *
 *
 * Input: mat =
 * [[1,0,0,0],
 * [1,1,1,1],
 * [1,0,0,0],
 * [1,0,0,0]],
 * k = 2
 * Output: [0,2]
 * Explanation:
 * The number of soldiers for each row is:
 * row 0 -> 1
 * row 1 -> 4
 * row 2 -> 1
 * row 3 -> 1
 * Rows ordered from the weakest to the strongest are [0,2,3,1]
 *
 *
 *
 * Constraints:
 *
 *
 * m == mat.length
 * n == mat[i].length
 * 2 <= n, m <= 100
 * 1 <= k <= m
 * matrix[i][j] is either 0 or 1.
 *
 */
package main

import "fmt"

func main() {
	mat := [][]int{
		[]int{1, 1, 0, 0, 0},
		[]int{1, 1, 1, 1, 0},
		[]int{1, 0, 0, 0, 0},
		[]int{1, 1, 0, 0, 0},
		[]int{1, 1, 1, 1, 1},
	}
	k := 3
	res := kWeakestRows(mat, k)
	fmt.Printf("%v\n", res)
}

// @lc code=start
func kWeakestRows(mat [][]int, k int) []int {
	m, n := len(mat), len(mat[0])
	soldiersCount := make([]int, m)
	for i := 0; i < m; i++ {
		values := mat[i]
		for j := 0; j < n; j++ {
			if values[j] == 1 {
				soldiersCount[i]++
			} else {
				break
			}
		}
	}

	tmp := make([][]int, 101)
	for i, v := range soldiersCount {
		if tmp[v] == nil {
			tmp[v] = []int{i}
		} else {
			tmp[v] = append(tmp[v], i)
		}
	}
	res := []int{}
	for _, v := range tmp {
		if len(v) > 0 && k > 0 {
			for _, e := range v {
				if k > 0 {
					res = append(res, e)
					k--

				}
			}
		}
	}

	return res
}

// @lc code=end
