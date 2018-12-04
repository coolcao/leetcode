'use strict';
/**
Given an integer, write an algorithm to convert it to hexadecimal. For negative integer, two’s complement method is used.

Note:

All letters in hexadecimal (a-f) must be in lowercase.
The hexadecimal string must not contain extra leading 0s. If the number is zero, it is represented by a single zero character '0'; otherwise, the first character in the hexadecimal string will not be the zero character.
The given number is guaranteed to fit within the range of a 32-bit signed integer.
You must not use any method provided by the library which converts/formats the number to hex directly.
Example 1:

Input:
26

Output:
"1a"
Example 2:

Input:
-1

Output:
"ffffffff"
 */

/**
 * @param {number} num
 * @return {string}
 */
const toHex = function(num) {
    const num2hex = function(num) {
        const table = {
            '10': 'a',
            '11': 'b',
            '12': 'c',
            '13': 'd',
            '14': 'e',
            '15': 'f',
        };
        const hex = [];
        while (num > 0) {
            const mod = num % 16;
            if (mod>9) {
                hex.push(table[mod]);
            } else {
                hex.push(`${mod}`);
            }
            num = (num/16) >>> 0;
        }
        console.log(hex);
        return hex.reduceRight((pre,cur) => {
            pre += cur;
            return pre;
        }, '');
    }
    if (num > 0) {
        return num2hex(num);
    } else if (num < 0) {
        return num2hex(Math.pow(2,32) + num);
    } else {
        return '0';
    }
};

console.log(toHex(26));