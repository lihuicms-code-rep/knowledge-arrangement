package main

//题目36:二叉搜索树与双向链表

//pre:记录当前结点的上一个结点
//head:记录二叉搜索树的最左下角的值(也就是双向链表的头了)
//采用中序遍历
func treeToDoubleLinkLsit(root *TreeNode) *TreeNode {
	if root == nil {
		return nil
	}

	var pre, head *TreeNode = nil, nil
	inOrder(root, pre, head)
	//首尾相连
	pre.Right = head
	head.Left = pre

	return head
}

//中序遍历
func inOrder(root *TreeNode, pre, head *TreeNode) {
	if root == nil {
		return
	}

	inOrder(root.Left, pre, head)
	if pre == nil {
		head = root
		pre = root
	} else {
		pre.Right = root
		root.Left = pre
		pre = root
	}

	inOrder(root.Right, pre, head)

}
