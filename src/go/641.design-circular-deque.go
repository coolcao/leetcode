/*
 * @lc app=leetcode id=641 lang=golang
 *
 * [641] Design Circular Deque
 *
 * https://leetcode.com/problems/design-circular-deque/description/
 *
 * algorithms
 * Medium (51.29%)
 * Likes:    197
 * Dislikes: 39
 * Total Accepted:    14.3K
 * Total Submissions: 27.9K
 * Testcase Example:  '["MyCircularDeque","insertLast","insertLast","insertFront","insertFront","getRear","isFull","deleteLast","insertFront","getFront"]\n[[3],[1],[2],[3],[4],[],[],[],[4],[]]'
 *
 * Design your implementation of the circular double-ended queue (deque).
 *
 * Your implementation should support following operations:
 *
 *
 * MyCircularDeque(k): Constructor, set the size of the deque to be k.
 * insertFront(): Adds an item at the front of Deque. Return true if the
 * operation is successful.
 * insertLast(): Adds an item at the rear of Deque. Return true if the
 * operation is successful.
 * deleteFront(): Deletes an item from the front of Deque. Return true if the
 * operation is successful.
 * deleteLast(): Deletes an item from the rear of Deque. Return true if the
 * operation is successful.
 * getFront(): Gets the front item from the Deque. If the deque is empty,
 * return -1.
 * getRear(): Gets the last item from Deque. If the deque is empty, return
 * -1.
 * isEmpty(): Checks whether Deque is empty or not.
 * isFull(): Checks whether Deque is full or not.
 *
 *
 *
 *
 * Example:
 *
 *
 * MyCircularDeque circularDeque = new MycircularDeque(3); // set the size to
 * be 3
 * circularDeque.insertLast(1);            // return true
 * circularDeque.insertLast(2);            // return true
 * circularDeque.insertFront(3);            // return true
 * circularDeque.insertFront(4);            // return false, the queue is full
 * circularDeque.getRear();              // return 2
 * circularDeque.isFull();                // return true
 * circularDeque.deleteLast();            // return true
 * circularDeque.insertFront(4);            // return true
 * circularDeque.getFront();            // return 4
 *
 *
 *
 *
 * Note:
 *
 *
 * All values will be in the range of [0, 1000].
 * The number of operations will be in the range of [1, 1000].
 * Please do not use the built-in Deque library.
 *
 *
 */

package main

// @lc code=start
type MyCircularDequeNode struct {
	Val  int
	Next *MyCircularDequeNode
	Pre  *MyCircularDequeNode
}
type MyCircularDeque struct {
	front *MyCircularDequeNode
	last  *MyCircularDequeNode
	size  int // 大小
	used  int // 已使用
}

/** Initialize your data structure here. Set the size of the deque to be k. */
func Constructor(k int) MyCircularDeque {
	return MyCircularDeque{
		front: nil,
		last:  nil,
		size:  k,
		used:  0,
	}
}

/** Adds an item at the front of Deque. Return true if the operation is successful. */
func (this *MyCircularDeque) InsertFront(value int) bool {
	// 已满，无法插入
	if this.IsFull() {
		return false
	}
	node := new(MyCircularDequeNode)
	node.Val = value
	// 空的，直接把前后指针指到新节点即可
	if this.IsEmpty() {
		this.front = node
		this.last = node
		this.used++
		return true
	}

	// 加入到头节点，调整指针
	current := this.front
	node.Next = current
	current.Pre = node
	this.front = node

	this.used++

	return true
}

/** Adds an item at the rear of Deque. Return true if the operation is successful. */
func (this *MyCircularDeque) InsertLast(value int) bool {
	// 已满，无法插入
	if this.IsFull() {
		return false
	}
	node := new(MyCircularDequeNode)
	node.Val = value
	// 空的，直接把前后指针指到新节点即可
	if this.IsEmpty() {
		this.front = node
		this.last = node
		this.used++
		return true
	}

	// 加入到尾节点，调整指针
	lastNode := this.last
	lastNode.Next = node
	node.Pre = lastNode
	this.last = node

	this.used++

	return true

}

/** Deletes an item from the front of Deque. Return true if the operation is successful. */
func (this *MyCircularDeque) DeleteFront() bool {
	if this.IsEmpty() {
		return false
	}
	frontNode := this.front
	// 只有一个节点
	if this.used == 1 {
		this.front = nil
		this.last = nil
	} else {
		nextNode := frontNode.Next
		nextNode.Pre = nil
		this.front = nextNode
	}
	this.used--
	return true
}

/** Deletes an item from the rear of Deque. Return true if the operation is successful. */
func (this *MyCircularDeque) DeleteLast() bool {
	if this.IsEmpty() {
		return false
	}
	lastNode := this.last
	if this.used == 1 {
		this.front = nil
		this.last = nil
	} else {
		preNode := lastNode.Pre
		preNode.Next = nil
		this.last = preNode
	}
	this.used--
	return true
}

/** Get the front item from the deque. */
func (this *MyCircularDeque) GetFront() int {
	if this.IsEmpty() {
		return -1
	}
	return this.front.Val
}

/** Get the last item from the deque. */
func (this *MyCircularDeque) GetRear() int {
	if this.IsEmpty() {
		return -1
	}
	return this.last.Val
}

/** Checks whether the circular deque is empty or not. */
func (this *MyCircularDeque) IsEmpty() bool {
	if this.used == 0 {
		return true
	}
	return false
}

/** Checks whether the circular deque is full or not. */
func (this *MyCircularDeque) IsFull() bool {
	if this.size == this.used {
		return true
	}
	return false
}

/**
 * Your MyCircularDeque object will be instantiated and called as such:
 * obj := Constructor(k);
 * param_1 := obj.InsertFront(value);
 * param_2 := obj.InsertLast(value);
 * param_3 := obj.DeleteFront();
 * param_4 := obj.DeleteLast();
 * param_5 := obj.GetFront();
 * param_6 := obj.GetRear();
 * param_7 := obj.IsEmpty();
 * param_8 := obj.IsFull();
 */
// @lc code=end
