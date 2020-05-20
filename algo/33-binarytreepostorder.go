package main

import "fmt"

//题目33:输入一个整数数组,判断该数组是是不是某二叉搜索树的后序遍历结果
/*
  参考以下这颗二叉搜索树：

     5
    / \
   2   6
  / \
 1   3
示例 1：

输入: [1,6,3,2,5]
输出: false
示例 2：

输入: [1,3,2,6,5]
输出: true

//首先二叉搜索树(二叉查找树,二叉排序树)的定义
若它的左子树不空，则左子树上所有结点的值均小于它的根结点的值；
若它的右子树不空，则右子树上所有结点的值均大于它的根结点的值；
它的左、右子树也分别为二叉排序树。
二叉搜索树作为一种经典的数据结构，它既有链表的快速插入与删除操作的特点，又有数组快速查找的优势；所以应用十分广泛，
例如在文件系统和数据库系统一般会采用这种数据结构进行高效率的排序与检索操作。

 */

func verifyPostorder(postorder []int) bool {
	if postorder == nil || len(postorder) < 1 {
		return false
	}

	postLen := len(postorder)
	root := postorder[postLen-1] //根结点

	indexi := 0
	for i := 0; i < postLen - 1; i++ { //划分左右子树
		if postorder[i] > root {
			indexi = i
			break
		}
	}

	if indexi > 0 {               //如果有右子树,在右子树中查找是否有小于根结点的,有则不符合定义
		for j := indexi; j < postLen-1; j++ {
			if postorder[j] < root {
				return false
			}
		}
	}

	//递归判断左子树,右子树
	left, right := true, true
	if indexi > 0 {  //说明有左子树
	   left = verifyPostorder(postorder[:indexi])
	}

	if indexi < postLen - 1 {
		right = verifyPostorder(postorder[indexi:postLen-1])
	}

	return left && right
}

func main() {
	postOrder := []int{5, 7, 6, 9, 11, 10, 8}
	flag1 := verifyPostorder(postOrder)
	fmt.Printf("postOrder:%v is post order:%t\n", postOrder, flag1)

	postOrder1 := []int{5, 7, 6, 11, 9, 10, 8}
	flag2 := verifyPostorder(postOrder1)
	fmt.Printf("postOrder1:%v is post order:%t\n", postOrder1, flag2)

	postOrder3 := []int{5, 6, 8}
	flag3 := verifyPostorder(postOrder3)
	fmt.Printf("postOrder3:%v is post order:%t\n", postOrder3, flag3)

}
