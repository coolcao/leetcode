'use strict'

/**
 * Given a List of words, return the words that can be typed using letters of alphabet on only one row's of American keyboard like the image below.
 * Example 1:
 * Input: ["Hello", "Alaska", "Dad", "Peace"]
 * Output: ["Alaska", "Dad"]
 * Note:
 * You may use one character in the keyboard more than once.
 * You may assume the input string will only contain letters of alphabet.
 */

/**
 * @param {string[]} words
 * @return {string[]}
 */
const findWords = function(words) {
    //取下不取上
    //0-10
    //10-19
    //19-25
    let chars = ['q','w','e','r','t','y','u','i','o','p','a','s','d','f','g','h','j','k','l','z','x','c','v','b','n','m'];
    let map = new Map();
    for(let i=0;i<chars.length;i++){
        map.set(chars[i],i);
    }

    let isOneRow = function (word) {
        word = word.toLowerCase();
        let start = 0;
        let end = 26;
        let length = word.length;
        let index = map.get(word[0]);
        if(index>=0 && index <10){
            start = 0;
            end = 10;
        }else if(index >=10 && index < 19){
            start = 10;
            end = 19;
        }else if(index >= 19 && index < 26){
            start = 19;
            end = 26;
        }
        for(let i=1;i<length;i++){
            index = map.get(word[i]);
            if(!(index >= start && index < end)){
                return false;
            }
        }
        return true;
    }

    return words.filter(isOneRow);

};

console.log(findWords(["Hello","Dad"]));
