package main

import "fmt"

//题目7:重建二叉树
////输入某二叉树的前序遍历和中序遍历的结果，请重建出该二叉树。
//// 假设输入的前序遍历和中序遍历的结果中都不含重复的数字。

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }

   前序遍历 preorder = [3,9,20,15,7]
   中序遍历 inorder = [9,3,15,20,7]
    3
   / \
  9  20
    /  \
   15   7

 */

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

//思路1:跟树有关的,先考虑使用递归
func buildTree(preorder []int, inorder []int) *TreeNode {
	if len(preorder) == 0 || len(inorder) == 0 {
		return nil
	}

	rootValue := preorder[0] //根结点的值
	root := &TreeNode{ //创建当前树的根结点
		Val:   rootValue,
		Left:  nil,
		Right: nil,
	}

	if len(preorder) == 1 { //只有一个结点时返回
		return root
	}

	rootIndexInOrder := -1 //根结点在中序序列中的下标
	for i, v := range inorder {
		if v == rootValue {
			rootIndexInOrder = i
			break
		}
	}

	//创建左子树
	if rootIndexInOrder > 0 {
		leftPreOrder := preorder[1 : rootIndexInOrder+1] //左子树的前序
		leftInOrder := inorder[:rootIndexInOrder]        //左子树的中序
		root.Left = buildTree(leftPreOrder, leftInOrder) //递归生成
	}

	//创建右子树
	if len(preorder) - rootIndexInOrder > 0 {
		rightPreOrder := preorder[rootIndexInOrder+1:]
		rightInOrder := inorder[rootIndexInOrder+1:]
		root.Right = buildTree(rightPreOrder, rightInOrder)
	}

	return root
}

//前序遍历二叉树
func PreOrderTree(root *TreeNode) {
	if root == nil {
		return
	}

	fmt.Printf("前序当前结点:%d\n", root.Val)
	PreOrderTree(root.Left)
	PreOrderTree(root.Right)
}

//中序遍历二叉树
func InOrderTree(root *TreeNode) {
	if root == nil {
		return
	}

	InOrderTree(root.Left)
	fmt.Printf("中序当前结点:%d\n", root.Val)
	InOrderTree(root.Right)

}

//func main() {
//	preorder := []int{1, 2, 4, 7, 3, 5, 6, 8}
//	inorder := []int{4, 7, 2, 1, 5, 3, 8, 6}
//	root := buildTree(preorder, inorder)
//	fmt.Println("前序遍历.......")
//	PreOrderTree(root)
//	fmt.Println("中序遍历.......")
//	InOrderTree(root)
//}
