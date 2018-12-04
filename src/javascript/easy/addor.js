'use strict';

/**
 * 给定x,k，求满足 x + y = x|y 的第k小的正整数。| 是二进制的或（or）运算，例如 3|5=7.
 * 比如当x=5，k=1时返回2，因为5+1=6不等于5|1=5，而5+2=7等于5|2=7.
 * 输入描述：每组测试用例仅包含一组数据，每组数据为两个正整数x,k，满足 0<x,k<=2,000,000,000
 * 输出描述：输出一个数y
 * 输入例子： 5 1
 * 输出例子：2
 */

/**
 * 暴力计算，效率极低，当k的值接近2,000,000,000时，几乎解不出来
 */
const addor = function(x, k) {
    let result = 0;
    let i = 1;
    let count = 0;
    while (count < k) {
        if (x + i == (x | i)) {
            result = i;
            count ++;
        }
        i++;
    }
    return result;
}

/**
 * 该方法使用二进制的原理进行操作
 * 满足 x+y=x|y 的数，只需要将x的二进制中所有为1的位设置为0
 * 剩余的位进行排列组合即可
 * 要求第k大的数，即将k的二进制依次填充到x二进制中除了0剩余的所有位即可
 * 下面这个方法，达到了功能需求，效率也比第一种要快
 * 但是由于在运算过程中，转换字符串，数组逆转，转换为整数等各项操作是非常耗时的
 * 应该还有其他的，直接用位操作进行操作的方法
 */
const addor2 = function(x, k) {
    //将x二进制所有为1的位设置为0
    //其余的位自由组合
    //逆序存储
    let x_array = x.toString(2).split('').reverse();
    let y_array = x_array.map(item => {
        if (item == '1') {
            return '0';
        } else if (item == '0') {
            return undefined;
        }
    });
    let k_array = k.toString(2).split('').reverse();
    // console.log(y_array);
    let idx = 0;
    for(let i=0;i<k_array.length;i++){

        if(y_array[idx] == undefined){
            y_array[idx] = k_array[i];
        }else{
            while (y_array[idx] != undefined) {
                idx ++;
            }
            y_array[idx] = k_array[i];
        }
    }

    let result = 0;
    //y_array拼数
    for(let i=0;i<y_array.length;i++){
        let num = y_array[i] >>> 0;
        result = result + Math.pow(2,i) * num;
    }
    return result;


}

console.time('loop');
console.log(addor(5,150003200));
console.timeEnd('loop');
console.time('time');
console.log(addor2(5,150003200));
console.timeEnd('time');
