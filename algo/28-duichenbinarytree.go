package main

import "fmt"

//题目28:对称的二叉树
/*
  请实现一个函数，用来判断一棵二叉树是不是对称的。如果一棵二叉树和它的镜像一样，那么它是对称的。

例如，二叉树 [1,2,2,3,4,4,3] 是对称的。

    1
   / \
  2   2
 / \ / \
3  4 4  3
但是下面这个 [1,2,2,null,3,null,3] 则不是镜像对称的:

    1
   / \
  2   2
   \   \
   3    3

 */

 //首先要了解:对称二叉树的定义
 //对树中任意的两个对称结点L和R
 //1)L.Val=R.Val 2)L.Left.Val = R.Right.Val 3)L.Right.Val = R.Left.Val
 //从顶部往下递归,判断每一对结点是否对称即可

func isSymmetric(root *TreeNode) bool {
	if root == nil {
		return true
	}

	return recur(root.Left, root.Right)
}

func recur(l *TreeNode, r *TreeNode) bool {
	if  l == nil && r == nil {
		return true
	}

	if l == nil || r == nil {
		return false
	}

	if l.Val != r.Val {
		return false
	}

	return recur(l.Left, r.Right) && recur(l.Right, r.Left)

}


func main() {
	root := &TreeNode{4, nil, nil}
	node2 := &TreeNode{2, nil, nil}
	node3 := &TreeNode{7, nil, nil}
	node4 := &TreeNode{1, nil, nil}
	node5 := &TreeNode{3, nil, nil}
	node6 := &TreeNode{6, nil, nil}
	node7 := &TreeNode{9, nil, nil}

	root.Left = node2
	root.Right = node3

	node2.Left = node4
	node2.Right = node5

	node3.Left = node6
	node3.Right = node7
	fmt.Println("**********原始树为**********")
	PreOrderTree(root)

	isM := isSymmetric(root)
	fmt.Printf("该数是否为对称的:%t\n", isM)



}
