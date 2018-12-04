'use strict'
/*
Given a sorted linked list, delete all duplicates such that each element appear only once.

For example,
Given 1->1->2, return 1->2.
Given 1->1->2->3->3, return 1->2->3.
*/

/**
 * Definition for singly-linked list.
 * function ListNode(val) {
 *     this.val = val;
 *     this.next = null;
 * }
 */
/**
 * @param {ListNode} head
 * @return {ListNode}
 */
const deleteDuplicates = function(head) {
  if (!head) return null;
  let current = head;
  let next = head.next;
  while (current) {
    while (next && next.val == current.val) {
      next = next.next;
      current.next = next;
    }
    current = next;
  }

  return head;
};

const utils = require('../utils/utils');

const head = utils.createSingleLinkedList([1,2,3]);

const result = deleteDuplicates(head);
console.log(utils.expandSingleLinkedList2Array(result));
