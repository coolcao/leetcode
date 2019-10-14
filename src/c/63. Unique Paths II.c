/*
A robot is located at the top-left corner of a m x n grid (marked 'Start' in the
diagram below).

![图片](https://assets.leetcode.com/uploads/2018/10/22/robot_maze.png)

The robot can only move either down or right at any point in time. The robot is
trying to reach the bottom-right corner of the grid (marked 'Finish' in the
diagram below).

Now consider if some obstacles are added to the grids. How many unique paths
would there be?



An obstacle and empty space is marked as 1 and 0 respectively in the grid.

Note: m and n will be at most 100.

Example 1:

Input:
[
  [0,0,0],
  [0,1,0],
  [0,0,0]
]
Output: 2
Explanation:
There is one obstacle in the middle of the 3x3 grid above.
There are two ways to reach the bottom-right corner:
1. Right -> Right -> Down -> Down
2. Down -> Down -> Right -> Right
*/

#include <stdio.h>
#include <stdlib.h>
int uniquePathsWithObstacles(int** obstacleGrid, int obstacleGridSize,
                             int* obstacleGridColSize) {
  if (obstacleGridSize == 0) return 0;
  int rows = obstacleGridSize, cols = obstacleGridColSize[0];
  unsigned int** steps;
  steps = (int**)malloc(sizeof(int*) * rows);
  for (int i = 0; i < rows; i++) {
    steps[i] = (int*)malloc(sizeof(int) * cols);
    for (int j = 0; j < cols; j++) {
      steps[i][j] = 0;
    }
  }

  if (cols > 0 && obstacleGrid[0][0] == 0) {
    steps[0][0] = 1;
  }

  for (int i = 0; i < rows; i++) {
    for (int j = 0; j < cols; j++) {
      int current = obstacleGrid[i][j];
      if (current == 1 || (i == 0 && j == 0)) {
        continue;
      }
      if (i == 0) {
        steps[i][j] = steps[0][j] + steps[i][j - 1];
      } else if (j == 0) {
        steps[i][j] = steps[i - 1][j] + steps[i][0];
      } else {
        steps[i][j] = steps[i - 1][j] + steps[i][j - 1];
      }
    }
  }

  return steps[rows-1][cols-1];
}

