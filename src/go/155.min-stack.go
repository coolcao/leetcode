/*
 * @lc app=leetcode id=155 lang=golang
 *
 * [155] Min Stack
 */

package main

import "fmt"

// MinStack 最小栈
type MinStack struct {
	data []int
	min  []int
}

// Constructor 构造函数
func Constructor() MinStack {
	stack := new(MinStack)
	return *stack
}

// Push 入栈
func (this *MinStack) Push(x int) {
	this.data = append(this.data, x)
	minLength := len(this.min)
	if minLength == 0 || x <= this.min[minLength-1] {
		this.min = append(this.min, x)
	}
}

// Pop 出栈
func (this *MinStack) Pop() {
	length, minLength := len(this.data), len(this.min)

	value := this.data[length-1]
	this.data = this.data[0 : length-1]

	if value == this.min[minLength-1] {
		this.min = this.min[0 : minLength-1]
	}
}

// Top 栈顶元素
func (this *MinStack) Top() int {
	length := len(this.data)
	if length <= 0 {
		return 0
	}
	return this.data[length-1]
}

// GetMin 获取栈内元素最小值
func (this *MinStack) GetMin() int {
	length, minLength := len(this.data), len(this.min)
	if length <= 0 {
		return 0
	}
	return this.min[minLength-1]
}

func main() {
	stack := Constructor()
	stack.Push(2147483646)
	fmt.Printf("%v\n", stack.data)
	fmt.Printf("%v\n", stack.min)
	stack.Push(2147483646)
	fmt.Printf("%v\n", stack.data)
	fmt.Printf("%v\n", stack.min)
	stack.Push(2147483647)
	fmt.Printf("%v\n", stack.data)
	fmt.Printf("%v\n", stack.min)
	fmt.Printf("%d\n", stack.Top())
	stack.Pop()
	fmt.Printf("%v\n", stack.data)
	fmt.Printf("%v\n", stack.min)
	fmt.Printf("%d\n", stack.GetMin())
	stack.Pop()
	fmt.Printf("%v\n", stack.data)
	fmt.Printf("%v\n", stack.min)
	fmt.Printf("%d\n", stack.GetMin())
	stack.Pop()
	fmt.Printf("%v\n", stack.data)
	fmt.Printf("%v\n", stack.min)
	stack.Push(2147483647)
	fmt.Printf("%v\n", stack.data)
	fmt.Printf("%v\n", stack.min)
	fmt.Printf("%d\n", stack.Top())
	fmt.Printf("%d\n", stack.GetMin())
	stack.Push(-2147483648)
	fmt.Printf("%d\n", stack.Top())
	fmt.Printf("%d\n", stack.GetMin())
	stack.Pop()
	fmt.Printf("%d\n", stack.GetMin())

}

/**
 * Your MinStack object will be instantiated and called as such:
 * obj := Constructor();
 * obj.Push(x);
 * obj.Pop();
 * param_3 := obj.Top();
 * param_4 := obj.GetMin();
 */
// @lc code=end
