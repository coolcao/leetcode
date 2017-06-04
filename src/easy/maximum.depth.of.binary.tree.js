'use strict';

/**
 * Given a binary tree, find its maximum depth.

    The maximum depth is the number of nodes along the longest path from the root node down to the farthest leaf node.

    Subscribe to see which companies asked this question.
 */

const utils = require('./utils');

const maxDepth = function(root) {
    if(!root){
        return 0;
    }
    root.deepth = 1;
    const q = [root];
    while (q.length > 0) {

        let node = q.shift();
        if (node.left) {
            node.left.deepth = node.deepth + 1;
            q.push(node.left);
        }
        if (node.right) {
            node.right.deepth = node.deepth + 1;
            q.push(node.right);
        }

        if (q.length === 0) {
            return node.deepth;
        }
    }
}


const root = utils.createTree([1,2,3,null,4,5]);
console.log(root);
console.log(maxDepth(root));