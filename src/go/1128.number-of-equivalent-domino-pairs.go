/*
 * @lc app=leetcode id=1128 lang=golang
 *
 * [1128] Number of Equivalent Domino Pairs
 *
 * https://leetcode.com/problems/number-of-equivalent-domino-pairs/description/
 *
 * algorithms
 * Easy (47.27%)
 * Likes:    205
 * Dislikes: 111
 * Total Accepted:    23.3K
 * Total Submissions: 49.2K
 * Testcase Example:  '[[1,2],[2,1],[3,4],[5,6]]'
 *
 * Given a list of dominoes, dominoes[i] = [a, b] is equivalent to dominoes[j]
 * = [c, d] if and only if either (a==c and b==d), or (a==d and b==c) - that
 * is, one domino can be rotated to be equal to another domino.
 *
 * Return the number of pairs (i, j) for which 0 <= i < j < dominoes.length,
 * and dominoes[i] is equivalent to dominoes[j].
 *
 *
 * Example 1:
 * Input: dominoes = [[1,2],[2,1],[3,4],[5,6]]
 * Output: 1
 *
 *
 * Constraints:
 *
 *
 * 1 <= dominoes.length <= 40000
 * 1 <= dominoes[i][j] <= 9
 *
 */
package main

func main() {
	dominoes := [][]int{[]int{1, 2}, []int{2, 1}, []int{3, 4}, []int{5, 6}}
	numEquivDominoPairs(dominoes)
}

// @lc code=start
func getSum(nums []int) int {
	a, b := nums[0], nums[1]
	if nums[0] > nums[1] {
		a, b = nums[1], nums[0]
	}
	return 10*a + b
}
func numEquivDominoPairs(dominoes [][]int) int {
	length := len(dominoes)
	counts := [100]int{}
	sum := 0
	for i := 0; i < length; i++ {
		count := getSum(dominoes[i])
		counts[count]++
	}
	for _, c := range counts {
		if c > 1 {
			sum += c * (c - 1) / 2
		}
	}
	return sum
}

// @lc code=end
