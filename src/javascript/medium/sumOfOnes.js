'use strict';

/**
 * 计算一个整数的二进制表示中1的位数
 * 比如，4的二进制：100 有1位是1
 */

/**
 * 对于二进制表示，我们将其除以2，如果余数是1，则该位是1，如果余数是0，则该位是0
 * 我们可以不断除以2，通过判断其余数来进行统计1的位数
 * @param  {number} n [整数n]
 * @return {number}   [二进制中1的位数]
 */
const sumOfOnes1 = function (n) {
    let count = 0;
    while (n) {
        if(n % 2 == 1){
            count++;
        }
        n = (n/2)>>>0;
    }
    return count;
}

/**
 * 从二进制表示上，我们可以通过移位进行除以2的运算，而且速度比除法更快
 */
const sumOfOnes2 = function (n) {
    let count = 0;
    while (n) {
        count = count + (n & 1);
        n = n >>> 1;
    }
    return count;
}

/**
 * 第二种方式，很快，但也还不是最优的
 * 因为重复计算了其中为0的位
 * 其实可以直接计算其中为1的位，而跳过为0的位
 * 我们从最低位开始，依次消除其中为1的位即可
 * 使用 n & (n-1) 进行消除
 */
const sumOfOnes3 = function (n) {
    let count = 0;
    while (n) {
        count ++;
        n = n & (n-1);
    }
    return count;
}

console.log(sumOfOnes2(600));
console.log(sumOfOnes3(600));