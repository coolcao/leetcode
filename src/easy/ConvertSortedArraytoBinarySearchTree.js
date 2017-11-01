'use strict';

/*
Given an array where elements are sorted in ascending order, convert it to a height balanced BST.
将一个已按升序拍好序的数组转换成高度平衡的二叉搜索树
有个疑问，到底什么是 `height balanced BST`?
*/

const TreeNode = require('../utils/TreeNode');
const utils = require('../utils/utils');

/**
 * Definition for a binary tree node.
 * function TreeNode(val) {
 *     this.val = val;
 *     this.left = this.right = null;
 * }
 */
/**
 * @param {number[]} nums
 * @return {TreeNode}
 */
const sortedArrayToBST = function(nums) {
    if (nums.length === 0) {
        return null;
    }
    const nodes = [];
    nums.reduce((pre, current) => {
        pre.push({
            val: current,
            left: null,
            right: null,
        });
        return pre;
    }, nodes);
    while(nodes.length > 1) {
        let current = nodes.shift();
        let next = nodes.shift();
        if (!current.left) {
            next.left = current;
            current = next;
            nodes.unshift(current);
        } else {
            if (!current.right) {
                current.right = next;
                nodes.unshift(current);
            } else {
                next.left = current;
                current = next;
                nodes.unshift(current);
            }
        }
        
    }
    return nodes[0];
};

const sortedArrayToBST2 = function(nums) {
    const length = nums.length;
    if (length === 0) {
        return null;
    }
    if (length === 1) {
        return {
            val: nums[0],
            left: null,
            right: null,
        }
    }
    let current = (length)/2 >>> 0;
    let left = nums.slice(0,current);
    let right = nums.slice(current + 1,length);
    const node = new TreeNode();
    node.val = nums[current];
    node.left = sortedArrayToBST2(left);
    node.right = sortedArrayToBST2(right);
    return node;
}

const sortedArrayToBST3 = function(nums) {
    const length = nums.length;
    if (length === 0) {
        return null;
    }
    if (length === 1) {
        return new TreeNode(nums[0]);
    }
    let current = (length - 1)/2 >>> 0;
    let left = nums.slice(0,current);
    let right = nums.slice(current + 1,length);
    const node = new TreeNode();
    node.val = nums[current];
    node.left = sortedArrayToBST3(left);
    node.right = sortedArrayToBST3(right);
    return node;
}



const nums = [2,4,6,8,10];

const tree = sortedArrayToBST3(nums);

console.log(utils.expandTree(tree));

