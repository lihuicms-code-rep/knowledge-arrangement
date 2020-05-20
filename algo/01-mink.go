package main

import "fmt"

//题目1.最小的k个数
//问题描述:输入n个整数，找出其中最小的K个数。例如输入4,5,1,6,2,7,3,8这8个数字，则最小的4个数字是1,2,3,4,。

//思路1.对n个数进行冒泡排序,依次挑出最小的k个数
//时间复杂度度为O(k*n)
func minK(arr []int, k int) {
	if k < 1 || len(arr) < 0 || k > len(arr) {
		return
	}

    fmt.Print("原数组:", arr)
	kArr := []int{}
	arrLen := len(arr)
	//冒泡排序

	for i := 0; i < k; i++ { //只要进行k轮
		for j := 0; j < arrLen - i - 1; j++ {
			if arr[j] < arr[j+1] { //前面比后面的小,就往后移动
				arr[j], arr[j+1] = arr[j+1], arr[j]
			}
		}
		kArr = append(kArr, arr[arrLen-i-1]) //当前
	}

	fmt.Printf("中挑选出最小的%d个数为:%v\n", k, kArr)
}

//思路2.采用快排的话,时间复杂度为o(nlogn)



func main() {
	arr := []int{4,5,1,6,2,7,3,8}
	k := 1
	minK(arr, k)
}
