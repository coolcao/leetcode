/**
 * Given a non-negative integer num, repeatedly add all its digits until the result has only one digit.
    For example:
    Given num = 38, the process is like: 3 + 8 = 11, 1 + 1 = 2. Since 2 has only one digit, return it.
    Follow up:
    Could you do it without any loop/recursion in O(1) runtime?
    Credits:
    Special thanks to @jianchao.li.fighter for adding this problem and creating all test cases.

    ---------------------------

    给定一个非负整数，将其所有位的数字加和，直至剩下一个数字。
    例如，给定的数字 num=38, 计算过程是： 3+8=11, 1+1=2,因此，最后返回 2
    请遵守如下规则：
    不使用循环和递归，保证时间复杂度在 O(1) 。
 */


/**
 * 这个题目有明确的规定，不能使用循环和递归，且时间复杂度在O(1)内。
 * 那么，肯定不能用题中所示的计算过程。
 * 我们找一下规律：
 * 当数字为1-9时，直接输出。
 * 10 ----> 1
 * 11 ----> 2
 * 12 ----> 3
 * 当数字大于10时，规律和9有关，对9取余。
 * 但是对9取余，范围在 0-8 之间，当余数为0时，其结果对应为9
 * 因此，最后的程序，其实就是对9取余即可。
 */


'use strict';

const addDigits = function (num) {
    if(num < 10){
        return num;
    }
    const mod = num % 9;
    return mod===0?9:mod;
}

console.log(addDigits(18));
