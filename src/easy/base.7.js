'use strict';

/**
 * Given an integer, return its base 7 string representation.

    Example 1:
    Input: 100
    Output: "202"
    Example 2:
    Input: -7
    Output: "-10"
    Note: The input will be in range of [-1e7, 1e7].
 */

/**
* @param {number} num
* @return {string}
*/
const convertToBase7 = function(num) {
    if(num === 0){
        return '0';
    }
    let s = '';
    if(num < 0){
        s += '-';
    }
    const tmp = [];
    num = Math.abs(num);
    const mod = num % 7;
    const base = (num / 7) >>> 0;
    while (num / 7 !== 0) {
        tmp.push(num % 7);
        num = (num / 7) >>> 0;
    }
    return tmp.reduceRight((pre,item) => {
        pre += item;
        return pre;
    },s);
};

console.log(convertToBase7(0));
