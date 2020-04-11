/*
 * @lc app=leetcode id=225 lang=golang
 *
 * [225] Implement Stack using Queues
 *
 * https://leetcode.com/problems/implement-stack-using-queues/description/
 *
 * algorithms
 * Easy (43.35%)
 * Likes:    525
 * Dislikes: 530
 * Total Accepted:    165.8K
 * Total Submissions: 382.4K
 * Testcase Example:  '["MyStack","push","push","top","pop","empty"]\n[[],[1],[2],[],[],[]]'
 *
 * Implement the following operations of a stack using queues.
 *
 *
 * push(x) -- Push element x onto stack.
 * pop() -- Removes the element on top of the stack.
 * top() -- Get the top element.
 * empty() -- Return whether the stack is empty.
 *
 *
 * Example:
 *
 *
 * MyStack stack = new MyStack();
 *
 * stack.push(1);
 * stack.push(2);
 * stack.top();   // returns 2
 * stack.pop();   // returns 2
 * stack.empty(); // returns false
 *
 * Notes:
 *
 *
 * You must use only standard operations of a queue -- which means only push to
 * back, peek/pop from front, size, and is empty operations are valid.
 * Depending on your language, queue may not be supported natively. You may
 * simulate a queue by using a list or deque (double-ended queue), as long as
 * you use only standard operations of a queue.
 * You may assume that all operations are valid (for example, no pop or top
 * operations will be called on an empty stack).
 *
 *
 */
package main

import "fmt"

// Queue 定义队列
type Queue struct {
	data []int
}

func (q *Queue) push(x int) {
	q.data = append(q.data, x)
}
func (q *Queue) pop() int {
	item := q.data[0]
	q.data = q.data[1:]
	return item
}
func (q *Queue) peek() int {
	item := q.data[0]
	return item
}
func (q *Queue) isEmpty() bool {
	return q.size() == 0
}
func (q *Queue) size() int {
	return len(q.data)
}

// @lc code=start
type MyStack struct {
	length int
	data   [2]*Queue
}

/** Initialize your data structure here. */
func Constructor() MyStack {
	s := MyStack{
		length: 0,
		data: [2]*Queue{
			new(Queue), new(Queue),
		},
	}
	return s
}

/** Push element x onto stack. */
func (this *MyStack) Push(x int) {
	idx := this.length % 2
	for !this.data[idx].isEmpty() {
		this.data[(this.length+1)%2].push(this.data[idx].pop())
	}
	this.data[idx].push(x)
	this.length++
}

/** Removes the element on top of the stack and returns that element. */
func (this *MyStack) Pop() int {
	idx := (this.length - 1) % 2
	item := this.data[idx].pop()
	idx2 := this.length % 2
	if !this.data[idx2].isEmpty() {
		this.data[idx].push(this.data[this.length%2].pop())
	}
	this.length--
	return item
}

/** Get the top element. */
func (this *MyStack) Top() int {
	idx := (this.length - 1) % 2
	return this.data[idx].peek()
}

/** Returns whether the stack is empty. */
func (this *MyStack) Empty() bool {
	return this.length == 0
}

func main() {
	s := Constructor()
	s.Push(1)
	fmt.Printf("%V\n", s.data)
	fmt.Printf("%V\n", s.Pop())
	fmt.Printf("%v\n", s.Empty())
}

/**
 * Your MyStack object will be instantiated and called as such:
 * obj := Constructor();
 * obj.Push(x);
 * param_2 := obj.Pop();
 * param_3 := obj.Top();
 * param_4 := obj.Empty();
 */
// @lc code=end
