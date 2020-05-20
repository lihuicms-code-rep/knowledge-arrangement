package main

import "fmt"

//题目22:链表中倒数第K个结点
/*
 输入一个链表，输出该链表中倒数第k个节点。
为了符合大多数人的习惯，本题从1开始计数，即链表的尾节点是倒数第1个节点。
例如，一个链表有6个节点，从头节点开始，它们的值依次是1、2、3、4、5、6。这个链表的倒数第3个节点是值为4的节点。

给定一个链表: 1->2->3->4->5, 和 k = 2.
返回链表 4->5.
 */
//最直接的解法就是从头遍历
//这种解法需要遍历两次链表,第一次获得链表长度,第二次找到倒数第k个结点
func getKthFromEnd(head *ListNode, k int) *ListNode {
	if head == nil || k < 0 {
		return nil
	}

	len := 0
	for p := head; p != nil; p = p.Next {
		len++
	}

	//倒数第k个结点就往前走len-k+1步
	p := head
	for i := 1; i < len - k + 1; i++ {
		p = p.Next
	}


	return p
}

//只遍历一遍的做法:前后指针
func getKthFromEnd2(head *ListNode, k int) *ListNode {
	if head == nil || k < 0 {
		return nil
	}

	p1 := head       //p1先往前走k步
	p2 := head
	i := 1
	for {
		if p1 == nil {
			break
		}
		p1 = p1.Next
		if i > k {   //p1先往前走k步,然后p2才走,p1走到末尾,p2就走到了len-k+1
			p2 = p2.Next
		}
		i++
	}

	return p2
}

func main() {
	node1 := &ListNode{1,nil}
	node2 := &ListNode{2,nil}
	node3 := &ListNode{3,nil}
	node4 := &ListNode{4,nil}
	node5 := &ListNode{5,nil}
	node1.Next = node2
	node2.Next = node3
	node3.Next = node4
	node4.Next = node5

	kNode1 := getKthFromEnd(node1, 3)
	kNode2 := getKthFromEnd(node1, 5)
	kNode3 := getKthFromEnd(node1, 1)

	kNode4 := getKthFromEnd2(node1, 3)
	kNode5 := getKthFromEnd2(node1, 5)
	kNode6 := getKthFromEnd2(node1, 1)

	fmt.Printf("knode1:%v, knode2:%v,knode3:%v\n", kNode1, kNode2, kNode3)
	fmt.Printf("knode4:%v, knode5:%v,knode6:%v\n", kNode4, kNode5, kNode6)

}
