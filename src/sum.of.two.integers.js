'use strict';

/**
 * Calculate the sum of two integers a and b, but you are not allowed to use the operator + and -.

    Example:
    Given a = 1 and b = 2, return 3.

    Credits:
    Special thanks to @fujiaozhu for adding this problem and creating all test cases.
 */

const getSum = function(a, b) {
    if (b === 0) {
        return a;
    } else {
        const tmp = a ^ b;
        const carry = (a & b) << 1;
        return getSum(tmp, carry);
    }
};

console.log(getSum(71, -500));