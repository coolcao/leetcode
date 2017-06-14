'use strict';

/**
 * Given a roman numeral, convert it to an integer.
 * Input is guaranteed to be within the range from 1 to 3999.
 */

/**
 * @param {string} s
 * @return {number}
 */
const romanToInt = function(s) {
    const sym = {
        "I": 1,
        "V": 5,
        "X": 10,
        "L": 50,
        "C": 100,
        "D": 500,
        "M": 1000
    };
    let sum = 0;
    const length = s.length;
    for(let i=0;i<length;i++){
        const char = s.charAt(i);
        const charN = s.charAt(i+1);
        if(sym[char] < sym[charN]){
            sum -= sym[char];
        }else{
            sum += sym[char];
        }
    }
    return sum;
};


console.log(romanToInt('MMMCCCXXXIII'));
