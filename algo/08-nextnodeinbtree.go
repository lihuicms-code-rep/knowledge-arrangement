package main

import "fmt"

//题目8:二叉树的下一个结点
/*
  中序遍历的一些特点
  1)如果结点有右子树,那么该结点的下一个结点就是其右子树的最左结点
  2)如果结点没有右子树,并且这个结点是它父亲结点的左子树,那么该结点的下一个结点就是其父节点
  3)最复杂的一种情况,该结点既没有右子树,并且还是其父节点的右子结点,就沿着父指针向上遍历,直到找到
    一个是它父节点的左子结点的结点

  举例:
          a
        /   \
       b     c
      / \   / \
     d   e f   g
        / \
        h  i
  其中序遍历结果为:dbheiafcg
 */

type BinaryTreeNode struct {
	Data   string            //数据域
	Left   *BinaryTreeNode //左子树
	Right  *BinaryTreeNode //右子树
	Parent *BinaryTreeNode //其父节点
}

func NewBNode(data string) *BinaryTreeNode {
	return &BinaryTreeNode{
		Data:data,
		Left:nil,
		Right:nil,
		Parent:nil,
	}
}

//找出中序遍历的下一个结点并返回
func GetNextNode(T *BinaryTreeNode) *BinaryTreeNode {
	//1.结点有右子树
	if T.Right != nil {
		p := T.Right
		for ; p.Left != nil; p = p.Left {
		}
		return p
	}

	//2.结点没有右子树,并且该结点是其父亲结点的左子树
	if T.Right == nil && T.Parent.Left == T {
		return T.Parent
	}

	//3.结点没有右子树,并且还是其父结点的右子树
	if T.Right == nil && T.Parent.Right == T {
		p := T.Parent
		//沿着父亲往上找
		for ; p != nil; p = p.Parent {
			if p.Parent != nil && p == p.Parent.Left {
				break
			}
		}

		return p
	}

	return nil
}

func main() {
	//创建一颗二叉树测试
	root := NewBNode("A")
	BNode := NewBNode("B")
	CNode := NewBNode("C")
	DNode := NewBNode("D")
	ENode := NewBNode("E")
	FNode := NewBNode("F")
	GNode := NewBNode("G")
	HNode := NewBNode("H")
	INode := NewBNode("I")

	root.Left = BNode
	root.Right = CNode
	BNode.Parent, CNode.Parent = root, root

	BNode.Left = DNode
	BNode.Right = ENode
	DNode.Parent, ENode.Parent = BNode, BNode

	CNode.Left = FNode
	CNode.Right = GNode
	FNode.Parent, GNode.Parent = CNode, CNode

	ENode.Left = HNode
	ENode.Right = INode
	HNode.Parent, INode.Parent = ENode, ENode

	get1 := GetNextNode(root)
	fmt.Printf("root node next node is:%+v\n", get1)

	get2 := GetNextNode(BNode)
	fmt.Printf("B node next node is:%+v\n", get2)

	get3 := GetNextNode(CNode)
	fmt.Printf("C node next node is:%+v\n", get3)

	get4 := GetNextNode(DNode)
	fmt.Printf("D node next node is:%+v\n", get4)

	get5 := GetNextNode(ENode)
	fmt.Printf("E node next node is:%+v\n", get5)

	get6 := GetNextNode(FNode)
	fmt.Printf("F node next node is:%+v\n", get6)

	get7 := GetNextNode(GNode)
	fmt.Printf("G node next node is:%+v\n", get7)

	get8 := GetNextNode(HNode)
	fmt.Printf("H node next node is:%+v\n", get8)

	get9 := GetNextNode(INode)
	fmt.Printf("I node next node is:%+v\n", get9)
}
