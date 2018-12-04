'use strict'

/*
Write a function that takes an unsigned integer and returns the number of ’1' bits it has (also known as the Hamming weight).

For example, the 32-bit integer ’11' has binary representation 00000000000000000000000000001011, so the function should return 3.

求一个正整数的二进制表示中，有多少个1的位
*/

/**
 * 最简单的就是将正整数改写成二进制字符串，然后遍历字符串，统计1的个数。
 * 简单易理解，但是，整数在转换成二进制时，使用了js的方法，并不优雅，而且效率也不高。
 * @param {number} n - a positive integer
 * @return {number}
 */
const hammingWeight = function(n) {
  const str = n.toString(2);
  let sum = 0;
  for (let i = 0; i < str.length; i++) {
    if (str[i] === '1') {
      sum += 1;
    }
  }
  return sum;
};
/**
 * 题目中已经指定了正整数范围为 32 位整数，那么可以考虑使用位运算。
 * 我们只需要在每位上和1做 与(&) 运算，结果为1，则说明该位为1，加和即可。
 * 注意：题目中的32位整数为无符号整数。js中位运算是有符号整数，因此，需要注意最高位为1的区别。
 */
const hammingWeight2 = function(n) {
  let m = 1;
  let sum = 0;
  while (n > 0) {
    if ((m & n) === 1) sum += 1;
    n >>>= 1;
  }
  return sum;
}

/**
 * 第二个方法中，使用位运算，将最低位至最高位逐一和1做与(&)运算，不管该位是0还是1，都做运算。
 * 其实我们还可以，只将位为1的位做运算，每次逐一消除最低位的1.
 * 计算中，最关键的一个计算是 ： n &= (n-1)
 * 由于二进制运算中借位的关系， n-1 是将n的最低位的1减为0，而最低位1后面的0改为1，再与n 做与运算后，便可将最低位的1消除。
 * 可以随便找个整数，写成二进制，计算看一下即可。
 * 但是但是，需要注意的，还是上面说的一点，js位运算是32位整数，且计算机中存储数字时，是有正负之分，因此，对于位运算，可计算的范围是除了最高位的符号位后的31位。
 * 而题目中给出的是32位数字，是无符号的，所以，在使用 n &= (n-1) 运算时，还是要主要最高位为1的情况，最高位为1时，先加 1，然后将最高位剔除，再计算。
 */
const hammingWeight3 = function(n) {
  let sum = 0;
  if (n > Math.pow(2, 31)) {
    sum += 1;
    n -= Math.pow(2,31);
  }
  while (n > 0) {
    n &= (n - 1);
    sum += 1;
  }
  return sum;
}

console.log(hammingWeight(4294917295));
console.log(hammingWeight2(4294917295));
console.log(hammingWeight3(4294917295));
