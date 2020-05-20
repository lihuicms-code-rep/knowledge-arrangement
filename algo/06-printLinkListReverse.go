package main

//题目6:倒序打印链表
/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
*/

type ListNode struct {
	Val int
	Next *ListNode
}

//解法1:允许使用辅助空间,只要将原来的数据放在数组中反转就可以
func reversePrint(head *ListNode) []int {
	arr := []int{}
	for p := head; p != nil; p = p.Next {
		arr = append(arr, p.Val)
	}

	for i := 0; i < len(arr) / 2; i++ {
		arr[i], arr[len(arr)-1-i] = arr[len(arr)-1-i], arr[i]
	}

	return arr
}

//解法2:可以使用一个栈,先进后出
type Stack struct {
	data []int    //数据
	top  int      //栈顶指针
	num  int      //当前元素
}

func NewStack() *Stack {
	return &Stack{
		data:make([]int, 0),
		top:-1,
		num:0,
	}
}

//入栈
func (s *Stack) push(elem int) {
	s.data = append(s.data, elem)
	s.top++
	s.num++
}

//出栈
func (s *Stack) pop() (bool, int) {
	if s.top == -1 {
		return false, -1
	}

	elem := s.data[s.top]
	s.data = s.data[:s.top]  //长度变小
	s.top--
	return true, elem
}

func reversePrint2(head *ListNode) []int {
	stack := NewStack()
	for p := head; p != nil; p = p.Next {
		stack.push(p.Val)
	}

	arr := []int{}
	for {
		if stack.top == -1 {
			break
		}

		success, elem :=  stack.pop()
		if !success {
			break
		}
		arr = append(arr, elem)
	}

	return arr

}
