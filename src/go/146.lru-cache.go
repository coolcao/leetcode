/*
 * @lc app=leetcode id=146 lang=golang
 *
 * [146] LRU Cache
 *
 * https://leetcode.com/problems/lru-cache/description/
 *
 * algorithms
 * Medium (30.56%)
 * Likes:    6308
 * Dislikes: 267
 * Total Accepted:    585K
 * Total Submissions: 1.8M
 * Testcase Example:  '["LRUCache","put","put","get","put","get","put","get","get","get"]\n' +
  '[[2],[1,1],[2,2],[1],[3,3],[2],[4,4],[1],[3],[4]]'
 *
 * Design and implement a data structure for Least Recently Used (LRU) cache.
 * It should support the following operations: get and put.
 *
 * get(key) - Get the value (will always be positive) of the key if the key
 * exists in the cache, otherwise return -1.
 * put(key, value) - Set or insert the value if the key is not already present.
 * When the cache reached its capacity, it should invalidate the least recently
 * used item before inserting a new item.
 *
 * The cache is initialized with a positive capacity.
 *
 * Follow up:
 * Could you do both operations in O(1) time complexity?
 *
 * Example:
 *
 * LRUCache cache = new LRUCache( 2 );
 *
 * cache.put(1, 1);
 * cache.put(2, 2);
 * cache.get(1);       // returns 1
 * cache.put(3, 3);    // evicts key 2
 * cache.get(2);       // returns -1 (not found)
 * cache.put(4, 4);    // evicts key 1
 * cache.get(1);       // returns -1 (not found)
 * cache.get(3);       // returns 3
 * cache.get(4);       // returns 4
 *
*/
package main

import (
	"fmt"
)

func main() {

}
func printList(n *node) {
	for n != nil {
		fmt.Printf("%v->", n.val)
		n = n.next
	}
	fmt.Println("")
}
func printListFromTail(tail *node) {
	for tail != nil {
		fmt.Printf("<-%v", tail.val)
		tail = tail.pre
	}
	fmt.Println("")

}

// @lc code=start

// 双向链表节点
type node struct {
	key  int
	val  int
	pre  *node
	next *node
}

func moveToTail(head *node, tail *node, n *node) (*node, *node) {
	if head == nil && tail == nil {
		return head, tail
	}
	// 已经是尾节点，不需要再调整
	if n == tail {
		return head, tail
	}

	// 头节点，移动到尾部并调整head指针，tail指针
	if n == head {
		head = head.next
		head.pre = nil

		tail.next = n
		n.pre = tail
		tail = tail.next
		tail.next = nil
		return head, tail
	}

	pre, next := n.pre, n.next
	n.pre, n.next = nil, nil
	pre.next = next
	next.pre = pre

	tail.next = n
	n.pre = tail
	tail = tail.next
	return head, tail

}

func removeHead(head *node, tail *node) (*node, *node) {
	if head == tail {
		return nil, nil
	}
	next := head.next
	head.next = nil
	head = next
	head.pre = nil
	return head, tail
}

func addToTail(head, tail, n *node) (*node, *node) {
	if head == nil && tail == nil {
		head = n
		tail = n
		return head, tail
	}
	tail.next = n
	n.pre = tail
	tail = tail.next
	return head, tail
}

type LRUCache struct {
	capacity int           // 容量
	used     int           // 已使用
	data     map[int]*node // 哈希表数据域
	head     *node         // 链表头节点，标识最先入缓存的节点
	tail     *node         // 链表尾节点，标识最后入缓存的节点
}

func Constructor(capacity int) LRUCache {
	cache := LRUCache{
		capacity: capacity,
		used:     0,
		data:     map[int]*node{},
		head:     nil,
		tail:     nil,
	}
	return cache
}

func (this *LRUCache) Get(key int) int {
	n, ok := this.data[key]
	if !ok {
		return -1
	}

	// 存在节点，调整其在链表中的位置
	this.head, this.tail = moveToTail(this.head, this.tail, n)

	return n.val
}

func (this *LRUCache) Put(key int, value int) {
	n, ok := this.data[key]
	if !ok {
		// 不存在
		// 判断容量是否已满
		n := &node{val: value, key: key}
		if this.capacity == this.used {
			// 已满
			// 删除头节点，并将新节点添加到尾部
			headKey := this.head.key
			fmt.Printf("to delete key: %v\n", headKey)
			delete(this.data, headKey)
			this.head, this.tail = removeHead(this.head, this.tail)
			this.head, this.tail = addToTail(this.head, this.tail, n)
		} else {
			// 未满
			this.head, this.tail = addToTail(this.head, this.tail, n)
			this.used++
		}
		this.data[key] = n
	} else {
		// 存在，更新
		n.val = value
		this.head, this.tail = moveToTail(this.head, this.tail, n)
		this.data[key] = n
	}
}

/**
 * Your LRUCache object will be instantiated and called as such:
 * obj := Constructor(capacity);
 * param_1 := obj.Get(key);
 * obj.Put(key,value);
 */
// @lc code=end
