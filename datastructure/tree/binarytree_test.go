package tree

import (
	"fmt"
	"testing"
)

func TestPreOrderTraverse(t *testing.T) {
	root := CreateBinaryTree()
	fmt.Println("-----前序遍历------")
	PreOrderTraverse(root)
	fmt.Println()
	fmt.Println("-----前序完成------")
}

func TestInOrderTraverse(t *testing.T) {
	root := CreateBinaryTree()
	fmt.Println("-----中序遍历------")
	InOrderTraverse(root)
	fmt.Println()
	fmt.Println("-----中序完成------")
}

func TestPostOrderTraverse(t *testing.T) {
	root := CreateBinaryTree()
	fmt.Println("-----后序遍历------")
	PostOrderTraverse(root)
	fmt.Println()
	fmt.Println("-----后序完成------")
}

func TestCreateBTByPreOrder(t *testing.T) {
	root := NewBiTNode('A')
	CreateBTByPreOrder(root)
	fmt.Println("前序遍历通过前序生成的二叉树")
	InOrderTraverse(root)
}
