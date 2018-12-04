'use strict';

/**
 * You are given two non-empty linked lists representing two non-negative integers. The digits are stored in reverse order and each of their nodes contain a single digit. Add the two numbers and return it as a linked list.
 * You may assume the two numbers do not contain any leading zero, except the number 0 itself.
 * Input: (2 -> 4 -> 3) + (5 -> 6 -> 4)
 * Output: 7 -> 0 -> 8
 * Subscribe to see which companies asked this question.
 * ====================
 * ------Medium--------
 * ====================
 * 给定两个非空链表，代表两个非负整数。
 * 链表中数字逆序存储，每个节点只有一个数字。
 * 将两个数字整数相加，并以链表的形式返回。
 * 假设两个数字都不存在前导0，当然，除了0本身之外。
 * 本题的难度被标记为“中等”。
 * 其实本题还有一个变种形式，不是使用链表存储，而是采用数组存储。
 * 存储方式和链表类似，也是逆序。
 * 其实这道题难度并不大，只能算简单。
 * 我们只要处理好进位的逻辑即可。
 */

/**
 * Definition for singly-linked list.
 * function ListNode(val) {
 *     this.val = val;
 *     this.next = null;
 * }
 */
/**
 * @param {ListNode} l1
 * @param {ListNode} l2
 * @return {ListNode}
 */

 class ListNode {
    constructor(val){
        this.val = val;
        this.next = null;
    }
 }

var addTwoNumbers = function(l1, l2) {

    let root_node = new ListNode(l1.val + l2.val);

    let current = root_node;

    while (l1.next && l2.next) {
        l1 = l1.next;
        l2 = l2.next;
        let node = new ListNode(l1.val + l2.val);
        current.next = node;
        current = node;
    }

    while (l1.next) {
        let node = new ListNode(l1.next.val);
        current.next = node;
        current = node;
        l1 = l1.next;
    }

    while (l2.next) {
        let node = new ListNode(l2.next.val);
        current.next = node;
        current = node;
        l2 = l2.next;
    }
    current = root_node;
    while(current){
        if(current.val > 9){
            if(current.next){
                current.next.val = current.next.val + ((current.val/10)>>>0);
                current.val = (current.val % 10);
            }else{
                current.next = {};
                current.next.val = (current.val/10)>>>0;
                current.val = (current.val % 10);
            }
        }
        current = current.next;
    }

    return root_node;
};

let l1 = {
    val: 2,
    next: {
        val: 4,
        next: {
            val: 3,
            next: null
        }
    }
};

let l2 = {
    val: 5,
    next: {
        val: 6,
        next:null
    }
};

console.log(JSON.stringify(addTwoNumbers(l1, l2),null,2));
