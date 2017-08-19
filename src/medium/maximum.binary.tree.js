/**
Given an integer array with no duplicates. A maximum tree building on this array is defined as follow:

The root is the maximum number in the array.
The left subtree is the maximum tree constructed from left part subarray divided by the maximum number.
The right subtree is the maximum tree constructed from right part subarray divided by the maximum number.
Construct the maximum tree by the given array and output the root node of this tree.

Example 1:
Input: [3,2,1,6,0,5]
Output: return the tree root node representing the following tree:

      6
    /   \
   3     5
    \    / 
     2  0   
       \
        1
Note:
The size of the given array will be in the range [1,1000].
 */

/**
* Definition for a binary tree node.
* function TreeNode(val) {
*     this.val = val;
*     this.left = this.right = null;
* }
*/

class TreeNode {
    constructor(val) {
        this.val = val;
        this.left = this.right = null;
    }
}

/**
 * @param {number[]} nums
 * @return {TreeNode}
 */
const constructMaximumBinaryTree = function(nums) {
    const findMaxAndIdx = function (nums,start,end) {
        let max = nums[start];
        let idx = start;
        // 寻找最大值
        for (let i = start; i < end; i++) {
            if (max < nums[i]) {
                max = nums[i];
                idx = i;
            }
        }
        return { idx, max }
    }
    const constructNode = function (nums,start,end) {
        if (start >= end) {
            return null;
        }
        const { idx, max } = findMaxAndIdx(nums,start,end);
        const node = new TreeNode(max);
        node.left = constructNode(nums,start,idx);
        node.right = constructNode(nums,idx+1,end);
        return node;
    }

    return constructNode(nums,0,nums.length);
};

const array = [3, 2, 1, 6, 0, 5];
const root = constructMaximumBinaryTree(array);
console.log(JSON.stringify(root,null,2));