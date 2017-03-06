'use strict';

/**
 * 阶乘
 */

let factorial = function factorial(n) {
    let result = [1,1];
    if(n===0||n===1){
        return result[n];
    }else{
        for(let i=2;i<=n;i++){
            result[i] = i*result[i-1];
        }
        return result[n];
    }
}

console.log(factorial(20));