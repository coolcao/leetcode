'use strict';

/**
 * Given a string, you need to reverse the order of characters in each word within a sentence while still preserving whitespace and initial word order.
    Example 1:
    Input: "Let's take LeetCode contest"
    Output: "s'teL ekat edoCteeL tsetnoc"
    Note: In the string, each word is separated by single space and there will not be any extra space in the string.
 */

/**
 * @param {string} s
 * @return {string}
 */
const reverseWords1 = function(s) {

    const reverse = (s) => {
        const length = s.length;
        let r = '';
        for(let i=length-1;i>=0;i--){
            r += s.charAt(i);
        }
        return r;
    };

    let tmp = reverse(s);
    const array = tmp.split(' ');
    return array.reverse().join(' ');
};

const reverseWords2 = function(s) {

    const reverse = (s) => {
        const length = s.length;
        let r = '';
        for(let i=length-1;i>=0;i--){
            r += s.charAt(i);
        }
        return r;
    };

    const oa = s.split(' ');
    let result = [];
    for(const item of oa){
        result.push(reverse(item));
    }
    return result.join(' ');

};

let s = 'As a project maintainer, making your repo Commitizen friendly allows you to select pre-existing commit message conventions or to create your own custom commit message convention. When a contributor to your repo uses Commitizen, they will be prompted for the correct fields at commit time.';

console.time('1');
console.log(reverseWords1(s));
console.timeEnd('1');

console.time('2');
console.log(reverseWords2(s));
console.timeEnd('2');
