package main

import "fmt"

//题目24:反转链表
// 定义一个函数，输入一个链表的头节点，反转该链表并输出反转后链表的头节点
//输入: 1->2->3->4->5->NULL
//输出: 5->4->3->2->1->NULL

//解法1:遍历的过程中再创建一个新链表
//需要O(n)的时间和O(n)的空间
func reverseList(head *ListNode) *ListNode {
	if head == nil {
		return nil
	}

	var newHead *ListNode
	for p := head; p != nil; p = p.Next {
		newNode := &ListNode{p.Val, newHead}
		newHead = newNode
	}

	return newHead
}

//解法2:原地改变链表的指向
//利用三指针的想法很值得借鉴
func reverseList2(head *ListNode) *ListNode {
	if head == nil {
		return nil
	}

	var newHead *ListNode  //新链表的头
	var p = head           //要被改变链表指向的结点
	var prev *ListNode     //p的前一个结点
	var next *ListNode     //p的后一个

	for {
		if p == nil {
			break
		}

		if p.Next == nil {  //源链表的最后一个结点就是新链表的头
			newHead = p
		}

		next = p.Next
		p.Next = prev
		prev = p
		p = next
	}

	return newHead
}

func printLinkList(head *ListNode) {
	fmt.Println("-----当前链表情况-----")
	for p := head; p != nil; p = p.Next {
		fmt.Printf("%d->", p.Val)
	}
	fmt.Println()
}

func main() {
	node1 := &ListNode{1, nil}
	node2 := &ListNode{2, nil}
	node3 := &ListNode{3, nil}
	node4 := &ListNode{4, nil}
	node5 := &ListNode{5, nil}
	node1.Next = node2
	node2.Next = node3
	node3.Next = node4
	node4.Next = node5

	//newHead := reverseList(node1)
	newHead2 := reverseList2(node1)
	printLinkList(newHead2)
}
