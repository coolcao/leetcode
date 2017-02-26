### 题目描述
```
/**
 * You are given two non-empty linked lists representing two non-negative integers. The digits are stored in reverse order and each of their nodes contain a single digit. Add the two numbers and return it as a linked list.
 * You may assume the two numbers do not contain any leading zero, except the number 0 itself.
 * Input: (2 -> 4 -> 3) + (5 -> 6 -> 4)
 * Output: 7 -> 0 -> 8
 * ====================
 * ------Medium--------
 * ====================
 */
```

给定两个非空的链表，表示两个非负整数。
链表中的数字以逆向存储，每个节点只存储一个数字，例如(2->4->3)这个链表表示数字`342`。将两个数字相加，并以链表的形式返回。
假定两个数字都是不包含前导0的数字，除了0本身。
例如：
输入：(2 -> 4 -> 3) + (5 -> 6 -> 4)
输出：7 -> 0 -> 8

### 难易程度：中等

### 解析

这个问题其实是大整数计算问题中的一个简单的环节，加法运算。

题目中给定一种情境，将整数使用链表的形式逆向保存。现在实现加法运算。

其实很简单，由于是逆向存储的，那么，我们不用考虑“对位”的问题，因为第一个永远是个位数。我们只需要将两个数按位加起来，如果每位上的数字大于9，则进行进位处理即可。

### 代码

```js
/**
 * Definition for singly-linked list.
 * function ListNode(val) {
 *     this.val = val;
 *     this.next = null;
 * }
 */
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
```
