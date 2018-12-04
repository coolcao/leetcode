'use strict'

/*
Find the contiguous subarray within an array (containing at least one number) which has the largest sum.

For example, given the array [-2,1,-3,4,-1,2,1,-5,4],
the contiguous subarray [4,-1,2,1] has the largest sum = 6.

click to show more practice.

More practice:
If you have figured out the O(n) solution, try coding another solution using the divide and conquer approach, which is more subtle.
*/

/**
 * 这个方法使用了两层循环，复杂度 O(n^2)，很显然，太不符合题目要求了。。。
 * @param {number[]} nums
 * @return {number}
 */
const maxSubArray = function(nums) {
  const length = nums.length;
  let max = nums[0];
  for (let i = 0; i < length; i++) {
    let tmp_sum = nums[i];
    for (let j = i; j < length; j++) {
      if (i !== j) {
        tmp_sum += nums[j];
      }
      if (max < tmp_sum) {
        max = tmp_sum;
      }
    }
  }
  return max;
};
/*
Analysis of this problem:
Apparently, this is a optimization problem, which can be usually solved by DP. So when it comes to DP, the first thing for us to figure out is the format of the sub problem(or the state of each sub problem). The format of the sub problem can be helpful when we are trying to come up with the recursive relation.

At first, I think the sub problem should look like: maxSubArray(int A[], int i, int j), which means the maxSubArray for A[i: j]. In this way, our goal is to figure out what maxSubArray(A, 0, A.length - 1) is. However, if we define the format of the sub problem in this way, it's hard to find the connection from the sub problem to the original problem(at least for me). In other words, I can't find a way to divided the original problem into the sub problems and use the solutions of the sub problems to somehow create the solution of the original one.

So I change the format of the sub problem into something like: maxSubArray(int A[], int i), which means the maxSubArray for A[0:i ] which must has A[i] as the end element. Note that now the sub problem's format is less flexible and less powerful than the previous one because there's a limitation that A[i] should be contained in that sequence and we have to keep track of each solution of the sub problem to update the global optimal value. However, now the connect between the sub problem & the original one becomes clearer:

maxSubArray(A, i) = maxSubArray(A, i - 1) > 0 ? maxSubArray(A, i - 1) : 0 + A[i];

将题目抽象拆解成更小的步骤：maxSubArray(int A[], int i) 表示数组 A 从 0-i 中最大的子数组方法，A[i]为最后一个元素。
因此，会得到如下求解公式：
maxSubArray(A, i) = maxSubArray(A, i - 1) > 0 ? maxSubArray(A, i - 1) : 0 + A[i];
要求 maxSubArray(A, i) 即先求 maxSubArray(A, i-1)，如果maxSubArray(A, i-1)为整数，则maxSubArray(A,i)为maxSubArray(A,i+1)加上A[i]即可。
否则，如果maxSubArray(A,i-1)是负数，则maxSubArray(A,i)即为A[i]


*/
const maxSubArray2 = function(nums) {
  const length = nums.length;
  const dp = []; //dp[i] means the maximum subarray ending with A[i];
  dp[0] = nums[0];
  let max = dp[0];

  for (let i = 1; i < length; i++) {
    dp[i] = nums[i] + (dp[i - 1] > 0 ? dp[i - 1] : 0);
    max = Math.max(max, dp[i]);
  }

  return max;
}

/*
this problem was discussed by Jon Bentley (Sep. 1984 Vol. 27 No. 9 Communications of the ACM P885)

the paragraph below was copied from his paper (with a little modifications)

algorithm that operates on arrays: it starts at the left end (element A[1]) and scans through to the right end (element A[n]), keeping track of the maximum sum subvector seen so far. The maximum is initially A[0]. Suppose we've solved the problem for A[1 .. i - 1]; how can we extend that to A[1 .. i]? The maximum
sum in the first I elements is either the maximum sum in the first i - 1 elements (which we'll call MaxSoFar), or it is that of a subvector that ends in position i (which we'll call MaxEndingHere).

MaxEndingHere is either A[i] plus the previous MaxEndingHere, or just A[i], whichever is larger.
*/

const maxSubArray3 = function(nums) {
  let maxSoFar = nums[0], maxEndingHere = nums[0];
  for (let i = 1; i < nums.length; ++i) {
    maxEndingHere = Math.max(maxEndingHere + nums[i], nums[i]);
    maxSoFar = Math.max(maxSoFar, maxEndingHere);
  }
  return maxSoFar;
}

const nums = [-2, 1, -3, 4, -1, 2, 1, -5, 4, 9, 8, 7, 0, -1, 2, 5, -1, 5, 8, -8, -7];
console.log(maxSubArray(nums));
console.log(maxSubArray2(nums));
console.log(maxSubArray3(nums));
