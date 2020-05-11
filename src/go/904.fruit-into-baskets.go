/*
 * @lc app=leetcode id=904 lang=golang
 *
 * [904] Fruit Into Baskets
 *
 * https://leetcode.com/problems/fruit-into-baskets/description/
 *
 * algorithms
 * Medium (42.15%)
 * Likes:    727
 * Dislikes: 1098
 * Total Accepted:    87.2K
 * Total Submissions: 206.7K
 * Testcase Example:  '[1,2,1]'
 *
 * In a row of trees, the i-th tree produces fruit with type tree[i].
 *
 * You start at any tree of your choice, then repeatedly perform the following
 * steps:
 *
 *
 * 1. Add one piece of fruit from this tree to your baskets.  If you cannot,
 * stop.
 * 2. Move to the next tree to the right of the current tree.  If there is no tree
 * to the right, stop.
 *
 *
 * Note that you do not have any choice after the initial choice of starting
 * tree: you must perform step 1, then step 2, then back to step 1, then step
 * 2, and so on until you stop.
 *
 * You have two baskets, and each basket can carry any quantity of fruit, but
 * you want each basket to only carry one type of fruit each.
 *
 * What is the total amount of fruit you can collect with this procedure?
 *
 *
 *
 * Example 1:
 *
 *
 * Input: [1,2,1]
 * Output: 3
 * Explanation: We can collect [1,2,1].
 *
 *
 *
 * Example 2:
 *
 *
 * Input: [0,1,2,2]
 * Output: 3
 * Explanation: We can collect [1,2,2].
 * If we started at the first tree, we would only collect [0, 1].
 *
 *
 *
 * Example 3:
 *
 *
 * Input: [1,2,3,2,2]
 * Output: 4
 * Explanation: We can collect [2,3,2,2].
 * If we started at the first tree, we would only collect [1, 2].
 *
 *
 *
 * Example 4:
 *
 *
 * Input: [3,3,3,1,2,1,1,2,3,3,4]
 * Output: 5
 * Explanation: We can collect [1,2,1,1,2].
 * If we started at the first tree or the eighth tree, we would only collect 4
 * fruits.
 *
 *
 * Note:
 *
 * 1 <= tree.length <= 40000
 * 0 <= tree[i] < tree.length
 *
 *
 */
package main

import "fmt"

func main() {
	tree := []int{3, 3, 3, 1, 2, 1, 1, 2, 3, 3, 4}
	total := totalFruit(tree)
	fmt.Printf("total: %d\n", total)
}

// @lc code=start
func totalFruit(tree []int) int {
	length := len(tree)

	max := 0
	basketMap := map[int]int{}
	left, right := 0, 0

	for right < length && len(basketMap) <= 2 {
		rightItem := tree[right]
		basketMap[rightItem]++
		right++
		for len(basketMap) > 2 {
			leftItem := tree[left]
			basketMap[leftItem]--
			if basketMap[leftItem] == 0 {
				delete(basketMap, leftItem)
			}
			left++
		}
		current := right - left
		if max < current {
			// fmt.Printf("left: %d,right: %d\n", left, right)
			max = current
		}
	}

	return max
}

// @lc code=end
