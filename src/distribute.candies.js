'use strict';

/**
 * Given an integer array with even length, where different numbers in this array represent different kinds of candies. Each number means one candy of the corresponding kind. You need to distribute these candies equally in number to brother and sister. Return the maximum number of kinds of candies the sister could gain.

    Example 1:
    Input: candies = [1,1,2,2,3,3]
    Output: 3
    Explanation:
    There are three different kinds of candies (1, 2 and 3), and two candies for each kind.
    Optimal distribution: The sister has candies [1,2,3] and the brother has candies [1,2,3], too. 
    The sister has three different kinds of candies. 
    Example 2:
    Input: candies = [1,1,2,3]
    Output: 2
    Explanation: For example, the sister has candies [2,3] and the brother has candies [1,1]. 
    The sister has two different kinds of candies, the brother has only one kind of candies. 
    Note:

    The length of the given array is in range [2, 10,000], and will be even.
    The number in given array is in range [-100,000, 100,000].
 */

/**
 * 这个题目有点逻辑题的意思
 * 首先我们先看一下这些糖果总共有多少种
 * 如果糖果种类要比总糖果数量一半还少，那么很明显，将所有不同的糖果分给女生，男生只分同一种相同的糖果即可
 * 如果糖果种类多余总量一半，那么，分给女生的肯定就是糖果总量的一半
 * 这里都不用考虑最后分给男生女生具体的糖果
 */

/**
 * @param {number[]} candies
 * @return {number}
 */
const distributeCandies = function(candies) {
    const length = candies.length;
    const set = new Set(candies);
    const size = set.size;
    const half_length = length/2;
    if(half_length >= set.size){
        return size;
    }else{
        return half_length;
    }
};

console.log(distributeCandies([1,1,2,2,3,3,2,1,8,8]));