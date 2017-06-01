'use strict';

/**
 * Given an arbitrary ransom note string and another string containing letters from all the magazines, write a function that will return true if the ransom note can be constructed from the magazines ; otherwise, it will return false.

    Each letter in the magazine string can only be used once in your ransom note.

    Note:
    You may assume that both strings contain only lowercase letters.

    canConstruct("a", "b") -> false
    canConstruct("aa", "ab") -> false
    canConstruct("aa", "aab") -> true
 */

/**
 * @param {string} ransomNote
 * @param {string} magazine
 * @return {boolean}
 */
const canConstruct = function(ransomNote, magazine) {
    const ms = magazine.split('').map(item => {
        return {
            char: item,
            used: false
        }
    });

    const contains = function(char) {
        for (let item of ms) {
            if (char === item.char && item.used === false) {
                item.used = true;
                return true;
            }
        }
        return false;
    }

    let result = true;

    let rl = ransomNote.length;
    for (let i = 0; i < rl; i++) {
        let char = ransomNote.charAt(i);
        if (!contains(char)) {
            return false;
        }
    }

    return result;
};

const canConstruct2 = function (ransomNote,magazine) {
    let rl = ransomNote.length;
    let ml = magazine.length;
    if(rl > ml){
        return false;
    }
    let hash = {};
    for(let i=0;i<rl;i++){
        let char = ransomNote.charAt(i);
        for(let j=0;j<ml;j++){
            let mchar = magazine.charAt(j);
            if(char !== mchar){
                if(j === ml - 1){
                    return false;
                }
                continue;
            }
            if(!hash[char+j]){
                hash[char+j] = true;
                break;
            }else{
                if(j === ml-1){
                    return false;
                }
                continue;
            }
        }
    }
    return true;
}

const r = 'a';
const m = '';

console.log(canConstruct2(r, m));





