package main

//题目27:二叉树的镜像
/*
  请完成一个函数，输入一个二叉树，该函数输出它的镜像。

例如输入：

     4
   /   \
  2     7
 / \   / \
1   3 6   9
镜像输出：

     4
   /   \
  7     2
 / \   / \
9   6 3   1
 */

 //
func mirrorTree(root *TreeNode) *TreeNode {
	if root == nil {
		return nil
	}

	if root.Left == nil && root.Right == nil {
		return root
	}

	//当前子树有左右子树,则交换
	root.Left, root.Right = root.Right, root.Left

	//递归处理其左右子树
	if root.Left != nil {
		mirrorTree(root.Left)
	}

	if root.Right != nil {
		mirrorTree(root.Right)
	}

	return root
}

//func main() {
//	root := &TreeNode{4,nil, nil}
//	node2 := &TreeNode{2, nil, nil}
//	node3 := &TreeNode{7, nil, nil}
//	node4 := &TreeNode{1, nil, nil}
//	node5 := &TreeNode{3, nil, nil}
//	node6 := &TreeNode{6, nil, nil}
//	node7 := &TreeNode{9, nil, nil}
//
//	root.Left = node2
//	root.Right = node3
//
//	node2.Left = node4
//	node2.Right = node5
//
//	node3.Left = node6
//	node3.Right = node7
//	fmt.Println("**********原始树为**********")
//	PreOrderTree(root)
//
//	root = mirrorTree(root)
//	fmt.Println("**********镜像树为**********")
//	PreOrderTree(root)
//
//}
