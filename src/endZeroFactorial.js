/**
 * 输入一个正整数n，求n!(即阶乘)末尾有多少个0？ 比如: n = 10; n! = 3628800,所以答案为2
 * 输入描述:
 * 输入为一行，n(1 ≤ n ≤ 1000)
 * 输出描述:
 * 输出一个整数,即题目所求
 * 输入例子:
 * 10
 * 输出例子:
 * 2
 */

 const endZeroFactorial = function (n) {
    return n/5 >>>0;
 }

 console.log(endZeroFactorial(1000));