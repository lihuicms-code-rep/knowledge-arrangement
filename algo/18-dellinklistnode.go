package main

import "fmt"

//题目18:删除链表结点
/*
给定单向链表的头指针和一个要删除的节点的值，定义一个函数删除该节点。返回删除后的链表的头节点。
输入: head = [4,5,1,9], val = 5
输出: [4,1,9]
解释: 给定你链表中值为 5 的第二个节点，那么在调用了你的函数之后，该链表应变为 4 -> 1 -> 9.
 */
 //这个题目如果改一下会有启发
 //要删除的参数给出的结点而不是值,那么可以通过得到结点的下一个元素而进行赋值到上一个然后删除下一个

type LNode struct {
	Val  int
	Next *LNode
}

//给定头指针和要删除的结点的值,删除该结点
func deleteNode(head *LNode, val int) *LNode {
	if head == nil {
		return nil
	}

	p1, p2 := head, head //p1指向要删除的结点,p2指向要删除结点的前一个
	find := false
	for {
		if p1 == nil {
			break
		}

		if p1.Val == val {
			find = true
			break
		}
		p2 = p1
		p1 = p1.Next
	}

	if !find {
		return nil
	}

	//删除
	if p1 == head {
		head = head.Next
		p1.Next = nil
	} else {
		p2.Next = p1.Next
		p1.Next = nil
	}

	return head
}

func printLNode(head *LNode) {
	fmt.Println("-----当前链表情况-----")
	for p := head; p != nil; p = p.Next {
		fmt.Printf("%d->", p.Val)
	}
	fmt.Println()
}

func main() {
	//node1 := LNode{
	//	Val:  4,
	//	Next: nil,
	//}
	//node2 := LNode{
	//	Val:  5,
	//	Next: nil,
	//}
	//node3 := LNode{
	//	Val:  1,
	//	Next: nil,
	//}
	//node4 := LNode{
	//	Val:  9,
	//	Next: nil,
	//}

	node5 := LNode{
		Val:10,
		Next:nil,
	}
	printLNode(&node5)
	head := deleteNode(&node5, 10)
	printLNode(head)

	//node1.Next = &node2
	//node2.Next = &node3
	//node3.Next = &node4
	//printLNode(&node1)
	//
	//head := deleteNode(&node1, 4)
	//printLNode(head)

}
