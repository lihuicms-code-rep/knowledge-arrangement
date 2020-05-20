package main

import "fmt"

//题目30:包含min函数的栈
/*
定义栈的数据结构，请在该类型中实现一个能够得到栈的最小元素的 min 函数在该栈中，调用 min、push 及 pop 的时间复杂度都是 O(1)。
示例:
MinStack minStack = new MinStack();
minStack.push(-2);
minStack.push(0);
minStack.push(-3);
minStack.min();   --> 返回 -3.
minStack.pop();
minStack.top();      --> 返回 0.
minStack.min();   --> 返回 -2.
 */

type MinStack struct {
	data []int //数据域
	min  []int //辅助域
	dTop int   //数据栈栈顶指针
	mTop int   //辅助栈栈顶指针
}

/** initialize your data structure here. */
func ConstructorMS() MinStack {
	return MinStack{
		data: make([]int, 0),
		min:  make([]int, 0),
		dTop: -1,
		mTop: -1,
	}
}

func (this *MinStack) Push(x int) {
	this.dTop++
	this.data = append(this.data, x)

	//同时在每次入栈时更新min的值
	//初次押入的就是最小的值
	if len(this.min) == 0 {
		this.min = append(this.min, x)
	} else {
		//如果当前押入的值小于当前min的栈顶,则押入x
		if x < this.min[this.mTop] {
			this.min = append(this.min, x)
		} else {
			//否则,继续押入当前min的栈顶
			this.min = append(this.min, this.min[this.mTop])
		}
	}
	this.mTop++
}

func (this *MinStack) Pop() {
	if this.IsEmpty() {
		return
	}

	this.data = this.data[:this.dTop]
	this.dTop--
	this.min = this.min[:this.mTop]
	this.mTop--
}

func (this *MinStack) Top() int {
	if this.IsEmpty() {
		return -1
	}

	return this.data[this.dTop]
}

func (this *MinStack) Min() int {
	if this.IsEmpty() {
		return -1
	}
	return this.min[this.mTop]
}

func (this *MinStack) IsEmpty() bool {
	return this.dTop == -1 || this.mTop == -1
}

func main() {
	minStack := ConstructorMS()
	minStack.Push(-2)
	minStack.Push(0)
	minStack.Push(-3)
	fmt.Printf("min1:%d\n", minStack.Min()) // --> 返回 -3.
	minStack.Pop()
	fmt.Printf("top1:%d\n", minStack.Top()) //      --> 返回 0.
	fmt.Printf("min2:%d\n", minStack.Min()) // --> 返回 -2.
}
