package tree

//树的孩子表示法
//首先有一个分析问题的过程
//1.树中结点的度是不同的,如果采用树的度作为结点指针域的个数
//好的地方是结构统一,每个结点都都是这么多指针
//不好的地方是浪费空间
const MAXDegree = 16

type ChildNode1 struct {
	data   byte                   //数据域
	childP [MAXDegree]*ChildNode1 //指向其孩子结点的指针域
}

//2.再改进一步,每个结点都有自己的特定的指针域
//好的地方是不那么浪费空间了
//不好的地方是每个指针域数组长度都不一致
type ChildNode2 struct {
	data   byte          //数据与
	degree uint8         //该结点的度
	childP []*ChildNode2 //指针域
}

//3.孩子表示法
//它的核心思想就是:每个结点记录其孩子结点，其所以孩子结点呢，实际上是一条链表

//1.每个结点描述
type ChildNode struct {
	data       byte           //数据域
	firstChild *ChildLinkNode //指向第一个孩子
}

//2.结点的孩子结点
type ChildLinkNode struct {
	index int            //孩子结点下标
	next  *ChildLinkNode //结点的下一个孩子
}

type ChildTree struct {
	nodes [MAXNUM]ChildNode //所有结点
	root  int               //根结点下标
	num   int               //结点数
}

//再进一步,双亲孩子表示法
//parent child
type PCNode struct {
	data       byte           //数据域
	parent     int            //结点的双亲
	firstChild *ChildLinkNode //孩子链
}

type PCTree struct {
	nodes [MAXNUM]PCNode //所有结点
	root  int            //根结点下标
	num   int            //结点数
}
