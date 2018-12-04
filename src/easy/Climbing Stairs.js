'use strict'

/*
You are climbing a stair case. It takes n steps to reach to the top.

Each time you can either climb 1 or 2 steps. In how many distinct ways can you climb to the top?

Note: Given n will be a positive integer.
这个题很有意思，采用递归的过程计算。我先是列出前几项，找到规律写的。但是，应该要先弄明白为何采用递归才是。

假设，我们要爬到第 N 阶台阶，那么我们的上一步，只能是在 (N-1) 或 (N-2) 阶这两种可能，因此，我们只需将前面两种情况的所有可能相加即可。
递归的终止条件是，只有一阶时，就一种可能。两阶时，可以一步两阶上去，也可以两次一步上去两种可能。

       |-- 1 (n=1)
f(n) = |-- 2 (n=2)
       |-- f(n-1) + f(n-2) (n>2)
*/

/**
 * @param {number} n
 * @return {number}
 */
const climbStairs = function(n) {
  const climb = function(n, a1, a2) {
    if (n === 1) return a1;
    else return climb(n - 1, a2, a1 + a2);
  }
  return climb(n, 1, 2)
};

const climbStairs2 = function(n) {
  const steps = [0, 1, 2];
  for (let i = 3; i <= n; i++) {
    steps[i] = steps[i - 1] + steps[i - 2];
  }
  console.log(steps);
  return steps[n];
}

const n = 43;

console.log(climbStairs2(n));
