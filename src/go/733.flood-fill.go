/*
 * @lc app=leetcode id=733 lang=golang
 *
 * [733] Flood Fill
 *
 * https://leetcode.com/problems/flood-fill/description/
 *
 * algorithms
 * Easy (53.34%)
 * Likes:    1172
 * Dislikes: 193
 * Total Accepted:    162.7K
 * Total Submissions: 295.9K
 * Testcase Example:  '[[1,1,1],[1,1,0],[1,0,1]]\n1\n1\n2'
 *
 *
 * An image is represented by a 2-D array of integers, each integer
 * representing the pixel value of the image (from 0 to 65535).
 *
 * Given a coordinate (sr, sc) representing the starting pixel (row and column)
 * of the flood fill, and a pixel value newColor, "flood fill" the image.
 *
 * To perform a "flood fill", consider the starting pixel, plus any pixels
 * connected 4-directionally to the starting pixel of the same color as the
 * starting pixel, plus any pixels connected 4-directionally to those pixels
 * (also with the same color as the starting pixel), and so on.  Replace the
 * color of all of the aforementioned pixels with the newColor.
 *
 * At the end, return the modified image.
 *
 * Example 1:
 *
 * Input:
 * image = [[1,1,1],[1,1,0],[1,0,1]]
 * sr = 1, sc = 1, newColor = 2
 * Output: [[2,2,2],[2,2,0],[2,0,1]]
 * Explanation:
 * From the center of the image (with position (sr, sc) = (1, 1)), all pixels
 * connected
 * by a path of the same color as the starting pixel are colored with the new
 * color.
 * Note the bottom corner is not colored 2, because it is not 4-directionally
 * connected
 * to the starting pixel.
 *
 *
 *
 * Note:
 * The length of image and image[0] will be in the range [1, 50].
 * The given starting pixel will satisfy 0 <= sr < image.length  and 0 <= sc < image[0].length .
 * The value of each color in image[i][j] and newColor will be an integer in
 * [0, 65535].
 *
 */
package main

import "fmt"

func main() {
	image := [][]int{[]int{1, 1, 1, 1, 0}, []int{1, 1, 0, 1, 1}, []int{1, 0, 1, 0, 1}}
	newColor := 2
	sr, sc := 1, 1
	floodFill(image, sr, sc, newColor)
	fmt.Printf("%v\n", image)
}

// @lc code=start
func fill(image [][]int, rows, cols, sr, sc int, oldColor, newColor int, visited map[string]bool) {
	k := fmt.Sprintf("%d-%d", sr, sc)
	if sr >= 0 && sr < rows && sc >= 0 && sc < cols && oldColor == image[sr][sc] && !visited[k] {
		image[sr][sc] = newColor
		visited[k] = true
		fill(image, rows, cols, sr+1, sc, oldColor, newColor, visited)
		fill(image, rows, cols, sr-1, sc, oldColor, newColor, visited)
		fill(image, rows, cols, sr, sc+1, oldColor, newColor, visited)
		fill(image, rows, cols, sr, sc-1, oldColor, newColor, visited)
	}
}
func floodFill(image [][]int, sr int, sc int, newColor int) [][]int {
	rows, cols := len(image), len(image[0])
	visited := map[string]bool{}
	oldColor := image[sr][sc]
	fill(image, rows, cols, sr, sc, oldColor, newColor, visited)
	return image
}

// @lc code=end
