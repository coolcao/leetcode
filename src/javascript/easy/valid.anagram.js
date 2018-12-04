'use strict'

/**
 * Given two strings s and t, write a function to determine if t is an anagram of s.

    For example,
    s = "anagram", t = "nagaram", return true.
    s = "rat", t = "car", return false.

    Note:
    You may assume the string contains only lowercase alphabets.

    Follow up:
    What if the inputs contain unicode characters? How would you adapt your solution to such case?
 */

/**
* @param {string} s
* @param {string} t
* @return {boolean}
*/
const isAnagram = function(s, t) {
    if (s.length !== t.length) {
        return false;
    }
    const sa = s.split('');
    const ta = t.split('');

    sa.sort();
    ta.sort();

    for(let i=0;i<sa.length;i++){
        if(sa[i] !== ta[i]) {
            return false;
        }
    }

    return true;

};
