package main

import (
	"fmt"
	"github.com/lihuicms-code-rep/knowledge-arrangement/datastructure/tree"
)

func main() {
	root := tree.NewBiTNode('A')
	tree.CreateBTByPreOrder(root)
	fmt.Println("前序遍历树....")
	tree.PreOrderTraverse(root)
}
