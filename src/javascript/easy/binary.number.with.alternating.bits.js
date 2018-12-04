'use strict'

/*
Given a positive integer, check whether it has alternating bits: namely, if two adjacent bits will always have different values.

Example 1:
Input: 5
Output: True
Explanation:
The binary representation of 5 is: 101
Example 2:
Input: 7
Output: False
Explanation:
The binary representation of 7 is: 111.
Example 3:
Input: 11
Output: False
Explanation:
The binary representation of 11 is: 1011.
Example 4:
Input: 10
Output: True
Explanation:
The binary representation of 10 is: 1010.
*/

/**
 * @param {number} n
 * @return {boolean}
 * 这个方法使用了位运算，所以，会存在溢出的情况。
 */
const hasAlternatingBits = function(n) {
  let b = 1;
  for (let i = 0; b <= n; i++) {
    b = b << 1;
  }
  b = b >>> 1;
  for (let i = 1; b > 0; i++, b = b >>> 1) {
    const mod = i % 2;
    const r = b & n;
    if (((mod === 1) && (r === b)) || ((mod === 0) && (r === 0))) {
      continue;
    } else {
      return false;
    }
  }
  return true;
};

const hasAlternatingBits2 = function(n) {
  const array = n.toString(2).split('');
  for (let i = 1; i < array.length; i++) {
    if ((i % 2 === 1 && array[i] === '0') || (i % 2 === 0 && array[i] === '1')) {
      continue;
    } else {
      return false;
    }
  }
  return true;
}

const hasAlternatingBits3 = function(n) {
  const array = [];
  while (n > 0) {
    array.unshift(n % 2);
    n = n >>> 1;
  }
  // console.log(array);
  for (let i = 1; i < array.length; i++) {
    if ((i % 2 === 1 && array[i] === 0) || (i % 2 === 0 && array[i] === 1)) {
      continue;
    } else {
      return false;
    }
  }
  return true;
}


// console.log((1431655764).toString(2));
// console.log(hasAlternatingBits2(1431655764));
// console.time('1');
// for(let i=1;i<=100;i++) {
//   if(hasAlternatingBits(i)) {
//     console.log((i).toString(2));
//   }
// }
// console.timeEnd('1');
console.time('2');
for(let i=1;i<=100;i++) {
  if(hasAlternatingBits2(i)) {
    console.log((i).toString(2));
  }
}
console.timeEnd('2');
console.time('3');
for(let i=1;i<=100;i++) {
  if(hasAlternatingBits3(i)) {
    console.log((i).toString(2));
  }
}
console.timeEnd('3');
