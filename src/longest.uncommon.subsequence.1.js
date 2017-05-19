'use strict';

const findLUSlength1 = function(a, b) {
    let l1 = a.length;
    let l2 = b.length;

    if(l1<l2){
        let tmp = a;
        a = b;
        b = tmp;
        tmp = l1;
        l1 = l2;
        l2 = tmp;
    }

    for(let i=l1;i>0;i--){
        for(let j=0;j<=l1-i;j++){
            if(!b.includes(a.substr(j,i))){
                return i;
            }
        }
    }
    return -1;
};

const findLUSlength2 = function (a,b) {
    let la = a.length;
    let lb = b.length;
    return a===b?-1:la>lb?la:lb;
}

const a = 'sdkfjlskfjklsjfklsdjfklsjlkfjaslkdfjaklsfjlsd';
const b = 'sdklafjksdlhfjksdahfsljkh';

console.log(findLUSlength1(a,b));
console.log(findLUSlength2(a,b));
