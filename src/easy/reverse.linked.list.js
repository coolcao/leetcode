'use strict';

/**
 * Definition for singly-linked list.
 * function ListNode(val) {
 *     this.val = val;
 *     this.next = null;
 * }
 */
/**
 * @param {} head
 * @return {ListNode}
 */
const reverseList = function(head) {
    if(head == null){
        return null;
    }
    let pre = head;
    let current = head.next;
    while (current) {
        let next = current.next;
        if(pre == head){
            pre.next = null;
        }
        current.next = pre;
        pre = current;
        current = next;
    }
    head = pre;
    return head;
};

const head = {
    val: 1,
    next: {
        val: 2,
        next: {
            val: 3,
            next: {
                val: 4,
                next: {
                    val: 5
                }
            }
        }
    }
};

console.log(JSON.stringify(reverseList(head)));