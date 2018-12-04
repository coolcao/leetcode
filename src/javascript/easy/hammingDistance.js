'use strict';

/**
 * The Hamming distance between two integers is the number of positions at which the corresponding bits are different.
 * Given two integers x and y, calculate the Hamming distance.
 * Note:
 * 0 ≤ x, y < 2^31.
 * Example:
 * Input: x = 1, y = 4
 * Output: 2
 * Explanation:
 * 1   (0 0 0 1)
 * 4   (0 1 0 0)
 *       ↑   ↑
 * The above arrows point to positions where the corresponding bits are different.
 * ================
 * ------easy------
 * ================
 */

/**
 * @param {number} x
 * @param {number} y
 * @return {number}
 */
const hammingDistance = function(x, y) {
    let tmp = x ^ y;
    let count = 0;
    while (tmp > 0) {
        if(tmp & 1 === 1){
            count ++;
        }
        tmp = tmp >>> 1;
    }
    return count;
};

console.log(hammingDistance(129,897));