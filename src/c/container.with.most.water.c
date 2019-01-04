/**
Given n non-negative integers a1, a2, ..., an , where each represents a point at
coordinate (i, ai). n vertical lines are drawn such that the two endpoints of
line i is at (i, ai) and (i, 0). Find two lines, which together with x-axis
forms a container, such that the container contains the most water.

Note: You may not slant the container and n is at least 2.

The above vertical lines are represented by array [1,8,6,2,5,4,8,3,7]. In this
case, the max area of water (blue section) the container can contain is 49.



Example:

Input: [1,8,6,2,5,4,8,3,7]
Output: 49
 */

/**
 * 这个题其实很好理解，但是却安排了中等难度。
 * 第一个参数，整数数组，表示一个有序的柱子的高度，然后根据这个柱子，求解哪两根柱子组成的“容器”的容积最大。
 * 最直观容易理解的办法就是，组合所有的柱子，将所有可能的“容器”的容积都计算一遍，找最大值即可。这样肯定效率低下。但能解。
 */

#include <stdio.h>
int maxArea(int* height, int heightSize) {
  int area = 0;
  for (int i = 0; i < heightSize - 1; i++) {
    for (int j = i + 1; j < heightSize; j++) {
      int h1 = height[i];
      int h2 = height[j];
      int h = h1 > h2 ? h2 : h1;
      int l = j - i;
      int a = h * l;
      if (a > area) {
        area = a;
      };
    }
  }
  return area;
}

/**
 * 上面一种方法，额外多算了很多。
 * 由于面积的计算为
 * 长*高，因此我们在选取两个柱子时，可以从距离最远的两个开始，也就是说，不断的缩小“长度”这个变量
 * 由于长度缩小，那么在判断高度时，如果高度比前一个高度还小，则可以直接略过。只计算比较长度缩小，但是高度增大的case。
 *
 */
int maxArea2(int* height, int heightSize) {
  int area = 0;
  int start = 0, end = 0, end_pre = 0;
  int h = 0, l = 0, a = 0;
  for (int i = 0; i < heightSize - 1; i++) {
    start = height[i];
    end_pre = 0;
    for (int j = heightSize - 1; j > i; j--) {
      end = height[j];
      if (end <= end_pre) {
        continue;
      } else {
        h = start > end ? end : start;
        l = j - i;
        a = h * l;
        if (a > area) {
          area = a;
        }
        end_pre = height[j];
      }
    }
  }
  return area;
}

/**
 * 上一个方案中，是固定一个柱子，即i，然后再不断和将i和其他进行组合，虽然中间略过了一些计算，但效率提升还是不明显。
 * 既然可以固定一个柱子，那么如果同时固定两个柱子呢？
 * 怎么同时固定两个柱子呢？
 * 可以使用两个参数标识起始位置。start标识开始位置，end标识结束位置。
 * 然后逐一缩减start与end之间的距离，计算并记录此过程中最大面积
 * 这样只需要O(n)复杂度即可
 * 好吧，其实这个方法我也是看了解析才弄明白的，还是得多多学习呀
 */
int maxArea3(int* height, int heightSize) {
  int area = 0;
  int start = 0, end = heightSize - 1;
  int start_pre = 0, end_pre = 0;
  int hstart = 0, hend = 0;
  int start_or_end = 0;
  while (start < end) {
    hstart = height[start];
    hend = height[end];
    if (start_or_end == -1 && start_pre > hstart) {
      start_pre = hstart;
      start += 1;
      continue;
    } else if (start_or_end == 1 && end_pre > hend) {
      end_pre = hend;
      end -= 1;
      continue;
    }
    int h = hstart > hend ? hend : hstart;
    int l = end - start;
    int a = h * l;
    if (a > area) {
      area = a;
    }
    if (hstart < hend) {
      start_pre = hstart;
      start += 1;
      start_or_end = -1;
    } else {
      end_pre = hend;
      end -= 1;
      start_or_end = 1;
    }
  }
  return area;
}

#define LEN 20
int main(int argc, char const* argv[]) {
  int height[LEN] = {3,  5,  2,  6,  7,  10, 1,  5,  8,  9,
                     11, 23, 16, 38, 94, 1,  28, 64, 38, 43};
  int a1 = maxArea(height, LEN);
  int a2 = maxArea3(height, LEN);
  printf("a1: %d\n", a1);
  printf("a2: %d\n", a2);

  return 0;
}
