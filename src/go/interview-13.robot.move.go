/*
地上有一个m行n列的方格，从坐标[0,0]到坐标[m-1,n-1].一个机器人从坐标[0,0]的格子开始移动，
它每次可以向左，右，上，下移动一格（不能移动到方格外），也不能进入行坐标和列坐标的数位之和大于k的格子。
例如，当k为18时，机器人能够进入方格[35,37]，因为3+5+3+7=18.但它不能进入方格[35,38]，因为3+5+3+8=19。
请问该机器人能够达到多少个格子？

示例1:
输入:m=2,n=3,k=1
输出：3

示例2:
输入：m=3,n=1,k=0
输出：1

提示：
1 <= n,m <= 100
0 <= k <= 20

**/

package main

import (
	"fmt"
)

// 检查当前格子是否符合条件
func check(x, y int, m, n, k int, matrix *[100][100]int, count *int) {
	// 检查是否符合方格条件
	if x < 0 || y < 0 {
		return
	}
	// 检查是否已超出方格
	if x >= m || y >= n {
		return
	}

	if matrix[x][y] == 1 {
		return
	}

	sum := 0
	sum += (x/10 + x%10)
	sum += (y/10 + y%10)

	if sum > k {
		return
	}

	matrix[x][y] = 1
	*count++

	check(x+1, y, m, n, k, matrix, count)
	check(x, y+1, m, n, k, matrix, count)
	check(x-1, y, m, n, k, matrix, count)
	check(x, y-1, m, n, k, matrix, count)
}

func movingCount(m, n, k int) int {
	count := 0
	var matrix [100][100]int
	check(0, 0, m, n, k, &matrix, &count)
	return count
}

func main() {
	m, n, k := 100, 100, 20
	count := movingCount(m, n, k)
	fmt.Println(count)
}
