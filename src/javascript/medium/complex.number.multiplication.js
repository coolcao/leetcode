'use strict';

/**
 * Given two strings representing two complex numbers.

    You need to return a string representing their multiplication. Note i2 = -1 according to the definition.

    Example 1:
    Input: "1+1i", "1+1i"
    Output: "0+2i"
    Explanation: (1 + i) * (1 + i) = 1 + i2 + 2 * i = 2i, and you need convert it to the form of 0+2i.
    Example 2:
    Input: "1+-1i", "1+-1i"
    Output: "0+-2i"
    Explanation: (1 - i) * (1 - i) = 1 + i2 - 2 * i = -2i, and you need convert it to the form of 0+-2i.
    Note:

    The input strings will not have extra blank.
    The input strings will be given in the form of a+bi, where the integer a and b will both belong to the range of [-100, 100]. And the output should be also in this form.

 */

class ComplexNumber {
    constructor(real=0,imaginary=0){
        this.real = real;
        this.imaginary = imaginary;
    }
    multi(cn){
        let real = this.real * cn.real - this.imaginary * cn.imaginary;
        let imaginary = this.real * cn.imaginary + this.imaginary * cn.real;
        return new ComplexNumber(real,imaginary);
    }
    toString(){
        return `${this.real}+${this.imaginary}i`;
    }
}

/**
 * @param {string} a
 * @param {string} b
 * @return {string}
 */
const complexNumberMultiply = function(a, b) {
    const aa = a.split('+');
    const ba = b.split('+');

    const cn1 = new ComplexNumber(parseInt(aa[0]),parseInt(aa[1]));
    const cn2 = new ComplexNumber(parseInt(ba[0]),parseInt(ba[1]));

    return cn1.multi(cn2).toString();

};

console.log(complexNumberMultiply('1+-1i','1+1i'));