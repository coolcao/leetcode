```
/**
 * The Hamming distance between two integers is the number of positions at which the corresponding bits are different.
 * Given two integers x and y, calculate the Hamming distance.
 * Note:
 * 0 ≤ x, y < 2^31.
 * Example:
 * Input: x = 1, y = 4
 * Output: 2
 * Explanation:
 * 1   (0 0 0 1)
 * 4   (0 1 0 0)
 *       ↑   ↑
 * The above arrows point to positions where the corresponding bits are different.
 * ================
 * ------easy------
 * ================
 */
 ```

 
### 汉明距离问题
两个整数间的汉明距离是指，两个数二进制位不同的位数。
给定两个整数x,y，计算它们的汉明距离。
**注意：**
0 ≤ x, y < 2^31.
例如，输入 `x = 1, y = 4`，输出`2`。
```
1   (0 0 0 1)
4   (0 1 0 0)
       ↑   ↑
```
上面箭头所指的位置就是1和4两个数的二进制不同的位，总共有2处位置，因此输出2。

### 难易程度：简单

### 解析


这个题目，首先，注意点中明确说明了x,y的取值范围在 0~2^31之间，而且和二进制位有关，因此我们首先考虑到的是位运算，因为位运算只能对32位整数有效。

题目中要统计二进制位不同的位数，位运算中的“异或”运算正好可以处理。

异或运算中，相同的位异或得0，不同的位异或得1。因此我们将两个数进行异或运算，结果二进制中，所有为1的位即是两个数不同的位。

那么，接下来的问题就是，统计这个结果中二进制表示中所有1的位数。

统计一个二进制中1的位数也有很多种方法，这里我们采用逐位与1进行“与”运算进行统计。

### 代码
```js
const hammingDistance = function(x, y) {
    let tmp = x ^ y;
    let count = 0;
    while (tmp > 0) {
        if(tmp & 1 === 1){
            count ++;
        }
        tmp = tmp >>> 1;
    }
    return count;
};
```
