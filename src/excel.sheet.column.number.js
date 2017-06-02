'use strict';

/**
 * Given a column title as appear in an Excel sheet, return its corresponding column number.

  For example:

      A -> 1
      B -> 2
      C -> 3
      ...
      Z -> 26
      AA -> 27
      AB -> 28
  Credits:
  Special thanks to @ts for adding this problem and creating all test cases.
 */


/**
 * @param {string} s
 * @return {number}
 */
var titleToNumber = function(s) {
    const obj = {
        'A': 1,
        'B': 2,
        'C': 3,
        'D': 4,
        'E': 5,
        'F': 6,
        'G': 7,
        'H': 8,
        'I': 9,
        'J': 10,
        'K': 11,
        'L': 12,
        'M': 13,
        'N': 14,
        'O': 15,
        'P': 16,
        'Q': 17,
        'R': 18,
        'S': 19,
        'T': 20,
        'U': 21,
        'V': 22,
        'W': 23,
        'X': 24,
        'Y': 25,
        'Z': 26
    };
    const length = s.length;
    const arr = [];
    for (let i = length - 1; i >= 0; i--) {
        arr.push(s.charAt(i));
    }
    console.log(arr);
    let sum = 0;
    for(let i=0;i<length;i++){
        sum += obj[arr[i]] * Math.pow(26,i);
    }
    return sum;
};

console.log(titleToNumber('AB'));
