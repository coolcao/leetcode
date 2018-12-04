'use strict'

/*
Given two non-empty binary trees s and t, check whether tree t has exactly the same structure and node values with a subtree of s. A subtree of s is a tree consists of a node in s and all of this node's descendants. The tree s could also be considered as a subtree of itself.

Example 1:
Given tree s:

     3
    / \
   4   5
  / \
 1   2
Given tree t:
   4
  / \
 1   2
Return true, because t has the same structure and node values with a subtree of s.
Example 2:
Given tree s:

     3
    / \
   4   5
  / \
 1   2
    /
   0
Given tree t:
   4
  / \
 1   2
Return false.

这个问题很有意思，我觉得我下面的代码很low，可是提交后，在js代码里，排名竟然靠前，不可思议。。。

*/

/**
 * Definition for a binary tree node.
 * function TreeNode(val) {
 *     this.val = val;
 *     this.left = this.right = null;
 * }
 */
/**
 * @param {TreeNode} s
 * @param {TreeNode} t
 * @return {boolean}
 */
const isSubtree = function(s, t) {
  const getSubTrees = function(root) {
    const subs = [];
    const q = [root];
    while (q.length > 0) {
      const node = q.shift();
      subs.push(node);
      if (node.left) q.push(node.left);
      if (node.right) q.push(node.right)
    }
    return subs;
  }
  const compare = function(s, t) {
    if (s && t) {
      if (s.val !== t.val) return false;
        return compare(s.left, t.left) && compare(s.right, t.right)
      return true;
    } else if (s === null && t === null) {
      return true;
    } else {
      return false;
    }
  }

  for(const sub of getSubTrees(s)) {
    if (compare(sub, t)) return true;
  }
  return false;
};

const utils = require('../utils/utils');
// const s = utils.createTree([29,28,30,27,29,29,31,26,26,28,28,28,28,30,32,25,25,25,25,27,29,null,29,29,29,null,29,29,29,31,null,26,24,26,26,26,null,24,null,null,26,28,null,28,28,30,28,30,30,null,null,30,30,30,30,null,32,27,27,null,25,25,null,null,25,27,27,null,null,null,null,27,27,27,29,null,null,null,31,29,null,31,null,29,29,null,null,29,31,null,29,29,31,null,31,null,null,null,28,24,24,24,26,24,24,null,28,26,28,26,null,null,null,28,28,null,28,null,null,28,30,32,null,30,28,28,28,null,null,null,null,28,30,28,28,null,null,null,null,27,null,null,null,23,25,null,null,null,null,null,null,null,null,27,27,null,null,null,29,null,null,null,null,27,29,null,27,27,null,null,null,null,31,29,29,27,29,null,29,27,29,null,null,null,null,27,null,null,29,null,null,22,22,null,26,null,null,26,28,28,28,null,28,null,28,null,28,null,null,null,null,null,null,null,28,null,28,28,null,30,null,null,null,null,null,26,null,28,30,21,23,null,null,null,25,null,27,null,null,null,null,27,29,27,29,27,27,null,null,null,null,29,null,27,null,null,null,25,27,null,null,null,null,null,null,null,null,null,null,null,null,null,null,null,null,null,28,null,null,null,null,null,null,null,null,26,null,null,24,null,28,null,null,null,null,null,23]);
// const t = utils.createTree([29]);
const s = utils.createTree([1,2,3]);
const t = utils.createTree([1,2]);

console.log(isSubtree(s, t));


