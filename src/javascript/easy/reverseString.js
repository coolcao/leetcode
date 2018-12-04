'use strict';

/**
 * @param {string} s
 * @return {string}
 */
var reverseString = function(s) {
    let length = s.length;
    let result = '';
    for(let i=length-1;i>=0;i--){
        result = result + s.charAt(i);
    }
    return result;
};

console.log(reverseString('abcdefg'));