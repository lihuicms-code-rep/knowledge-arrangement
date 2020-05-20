package main

import (
	"container/heap"
	"container/list"
	"container/ring"
	"fmt"
	"math/rand"
)

//container包中容器的使用
/*
  List数据结构的定义
  type Element struct {
	// Next and previous pointers in the doubly-linked list of elements.
	// To simplify the implementation, internally a list l is implemented
	// as a ring, such that &l.root is both the next element of the last
	// list element (l.Back()) and the previous element of the first list
	// element (l.Front()).
	next, prev *Element

	// The list to which this element belongs.
	list *List

	// The value stored with this element.
	Value interface{}
  }

  type List struct {
	root Element // sentinel list element, only &root, root.prev, and root.next are used
	len  int     // current list length excluding (this) sentinel element
  }

 */

//双向链表
func listDemo() {
	l := list.New() //一个双向链表,就会很方便
	l.PushBack('A')
	l.PushBack('B')
	l.PushBack('C')
	l.PushBack('D')

	l.PushFront('E')

	fmt.Println("从头向尾遍历列表")
	for e := l.Front(); e != nil; e = e.Next() {
		fmt.Printf("%c->", e.Value)
	}
	fmt.Println()

	fmt.Println("从尾向头遍历列表")
	for e := l.Back(); e != nil; e = e.Prev() {
		fmt.Printf("<-%c", e.Value)
	}
	fmt.Println()
}

//堆,可以被视为一棵树的数组对象
//1.堆中某个结点的值总是不大于或者不小于其父节点的值 2.是一颗完全二叉树
//四个方面的内容
//1.堆的初始化 2.堆的插入元素和删除元素 3.堆的排序 4.堆的向上/向下调整
/*
  //源码中关于一个堆类型的定义,匿名集成sort.Interface,所以需要实现Len,Swap,Less方法
  type Interface interface {
	sort.Interface
	Push(x interface{}) // add x as element Len()
	Pop() interface{}   // remove and return element Len() - 1.
  }

  //sort中Interface的定义
  type Interface interface {
	// Len is the number of elements in the collection.
	Len() int
	// Less reports whether the element with
	// index i should sort before the element with index j.
	Less(i, j int) bool
	// Swap swaps the elements with indexes i and j.
	Swap(i, j int)
  }

 */

//1.实现小根堆
type IntHeap []int

//实现soft.Interface接口
func (h IntHeap) Len() int           { return len(h) }
func (h IntHeap) Less(i, j int) bool { return h[i] < h[j] }
func (h IntHeap) Swap(i, j int)      { h[i], h[j] = h[j], h[i] }

//实现heap.Interface接口
//注意这里接收者要用指针类型
func (h *IntHeap) Push(x interface{}) {
	value, ok := x.(int) //注意需要先断言
	if !ok {
		return
	}
	*h = append(*h, value)
}

func (h *IntHeap) Pop() interface{} {
	lenH := len(*h)
	x := (*h)[lenH-1] //最后一个元素
	*h = (*h)[:lenH-1]
	return x
}

//2.可以实现优先级队列
//3.topK问题的思路

func heapDemo() {
	intH := &IntHeap{99, 5, 36, 7, 22, 17, 46, 12, 2, 19, 25, 28, 1, 92}
	heap.Init(intH)
	fmt.Printf("新建的小根堆堆顶为:%d\n", (*intH)[0])
	heap.Push(intH, 0)
	fmt.Printf("此时的小根堆堆顶为:%d\n", (*intH)[0])
	heap.Remove(intH, 0)
	fmt.Printf("此时的小根堆堆顶为:%d\n", (*intH)[0])

	//fix应该是在明确修改了heap中第几个元素
	for i, _ := range *intH {
		if i == 1 {
			(*intH)[i] = 66
			heap.Fix(intH, i)
		}
	}

	for i, v := range *intH {
		fmt.Printf("[%d:%d]->", i, v)
	}

	fmt.Println()
}

//3.ring的使用
/*
  type Ring struct {
	next, prev *Ring
	Value      interface{} // for use by client; untouched by this library
}
 */
func ringDemo() {
	r := ring.New(10)  //环形链表
	for i := 0; i < r.Len(); i++ {
		r.Value = rand.Intn(100)
		r = r.Next()
	}

	v := r.Move(6)
	fmt.Printf("move r:%d\n", v.Value)






}

func main() {
	//listDemo()
	//heapDemo()
	ringDemo()
}
