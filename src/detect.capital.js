'use strict';

/**
 * Given a word, you need to judge whether the usage of capitals in it is right or not.

    We define the usage of capitals in a word to be right when one of the following cases holds:

    All letters in this word are capitals, like "USA".
    All letters in this word are not capitals, like "leetcode".
    Only the first letter in this word is capital if it has more than one letter, like "Google".
    Otherwise, we define that this word doesn't use capitals in a right way.
    Example 1:
    Input: "USA"
    Output: True
    Example 2:
    Input: "FlaG"
    Output: False
    Note: The input will be a non-empty word consisting of uppercase and lowercase latin letters.

 */

const detectCapitalUse = function(word) {
    let tmp = word.split('').map(item => (item === item.toUpperCase()) >>> 0);

    let length = tmp.length;

    if (length === 0 || length === 1) {
        return true;
    } else {

        let sum = 0;
        for (let i = 1; i < length; i++) {
            sum += tmp[i];
        }
        if (tmp[0] === 0) {
            return sum === 0;
        } else {
            return sum === 0 || sum === tmp.length - 1;
        }

    }

};

/**
 * 还有一种更简单的方式
 * 大写字母的用法只有三种：全是大写，全是小写，只有首字母是大写
 * 那么好了，我们就直接构造这三种情况，看是否和原词相等，如果相等，那么就是正确的，否则就不正确。
 */

const detectCapitalUse2 = function (word) {
    const upperCase = word.toUpperCase();
    const lowerCase = word.toLowerCase();
    const first = word[0];
    return upperCase === word || lowerCase === word || first.toUpperCase() + word.slice(1,word.length).toLowerCase() === word;
}

console.log(detectCapitalUse2('FlaG'));