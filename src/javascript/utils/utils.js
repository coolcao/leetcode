'use strict';

const TreeNode = require('./TreeNode');
const ListNode = require('./ListNode');

/**
 * 由数组构建二叉树
 * @param  {number[]} array
 * @return {TreeNode}       二叉树的root节点
 */
const createTree = function (array) {
    const length = array.length;
    let root = null;
    let current = null;
    const q = [];
    if(length !== 0){
        root = new TreeNode(array[0]);
        q.push(root);

        for(let i=1;i<length;i+=2){
            current = q.shift();

            let left = null;
            let right = null;

            if(array[i]!==null && array[i]!==undefined){
                left =  new TreeNode(array[i]);
            }
            if(array[i+1]!==null && array[i+1]!== undefined){
                right = new TreeNode(array[i+1]);
            }

            current.left = left;
            current.right = right;

            if(left){
                q.push(left);
            }
            if(right){
                q.push(right);
            }
        }

    }
    return root;
};

const expandTree = function(root) {
    const q = [root];
    const array = [];
    while(q.length>0) {
        const node = q.shift();
        array.push(node.val);
        if (node.left)  q.push(node.left);
        if (node.right) q.push(node.right);
    }
    return array;
}

/**
 * 由数组构建单链表
 * @param {number[]} array
 * @return {ListNode} 单链表头
 */
const createSingleLinkedList = function(array) {
  const length = array.length;
  if (length === 0) return null;
  const head = new ListNode(array[0]);
  let current = head;
  for(let i=1;i<length;i++) {
    let node = new ListNode(array[i]);
    current.next = node;
    current = node;
  }
  return head;
}
const expandSingleLinkedList2Array = function(head) {
  let current = head;
  const result = [];
  while(current) {
    result.push(current.val);
    current = current.next;
  }
  return result;
}
module.exports = {
    createTree,
    expandTree,
    createSingleLinkedList,
    expandSingleLinkedList2Array,
}
