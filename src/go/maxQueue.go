/*
请定义一个队列并实现函数 max_value 得到队列里的最大值，要求函数max_value、push_back 和 pop_front 的均摊时间复杂度都是O(1)。

若队列为空，pop_front 和 max_value 需要返回 -1

示例 1：

输入:
["MaxQueue","push_back","push_back","max_value","pop_front","max_value"]
[[],[1],[2],[],[],[]]
输出: [null,null,null,2,1,2]
示例 2：

输入:
["MaxQueue","pop_front","max_value"]
[[],[],[]]
输出: [null,-1,-1]


限制：

1 <= push_back,pop_front,max_value的总操作数 <= 10000
1 <= value <= 10^5

来源：力扣（LeetCode）
链接：https://leetcode-cn.com/problems/dui-lie-de-zui-da-zhi-lcof
著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。
**/
package main

import (
	"fmt"
)

func main() {
	q := Constructor()
	q.Push_back(1)
	q.Push_back(2)
	q.Push_back(4)
	q.Push_back(3)
	fmt.Printf("size: %d\n", q.size)
	fmt.Printf("max: %d\n", q.Max_value())
	q.Pop_front()
	fmt.Printf("size: %d\n", q.size)
	fmt.Printf("max: %d\n", q.Max_value())
	q.Pop_front()
	fmt.Printf("size: %d\n", q.size)
	fmt.Printf("max: %d\n", q.Max_value())
	q.Pop_front()
	fmt.Printf("size: %d\n", q.size)
	fmt.Printf("max: %d\n", q.Max_value())
	q.Pop_front()
	fmt.Printf("size: %d\n", q.size)
	fmt.Printf("max: %d\n", q.Max_value())
}

// QueueNode 队列节点
type QueueNode struct {
	Val  int
	Next *QueueNode
	Pre  *QueueNode
}
type DoubleQueue struct {
	Left  *QueueNode
	Right *QueueNode
	Size  int
}

func (dq *DoubleQueue) LeftPeek() int {
	return dq.Left.Val
}

func (dq *DoubleQueue) LeftPush(num int) {
	node := new(QueueNode)
	node.Val = num
	if dq.Left == nil && dq.Right == nil {
		dq.Left = node
		dq.Right = node
	} else {
		current := dq.Left
		current.Pre = node
		node.Next = current
		dq.Left = node
	}

	dq.Size++
}

func (dq *DoubleQueue) LeftPop() int {
	current := dq.Left
	dq.Left = current.Next
	if dq.Left != nil {
		dq.Left.Pre = nil
	}
	dq.Size--
	if dq.Size == 0 {
		dq.Right = nil
	}
	return current.Val
}
func (dq *DoubleQueue) RightPeek() int {
	return dq.Right.Val
}
func (dq *DoubleQueue) RightPush(num int) {
	node := new(QueueNode)
	node.Val = num
	if dq.Left == nil && dq.Right == nil {
		dq.Left = node
		dq.Right = node
	} else {
		current := dq.Right
		current.Next = node
		node.Pre = current
		dq.Right = node
	}

	dq.Size++
}
func (dq *DoubleQueue) RightPop() int {
	current := dq.Right
	dq.Right = current.Pre
	if dq.Right != nil {
		dq.Right.Next = nil
	}
	dq.Size--
	if dq.Size == 0 {
		dq.Left = nil
	}
	return current.Val
}

// MaxQueue 最大队列，使用双端队列实现，左进右出
type MaxQueue struct {
	data *DoubleQueue
	max  *DoubleQueue
	size int
}

func Constructor() MaxQueue {
	return MaxQueue{
		size: 0,
		data: new(DoubleQueue),
		max:  new(DoubleQueue),
	}
}

func (this *MaxQueue) Max_value() int {
	if this.size == 0 {
		return -1
	}
	return this.max.RightPeek()
}

func (this *MaxQueue) Push_back(value int) {
	if this.size == 0 {
		this.max.LeftPush(value)
		this.data.LeftPush(value)
		this.size++
		return
	}
	for this.max.Size > 0 && this.max.RightPeek() < value {
		this.max.RightPop()
	}
	this.max.LeftPush(value)
	this.data.LeftPush(value)
	this.size++
}

func (this *MaxQueue) Pop_front() int {
	if this.size == 0 {
		return -1
	}
	this.size--
	if this.data.RightPeek() == this.max.RightPeek() {
		this.max.RightPop()
	}
	this.data.RightPop()
	return 0
}

/**
 * Your MaxQueue object will be instantiated and called as such:
 * obj := Constructor();
 * param_1 := obj.Max_value();
 * obj.Push_back(value);
 * param_3 := obj.Pop_front();
 */
