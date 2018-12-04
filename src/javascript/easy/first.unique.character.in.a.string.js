'use strict';

/**
 * Given a string, find the first non-repeating character in it and return it's index. If it doesn't exist, return -1.

    Examples:

    s = "leetcode"
    return 0.

    s = "loveleetcode",
    return 2.

    Note: You may assume the string contain only lowercase letters.
 */

 /**
  * @param {string} s
  * @return {number}
  */
const firstUniqChar = function(s) {
    const length = s.length;
    const map = new Map();
    for(let i=0;i<length;i++){
        const char = s.charAt(i);
        if(map.has(char)){
            map.set(char,{
                index: i,
                count: map.get(char).count + 1
            });
        }else{
            map.set(char,{
                index: i,
                count: 1
            });
        }
    }
    let result = length;
    const values = map.values();
    for(let val of values){
        if(val.count === 1 && result > val.index){
            result = val.index;
        }
    }
    return result === s.length ? -1 : result;
};

const firstUniqChar2 = function (s) {
    const length = s.length;
    const obj = {};
    for(let i=0;i<length;i++){
        const char = s.charAt(i);
        if(obj[char]){
            obj[char] = {
                index: i,
                count: obj[char].count + 1
            }
        }else{
            obj[char] = {
                index: i,
                count: 1
            }
        }
    }
    let result = length;
    for(let key in obj){
        const val = obj[key];
        if(val.count === 1 && result > val.index){
            result = val.index;
        }
    }
    return result === s.length ? -1 : result;
}

const firstUniqChar3 = function (s) {
    const arr = [];
    const length = s.length;
    for(let i=0;i<length;i++){
        const charCode = s.charCodeAt(i);
        arr[charCode] = (arr[charCode]>>>0) + 1;
    }
    for(let i=0;i<length;i++){
        if(arr[s.charCodeAt(i)] === 1) {
            return i;
        }
    }
    return -1;
}

const s = 'loveleetcode';
console.log(firstUniqChar3(s));
