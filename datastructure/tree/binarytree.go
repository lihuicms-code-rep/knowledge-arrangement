package tree

import "fmt"

//二叉树顺序存储结构
//完全二叉树可以使用一维数组来存储,因为直接按照层的序号一个个来
//极端情况,一颗深度为K的右斜树,需要分配2^k - 1个结点,实际使用的只有K个,比较浪费空间
//因此顺序存储方式用在完全二叉树上比较好

//二叉链表
//如果有需要还可以增加一个指向其双亲的指针域
//那样就称为三叉链表
type BiTNode struct {
	data   byte    //数据域
	lChild *BiTNode //左右孩子指针
	rChild *BiTNode
}

//相关operation实现
//构造空二叉树/生成一个结点
func NewBiTNode(data uint8) *BiTNode {
	return &BiTNode{
		data:   data,
		lChild: nil,
		rChild: nil,
	}
}

//前序遍历
//按照定义递归方式
func PreOrderTraverse(root *BiTNode) {
	if root == nil { //若二叉树为空,直接返回
		return
	}
	fmt.Printf("%c\t", root.data) //先访问根结点
	PreOrderTraverse(root.lChild) //然后前序遍历左子树
	PreOrderTraverse(root.rChild) //再前序遍历右子树
}

//中序遍历
func InOrderTraverse(root *BiTNode) {
	if root == nil { //若二叉树为空,直接返回
		return
	}

	InOrderTraverse(root.lChild)  //从根结点开始,先中序遍历左子树
	fmt.Printf("%c\t", root.data) //再访问根结点
	InOrderTraverse(root.rChild)  //然后中序遍历右子树
}

//后序遍历
func PostOrderTraverse(root *BiTNode) {
	if root == nil { //若二叉树为空,直接返回
		return
	}

	PostOrderTraverse(root.lChild) //从根结点开始,先后序遍历左子树
	PostOrderTraverse(root.rChild) //然后后序遍历右子树
	fmt.Printf("%c\t", root.data) //再访问根结点
}

//自定义构建一棵二叉树
func CreateBinaryTree() *BiTNode {
	root := NewBiTNode('A')
	BNode := NewBiTNode('B')
	CNode := NewBiTNode('C')
	DNode := NewBiTNode('D')
	ENode := NewBiTNode('E')
	FNode := NewBiTNode('F')
	GNode := NewBiTNode('G')
	HNode := NewBiTNode('H')
	INode := NewBiTNode('I')
	JNode := NewBiTNode('J')
	KNode := NewBiTNode('K')

	//指向关系
	root.lChild = BNode
	root.rChild = CNode

	BNode.lChild = DNode
	BNode.rChild = ENode

	CNode.lChild = FNode
	CNode.rChild = GNode

	DNode.lChild = HNode
	FNode.lChild = INode
	GNode.rChild = JNode
	HNode.rChild = KNode

	return root
}

//按二叉树前序遍历顺序构造一棵二叉树
//此时的前序遍历需要使用扩展二叉树
func CreateBTByPreOrder(t *BiTNode) {
	fmt.Print("请输入结点元素:")
	var data byte
	fmt.Scanln(&data)
	if data == '#' {      //扩展二叉树结点为空标记
		t = nil
	} else {
		t = NewBiTNode(data)
		CreateBTByPreOrder(t.lChild)
		CreateBTByPreOrder(t.rChild)
	}
}