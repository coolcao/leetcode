/*
 * @lc app=leetcode id=1232 lang=golang
 *
 * [1232] Check If It Is a Straight Line
 *
 * https://leetcode.com/problems/check-if-it-is-a-straight-line/description/
 *
 * algorithms
 * Easy (46.73%)
 * Likes:    350
 * Dislikes: 57
 * Total Accepted:    78.2K
 * Total Submissions: 170.7K
 * Testcase Example:  '[[1,2],[2,3],[3,4],[4,5],[5,6],[6,7]]'
 *
 * You are given an array coordinates, coordinates[i] = [x, y], where [x, y]
 * represents the coordinate of a point. Check if these points make a straight
 * line in the XY plane.
 *
 *
 *
 *
 * Example 1:
 *
 *
 *
 *
 * Input: coordinates = [[1,2],[2,3],[3,4],[4,5],[5,6],[6,7]]
 * Output: true
 *
 *
 * Example 2:
 *
 *
 *
 *
 * Input: coordinates = [[1,1],[2,2],[3,4],[4,5],[5,6],[7,7]]
 * Output: false
 *
 *
 *
 * Constraints:
 *
 *
 * 2 <= coordinates.length <= 1000
 * coordinates[i].length == 2
 * -10^4 <= coordinates[i][0], coordinates[i][1] <= 10^4
 * coordinates contains no duplicate point.
 *
 *
 */
package main

import "fmt"

func main() {
	coordinates := [][]int{[]int{2, 4}, []int{2, 5}, []int{2, 8}}
	fmt.Printf("%v\n", checkStraightLine(coordinates))
}

// @lc code=start
func checkStraightLine(coordinates [][]int) bool {
	length := len(coordinates)
	start, end := coordinates[0], coordinates[length-1]

	for i := 1; i < length-1; i++ {
		point := coordinates[i]
		if (start[1]-end[1])*(end[0]-point[0]) != (end[1]-point[1])*(start[0]-end[0]) {
			return false
		}
	}
	return true

}

// @lc code=end
