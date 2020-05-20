package tree

const MAXNUM = 65536

//树的存储结构
//顺序存储
//1.双亲表示法
type parentNode struct {
	data   byte //数据
	parent int  //该结点的双亲结点所在数组中的下标
}

type ParentTree struct {
	nodes [MAXNUM]parentNode //结点域
	root  int                //根结点在数组中的下标
	num   int                //当前结点数
}

//在这个基础的双亲结构上,可以发现可以在O(1)的时间内就找到结点的双亲结点
//但是如果想知道这个结点的孩子,就必须遍历整个数组

//可以在这个基础上进行扩充,比如加上这个结点的最左边孩子
//比如加上这个结点的兄弟结点


