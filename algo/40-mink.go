package main

import (
	"container/heap"
	"fmt"
	"sort"
)

//题目40:最小的k个数
//1.排序(快排)后取出前k个
func getLeastNumbers(arr []int, k int) []int {
	if arr == nil {
		return nil
	}

	sort.Ints(arr)
	return arr[:k]
}

//2.堆排序
/*
比较直观的想法是使用堆数据结构来辅助得到最小的 k 个数。
堆的性质是每次可以找出最大或最小的元素。
我们可以使用一个大小为 k 的最大堆（大顶堆），将数组中的元素依次入堆，当堆的大小超过 k 时，便将多出的元素从堆顶弹出
 */

//go中堆的使用
/*
  需要实现的接口
  type Interface interface {
   sort.Interface       //sort也是具体接口
   Push(x interface{}) // add x as element Len()
   Pop() interface{}   // remove and return element Len() - 1.
   }


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
type maxHeap []int
//实现
func (mh maxHeap) Len() int           { return len(mh) }
func (mh maxHeap) Less(i, j int) bool { return mh[i] > mh[j] }
func (mh maxHeap) Swap(i, j int)      { mh[i], mh[j] = mh[j], mh[i] }
func (mh *maxHeap) Push(x interface{}) {
	value, ok := x.(int)
	if !ok {
		return
	}
	*mh = append(*mh, value)
}

func (mh *maxHeap) Pop() interface{} {
	old := *mh
	h := len(old)
	x := old[h-1]
	*mh = old[:h-1]
	return x
}

func (mh *maxHeap) Size() int {
	return len(*mh)
}

func (mh *maxHeap) Peek() interface{} {
	if *mh != nil && len(*mh) > 0 {
		return (*mh)[0]
	}
	return nil
}

func getLeastNumbers2(arr []int, k int) []int {
	if arr == nil {
		return nil
	}

	mh := maxHeap{}
	heap.Init(&mh)
	for _, v := range arr {
		if mh.Size() < k {
			heap.Push(&mh, v)
		} else {
			peek := mh.Peek()
			if peek == nil {
				break
			}
			value, ok := peek.(int)
			if !ok {
				break
			}
			if v < value {
				heap.Pop(&mh)
				heap.Push(&mh, v)
			}
		}
	}

	res := make([]int, 0, k)
	for mh.Size() > 0 {
		value, ok := mh.Pop().(int)
		if !ok {
			break
		}
		res = append(res, value)
	}

	return res
}

func main() {
	arr := []int{5, 4, 1, 3, 6, 2, 9}
	res := getLeastNumbers2(arr, 3)
	fmt.Printf("res:%v\n", res)
}



