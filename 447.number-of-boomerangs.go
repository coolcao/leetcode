/*
 * @lc app=leetcode id=447 lang=golang
 *
 * [447] Number of Boomerangs
 *
 * https://leetcode.com/problems/number-of-boomerangs/description/
 *
 * algorithms
 * Easy (51.34%)
 * Likes:    366
 * Dislikes: 606
 * Total Accepted:    64.7K
 * Total Submissions: 125.9K
 * Testcase Example:  '[[0,0],[1,0],[2,0]]'
 *
 * Given n points in the plane that are all pairwise distinct, a "boomerang" is
 * a tuple of points (i, j, k) such that the distance between i and j equals
 * the distance between i and k (the order of the tuple matters).
 *
 * Find the number of boomerangs. You may assume that n will be at most 500 and
 * coordinates of points are all in the range [-10000, 10000] (inclusive).
 *
 * Example:
 *
 *
 * Input:
 * [[0,0],[1,0],[2,0]]
 *
 * Output:
 * 2
 *
 * Explanation:
 * The two boomerangs are [[1,0],[0,0],[2,0]] and [[1,0],[2,0],[0,0]]
 *
 *
 *
 *
 */
package main

import "fmt"

func main() {
	points := [][]int{[]int{0, 0}, []int{2, 0}, []int{1, 0}, []int{3, 0}}
	c := numberOfBoomerangs(points)
	numberOfBoomerangs2(points)
	fmt.Printf("count: %d\n", c)

}

// @lc code=start
// 方法1: 穷举，把所有可能都列举出来，思路简单，效率低
func numberOfBoomerangs(points [][]int) int {
	length := len(points)
	if length < 3 {
		return 0
	}
	count := 0
	for i := 0; i < length; i++ {
		for j := 0; j < length; j++ {
			for x := 0; x < length; x++ {
				if i == j || i == x || j == x {
					continue
				}
				pi, pj, px := points[i], points[j], points[x]
				if (pi[0]-pj[0])*(pi[0]-pj[0])+(pi[1]-pj[1])*(pi[1]-pj[1]) == (pi[0]-px[0])*(pi[0]-px[0])+(pi[1]-px[1])*(pi[1]-px[1]) {
					count++
				}
			}
		}
	}
	return count
}

func numberOfBoomerangs2(points [][]int) int {
	length := len(points)
	if length < 3 {
		return 0
	}

	distances := make([][]int, length-1)

	// 先把所有点的距离算出来
	for i := 0; i < length-1; i++ {
		tmp := make([]int, length)
		for j := i + 1; j < length; j++ {
			distance := (points[i][0]-points[j][0])*(points[i][0]-points[j][0]) + (points[i][1]-points[j][1])*(points[i][1]-points[j][1])
			tmp[j] = distance
			// fmt.Printf("i:%d,j:%d,distance: %d\n", i, j, distance)
		}
		distances[i] = tmp
	}
	fmt.Printf("%v\n", distances)
	return 0
}

// @lc code=end
