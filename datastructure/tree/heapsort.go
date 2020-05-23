package tree

//堆排序
//堆排序就是利用堆进行排序
//将待排序的序列改造成一颗大顶推

//堆排序
//1.构造一棵大顶推
//2.此时根结点就是最大值(有序)与第一个元素交换
//3.再调整前面n-1个结点
func HeapSort(arr []int) []int {
	if arr == nil || len(arr) < 0 {
		return nil
	}
	length := len(arr)
	buildMaxHeap(arr, length)
	for i := length - 1; i >=0; i-- {
		arr[i], arr[0] = arr[0], arr[i]
		length--
		adjustHeap(arr, 0, length)
	}

	return arr
}

//构造大顶推
//从最后一个非叶结点开始调整
func buildMaxHeap(arr []int, length int) {
	for i := length / 2 - 1; i >= 0; i-- {
		adjustHeap(arr, i, length)
	}
}


//调整结点i的顺序
func adjustHeap(arr []int, i int, length int) {
	if i < 0 || i >= length {
		return
	}

	left := 2 * i + 1
	right := 2 * i + 2
	largest := i

	if left < length && arr[left] > arr[largest] {
		largest = left
	}

	if right < length && arr[right] > arr[largest] {
		largest = right
	}

	if largest != i {
		arr[i], arr[largest] = arr[largest], arr[i]
		adjustHeap(arr, largest, length)
	}
}

