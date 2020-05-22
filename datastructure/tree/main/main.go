package main
import (
	"fmt"
	"github.com/lihuicms-code-rep/knowledge-arrangement/datastructure/tree"
)

func main() {
	heapSort()
}

func heapSort() {
	arr := []int{0,1,1,2,4,4,1,3,3,2}
	arr = tree.HeapSort(arr)
	fmt.Printf("堆排序后的arr:%v\n", arr)
}
