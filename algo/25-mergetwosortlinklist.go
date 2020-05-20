package main

import "fmt"

//题目25:合并两个排序的链表
/*
  输入两个递增排序的链表，合并这两个链表并使新链表中的节点仍然是递增排序的
  输入：1->2->4, 1->3->4
  输出：1->1->2->3->4->4
 */

 //它这里的思路是:依次比较
func mergeTwoLists(l1 *ListNode, l2 *ListNode) *ListNode {
	if l1 == nil {
		return l2
	}

	if l2 == nil {
		return l1
	}

	var head *ListNode
	if l1.Val <= l2.Val {
		head = l1
		l1 = l1.Next  //往前一步
	} else {
		head = l2
		l2 = l2.Next
	}

	tail := head       //合并的有序链表的末尾

	for {
		if l1 == nil || l2 == nil {
			break
		}

		if l1.Val <= l2.Val {
			tail.Next = l1
			l1 = l1.Next
		} else {
			tail.Next = l2
			l2 = l2.Next
		}
		tail = tail.Next
	}

	if l1 != nil {
		tail.Next = l1
	} else if l2 != nil {
		tail.Next = l2
	}

	return head
}

func main() {
	l1node1 := &ListNode{1, nil}
	l1node2 := &ListNode{2, nil}
	l1node3 := &ListNode{4, nil}
	l1node4 := &ListNode{7, nil}
	l1node5 := &ListNode{9, nil}
	l1node1.Next = l1node2
	l1node2.Next = l1node3
	l1node3.Next = l1node4
	l1node4.Next = l1node5

	l2node1 := &ListNode{1, nil}
	l2node2 := &ListNode{3, nil}
	l2node3 := &ListNode{4, nil}
	l2node1.Next = l2node2
	l2node2.Next = l2node3

	head := mergeTwoLists(l1node1, l2node1)
	for {
		if head == nil {
			break
		}
		fmt.Printf("%d->", head.Val)
		head = head.Next
	}
	fmt.Println()
}