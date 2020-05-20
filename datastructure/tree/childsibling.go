package tree

//孩子兄弟表示法
type ChildSiblingNode struct {
	date         byte              //数据域
	firstChild   *ChildSiblingNode //该结点的第一个孩子结点
	rightSibling *ChildSiblingNode //该结点的右兄弟结点
}

//最大的好处就是把一棵复杂的树的结点只用两个指针域就可以描述了
type ChildSiblingTree struct {
	nodes [MAXNUM]ChildSiblingNode
	root int
	num int
}
