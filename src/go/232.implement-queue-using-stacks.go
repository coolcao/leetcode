/*
 * @lc app=leetcode id=232 lang=golang
 *
 * [232] Implement Queue using Stacks
 *
 * https://leetcode.com/problems/implement-queue-using-stacks/description/
 *
 * algorithms
 * Easy (47.67%)
 * Likes:    911
 * Dislikes: 134
 * Total Accepted:    197.2K
 * Total Submissions: 413.3K
 * Testcase Example:  '["MyQueue","push","push","peek","pop","empty"]\n[[],[1],[2],[],[],[]]'
 *
 * Implement the following operations of a queue using stacks.
 *
 *
 * push(x) -- Push element x to the back of queue.
 * pop() -- Removes the element from in front of queue.
 * peek() -- Get the front element.
 * empty() -- Return whether the queue is empty.
 *
 *
 * Example:
 *
 *
 * MyQueue queue = new MyQueue();
 *
 * queue.push(1);
 * queue.push(2);
 * queue.peek();  // returns 1
 * queue.pop();   // returns 1
 * queue.empty(); // returns false
 *
 * Notes:
 *
 *
 * You must use only standard operations of a stack -- which means only push to
 * top, peek/pop from top, size, and is empty operations are valid.
 * Depending on your language, stack may not be supported natively. You may
 * simulate a stack by using a list or deque (double-ended queue), as long as
 * you use only standard operations of a stack.
 * You may assume that all operations are valid (for example, no pop or peek
 * operations will be called on an empty queue).
 *
 *
 */
package main

// Stack 栈
type Stack struct {
	data   []int
	length int
}

func (s *Stack) push(x int) {
	s.data = append(s.data, x)
	s.length++
}
func (s *Stack) peek() int {
	return s.data[s.length-1]
}
func (s *Stack) pop() int {
	if s.length == 0 {
		panic("栈为空")
	}
	item := s.data[s.length-1]
	s.data = s.data[:s.length-1]
	s.length--
	return item
}
func (s *Stack) size() int {
	return s.length
}
func (s *Stack) isEmpty() bool {
	return s.length == 0
}

// @lc code=start
type MyQueue struct {
	data   []*Stack
	length int
}

/** Initialize your data structure here. */
func Constructor() MyQueue {
	return MyQueue{
		data:   []*Stack{new(Stack), new(Stack)},
		length: 0,
	}
}

/** Push element x to the back of queue. */
func (this *MyQueue) Push(x int) {
	this.length++

	for this.data[0].length > 0 {
		this.data[1].push(this.data[0].pop())
	}

	this.data[0].push(x)

	for this.data[1].length > 0 {
		this.data[0].push(this.data[1].pop())
	}

}

/** Removes the element from in front of queue and returns that element. */
func (this *MyQueue) Pop() int {
	item := this.data[0].pop()
	this.length--
	return item

}

/** Get the front element. */
func (this *MyQueue) Peek() int {
	item := this.data[0].peek()
	return item
}

/** Returns whether the queue is empty. */
func (this *MyQueue) Empty() bool {
	return this.length == 0
}

/**
 * Your MyQueue object will be instantiated and called as such:
 * obj := Constructor();
 * obj.Push(x);
 * param_2 := obj.Pop();
 * param_3 := obj.Peek();
 * param_4 := obj.Empty();
 */
// @lc code=end
