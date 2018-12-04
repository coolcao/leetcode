'use strict'
/**
Write a program to check whether a given number is an ugly number.

Ugly numbers are positive numbers whose prime factors only include 2, 3, 5. For example, 6, 8 are ugly while 14 is not ugly since it includes another prime factor 7.

Note that 1 is typically treated as an ugly number.

Credits:
Special thanks to @jianchao.li.fighter for adding this problem and creating all test cases.

 */

/**
 * @param {number} num
 * @return {boolean}
 */
const isUgly = function (num) {
    if (num === 1) {
        return true;
    }
    if (num <= 0) {
        return false;
    }
    const getPrime = function (num) {
        const a = [];
        for (let x = 2; x <= num; x++) {
            if (num % x == 0) {
                num /= x;
                a.push(x);
                x -= 1;
            }
        }
        return a;
    }
    const getPrime2 = function (num) {
        let primes = [];
        for (let i = 2; i <= num; i++) {
            while (num !== i) {
                if (num % i === 0) {
                    primes.push(i);
                    num = (num / i) >>> 0;
                } else {
                    break;
                }
            }
        }
        primes.push(num);
        return primes;
    }
    const primes = getPrime2(num);
    for (let item of primes) {
        if (item !== 2 && item !== 3 && item !== 5) {
            return false;
        }
    }
    return true;
};

/**
 * 为啥思维要固定的先去分解质因数呢？分解质因数是一件效率很低的事情。。。
 * @param {number} num 
 * @return {boolean}
 */
const isUgly2 = function(num) {
    if (num <= 0) {
        return false;
    }
    while(num > 1) {
        if (num % 2 === 0) {
            num /= 2;
        } else if (num % 3 === 0) {
            num /= 3;
        } else if (num % 5 === 0) {
            num /= 5;
        } else {
            return false;
        }
    }
    return true;
}

console.log(isUgly2(1));
// console.log(isUgly(14));