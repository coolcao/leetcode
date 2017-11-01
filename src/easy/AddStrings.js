'use strict';

/**
Given two non-negative integers num1 and num2 represented as string, return the sum of num1 and num2.

Note:

The length of both num1 and num2 is < 5100.
Both num1 and num2 contains only digits 0-9.
Both num1 and num2 does not contain any leading zero.
You must not use any built-in BigInteger library or convert the inputs to integer directly.
 */


/**
* @param {string} num1
* @param {string} num2
* @return {string}
*/
const addStrings = function (num1, num2) {
    class BigNumber {
        constructor(str_num) {
            const length = str_num.length;
            this.length = length;
            this.datas = [];
            for (let i = length - 1; i >= 0; i--) {
                this.datas.push(str_num[i] >>> 0);
            }
        }
        add(big_number) {
            const length = big_number.length > this.length ? big_number.length : this.length;
            for (let i = 0; i < length; i++) {
                this.datas[i] = (this.datas[i] >>> 0) + (big_number.datas[i] >>> 0);
                if (this.datas[i] > 9) {
                    this.datas[i + 1] = (this.datas[i + 1] >>> 0) + (this.datas[i] / 10 >>> 0);
                    this.datas[i] = this.datas[i] % 10;
                }
            }
            return this;
        }

        toString() {
            return this.datas.reduceRight((pre, cur) => {
                pre += cur;
                return pre;
            }, '');
        }
    }

    const bn1 = new BigNumber(num1);
    const bn2 = new BigNumber(num2);
    return bn1.add(bn2).toString();
};

const addStrings2 = function (num1, num2) {
    const step = 9;
    class BigNumber {
        constructor(str_num) {
            const num_length = str_num.length;
            let length = 0;
            if (num_length % 9 === 0) {
                length = (str_num.length / 9 >>> 0);
            } else {
                length = (str_num.length / 9 >>> 0) + 1;
            }
            this.length = length;
            this.datas = [];
            for (let i = str_num.length; i > 0; i -= step) {
                const start = i - step > 0 ? (i - step) : 0;
                this.datas.push(str_num.slice(start, i));
            }
        }
        toString() {
            const to_unit_str = function(num) {
                if (num < 10) return `00000000${num}`
                if (num < 100) return `0000000${num}`
                if (num < 1000) return `000000${num}`
                if (num < 10000) return `00000${num}`
                if (num < 100000) return `0000${num}`
                if (num < 1000000) return `000${num}`
                if (num < 10000000) return `00${num}`
                if (num < 100000000) return `0${num}`
                return `${num}`
            }
            let s = '';
            // 这个地方有问题，使用字符串进行计算，不能用算数运算
            for(let i=this.datas.length - 1;i>=0;i--) {
                if (i === this.datas.length - 1) {
                    s += this.datas[i];
                } else {
                    s += to_unit_str(this.datas[i]);
                }
            }
            return s;
        }
        add(big_number) {
            const length = big_number.length > this.length ? big_number.length : this.length;
            for (let i = 0; i < length; i++) {
                this.datas[i] = (this.datas[i] >>> 0) + (big_number.datas[i] >>> 0);
                if (this.datas[i] > 999999999) {
                    this.datas[i + 1] = (this.datas[i + 1] >>> 0) + (this.datas[i] / 1000000000 >>> 0);
                    this.datas[i] = this.datas[i] % 1000000000;
                }
                
            }
            return this;
        }
    }
    const bn1 = new BigNumber(num1);
    const bn2 = new BigNumber(num2);
    return bn1.add(bn2).toString();
}

console.time('1');
const sum_str = addStrings("18582506933037463847746377834267361478236148729136487623519473826147826387462783468237647862452317947627835462349872164872356424537845646462914621934563194581269452423786383762627373666237826437864738546723542653462354867737466372372752", "366213273635438295738267932647238693296473617842561256743924930438369537257623918623493493864591634002345378652618236432597436503476239174523615683247283647823649123647326417853262773654627329703");
console.timeEnd('1');
console.time('2');
const sum_str2 = addStrings2("18582506933037463847746377834267361478236148729136487623519473826147826387462783468237647862452317947627835462349872164872356424537845646462914621934563194581269452423786383762627373666237826437864738546723542653462354867737466372372752", "366213273635438295738267932647238693296473617842561256743924930438369537257623918623493493864591634002345378652618236432597436503476239174523615683247283647823649123647326417853262773654627329703");
console.timeEnd('2');
console.log(sum_str);
console.log(sum_str2);