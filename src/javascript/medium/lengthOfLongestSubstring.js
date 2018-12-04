'use strict';

/**
 *  Given a string, find the length of the longest substring without repeating characters.
 *  Examples:
 *  Given "abcabcbb", the answer is "abc", which the length is 3.
 *  Given "bbbbb", the answer is "b", with the length of 1.
 *  Given "pwwkew", the answer is "wke", with the length of 3. Note that the answer must be a substring, "pwke" is a subsequence and not a substring.
 * ====================
 * ------Medium--------
 * ====================
 */

/**
 * @param {string} s
 * @return {number}
 */
let lengthOfLongestSubstring1 = function (s) {
    let length = s.length;
    if(length < 2){
        return s.length;
    }
    let maxlen = 0;
    let j = 0;
    let map = new Map();
    for (let i=0; i<length; i++) {
        map.clear();
        map.set(s.charAt(i), 1);

        for (j=i+1; j<length; j++) {
            if (map.get(s.charAt(j))) {
                break;
            }
            map.set(s.charAt(j), 1);
        }
        if (j-i > maxlen) {
            maxlen = j-i;
        }
    }

    return maxlen;
}
let lengthOfLongestSubstring2 = function (s) {
    if (s.length < 2) {
        return s.length;
    }
    let start = 0,
        end = 0,
        maxlen = 0,
        len = s.length,
        map = new Map();
    while (end < len) {
        if (map.get(s.charAt(end)) >= start) {
            if (end - start > maxlen) {
                maxlen = end - start;
            }
            start = map.get(s.charAt(end)) + 1;
        }
        map.set(s.charAt(end),end++);
    }
    if (end - start > maxlen) {
        maxlen = end - start;
    }
    return maxlen;
}

let lengthOfLongestSubstring3 = function (s) {
    let length = s.length;
    if (length < 2) {
        return s.length;
    }
    let maxlen = 0;
    let index = 0;
    let lastIndex = 0;
    let dp = [];
    dp[0] = 1;
    for(let i=1;i<length;i++){
        for(let j=i-1;j>=lastIndex;j--){
            if(s.charAt(i) == s.charAt(j)){
                dp[i] = i - j;
                lastIndex = j+1;
                break;
            }else{
                dp[i] = dp[i-1] + 1;
            }
        }
        if(dp[i] > maxlen){
            maxlen = dp[i];
            index = i+1-maxlen;
        }
    }
    return maxlen;
}

let lengthOfLongestSubstring4 = function (s) {
    let length = s.length;
    if (length < 2) {
        return s.length;
    }
    let maxlen = 0;
    let index = 0;
    let lastIndex = 0;
    let dp = [];
    dp[0] = 1;
    let map = new Map();
    map.set(s.charAt(0),0);
    for(let i=1;i<length;i++){
        if(map.get(s.charAt(i)) == null){
            map.set(s.charAt(i),i);
            dp[i] = dp[i-1] + 1;
        }else{
            if(lastIndex <= map.get(s.charAt(i))){
                dp[i] = i-map.get(s.charAt(i));
                lastIndex = map.get(s.charAt(i)) + 1;
                map.delete(s.charAt(i));
                map.set(s.charAt(i),i);
            }else{
                dp[i] = dp[i-1] + 1;
                map.delete(s.charAt(i));
                map.set(s.charAt(i),i);
            }
        }
        if(dp[i] > maxlen){
            maxlen = dp[i];
            index = i+1-maxlen;
        }
    }
    return maxlen;
}


console.log(lengthOfLongestSubstring4("dvdef"));
