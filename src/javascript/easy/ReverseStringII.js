'use strict';

/*
Given a string and an integer k, you need to reverse the first k characters for every 2k characters counting from the start of the string. If there are less than k characters left, reverse all of them. If there are less than 2k but greater than or equal to k characters, then reverse the first k characters and left the other as original.
Example:
Input: s = "abcdefg", k = 2
Output: "bacdfeg"
Restrictions:
The string consists of lower English letters only.
Length of the given string and k will in the range [1, 10000]

*/

/**
 * @param {string} s
 * @param {number} k
 * @return {string}
 */
const reverseStr = function(s, k) {
    const length = s.length;
    let result = '';
    let flag = 0;
    const reverse = function(s) {
        const length = s.length;
        let r = '';
        for(let i=length-1;i>=0;i--) {
            r += s.charAt(i);
        }
        return r;
    }
    for(let i=0;i<length;i+=k) {
        let str = s.substr(i, k);
        if (flag%2==0) {
            console.log(str);
            result += reverse(str);
        } else {
            result += str;
        }
        flag += 1;
    }
    return result;
};

const s = 'abcdefg';

console.log(reverseStr(s, 2));
