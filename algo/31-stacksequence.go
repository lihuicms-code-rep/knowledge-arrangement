package main

import "fmt"

//题目31:栈的压入,弹出序列
/*
  输入两个整数序列，第一个序列表示栈的压入顺序，请判断第二个序列是否为该栈的弹出顺序。
  假设压入栈的所有数字均不相等。
  例如，序列 {1,2,3,4,5} 是某栈的压栈序列，
  序列 {4,5,3,2,1} 是该压栈序列对应的一个弹出序列，但 {4,3,5,1,2} 就不可能是该压栈序列的弹出序列。

示例 1：
输入：pushed = [1,2,3,4,5], popped = [4,5,3,2,1]
输出：true
解释：我们可以按以下顺序执行：
push(1), push(2), push(3), push(4), pop() -> 4,
push(5), pop() -> 5, pop() -> 3, pop() -> 2, pop() -> 1

示例 2：
输入：pushed = [1,2,3,4,5], popped = [4,3,5,1,2]
输出：false
解释：1 不能在 2 之前弹出。
 */

func validateStackSequences(pushed []int, popped []int) bool {
	if pushed == nil || popped == nil {
		return false
	}

	if len(pushed) != len(popped) {
		return false
	}

	stack := NewCStack()
    index := 0	       //poped的比较下标

	for _, pushItem := range pushed {
		stack.push(pushItem)
		//循环判断当前栈顶元素是否和弹出序列的顺序一致,如果一致,则出栈,同时index后移
		for ; !stack.isEmpty() && stack.getTop() == popped[index]; {
			stack.pop()
			index++
		}
	}

	return stack.isEmpty()
}

func main() {
	pushed := []int{1, 2, 3, 4, 5}
	popped := []int{4, 5, 3, 2, 1}
	popped2 := []int{4, 3, 5, 1, 2}
	f1 := validateStackSequences(pushed, popped)
	fmt.Printf("f1:%t\n", f1)
	f2 := validateStackSequences(pushed, popped2)
	fmt.Printf("f2:%t\n", f2)

}