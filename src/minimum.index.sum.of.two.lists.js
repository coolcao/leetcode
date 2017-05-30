'use strict';

/**
 * Suppose Andy and Doris want to choose a restaurant for dinner, and they both have a list of favorite restaurants represented by strings.

    You need to help them find out their common interest with the least list index sum. If there is a choice tie between answers, output all of them with no order requirement. You could assume there always exists an answer.

    Example 1:
    Input:
    ["Shogun", "Tapioca Express", "Burger King", "KFC"]
    ["Piatti", "The Grill at Torrey Pines", "Hungry Hunter Steakhouse", "Shogun"]
    Output: ["Shogun"]
    Explanation: The only restaurant they both like is "Shogun".
    Example 2:
    Input:
    ["Shogun", "Tapioca Express", "Burger King", "KFC"]
    ["KFC", "Shogun", "Burger King"]
    Output: ["Shogun"]
    Explanation: The restaurant they both like and have the least index sum is "Shogun" with index sum 1 (0+1).
    Note:
    The length of both lists will be in the range of [1, 1000].
    The length of strings in both lists will be in the range of [1, 30].
    The index is starting from 0 to the list length minus 1.
    No duplicates in both lists.
 */

/**
 * @param {string[]} list1
 * @param {string[]} list2
 * @return {string[]}
 */
const findRestaurant = function(list1, list2) {
    const hash = {};
    const l1 = list1.length;
    const l2 = list2.length;

    for(let i=0;i<l1;i++){
        const item = list1[i];
        hash[item] = {
            val: item,
            repeat: false,
            idx: i
        }
    }

    for(let i=0;i<l2;i++){
        const item = list2[i];
        if(hash[item]){
            hash[item].idx += i;
            hash[item].repeat = true;
        }
    }

    let min = l1 + l2;

    for(let key in hash){
        if(min > hash[key].idx && hash[key].repeat){
            min = hash[key].idx;
        }
    }

    const result = [];

    for(let key in hash){
        if(min === hash[key].idx && hash[key].repeat){
            result.push(hash[key].val);
        }
    }

    return result; 

};

const list1 = ["Shogun","Tapioca Express","Burger King","KFC"];
const list2 = ["KFC","Burger King","Tapioca Express","Shogun"];

console.log(findRestaurant(list1,list2));





