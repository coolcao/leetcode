'use strict';

/**
 * 在字符串中找到第一个只出现一次的字符，如输入“abaccdeff”则输出 b
 */

const findOneTimeChar = function (str) {
    let length = str.length;
    let map = new Map();
    for(let i=0;i<length;i++){
        let char = str.charAt(i);
        let ctmp = map.get(char);
        if(ctmp){
            map.set(char,{count:ctmp.count+1,index:i});
        }else{
            map.set(char,{count:1,index:i});
        }
    }

    let items = map.entries();
    let firstIndex = length;
    let char = null;

    for(let [key,value] of items){
        if(value.index < firstIndex && value.count === 1){
            firstIndex = value.index;
            char = key;
        }

    }

    return char;

};

const findOneTimeChar2 = function (str) {
    let tmp = [];
    let length = str.length;
    for(let i=0;i<length;i++){
        let index = str.charCodeAt(i);
        if(tmp[index]){
            tmp[index] = {char:str.charAt(i),count:tmp[index].count + 1,index:i}
        }else{
            tmp[index] = {char:str.charAt(i),count:1,index:i};
        }
    }
    let firstIndex = length;
    let char = null;
    for(let item of tmp){
        if(item){
            if(item.index < firstIndex && item.count === 1){
                firstIndex = item.index;
                char = item.char;
            }
        }
    }
    return char;
}



console.log(findOneTimeChar2('abcca'));