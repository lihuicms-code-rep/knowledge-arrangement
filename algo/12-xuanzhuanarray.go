package main

import "fmt"

//题目12:旋转数组的最小数字
/*
把一个数组最开始的若干个元素搬到数组的末尾，我们称之为数组的旋转。
输入一个递增排序的数组的一个旋转，输出旋转数组的最小元素。
例如，数组 [3,4,5,1,2] 为 [1,2,3,4,5] 的一个旋转，该数组的最小值为1
 */
//经过观察,可以发现递增数组的旋转数组只需要遍历一遍
//第一次出现前一个数据>后一个数据,说明后一个就是最小的数
//这是一个O(n)的算法
//这个最好理解,也不难,最坏是O(n),最好O(1)
func minArray(numbers []int) int {
	if len(numbers) == 0 {
		return -1
	}

	if len(numbers) == 1 {
		return numbers[0]
	}

	for i := 0; i < len(numbers)-1; i++ {
		if numbers[i] > numbers[i+1] {
			return numbers[i+1]
		}
	}

	return numbers[0]
}

//解法2:如果出现排序数组中进行查找某个数字或者统计某个数字的次数,都可以考虑二分查找
//结合双指针的使用
//一般情况下是个O(logn)的算法,只有特殊情况时是个O(n)的算法
func minArray2(numbers []int) int {
	if len(numbers) == 0 {
		return -1
	}

	if len(numbers) == 1 {
		return numbers[0]
	}

	p1, p2 := 0, len(numbers)-1
	if numbers[p1] < numbers[p2] { //输入的是一个升序排序的数组
		return numbers[p1]
	}

	index := 0
	for {                         //输入的是升序排序的旋转
		if p1+1 == p2 {
			index = p2
			break
		}

		mid := (p1 + p2) / 2
		//特殊情况,无法判断中间元素在前面递增的子数组还是在后面递增的子数组
		if numbers[mid] == numbers[p1] && numbers[mid] == numbers[p2] {
			index = findMin(numbers)
			break
		}

		if numbers[mid] >= numbers[p1] {
			p1 = mid
		} else if numbers[mid] <= numbers[p2] {
			p2 = mid
		}
	}

	return numbers[index]
}

func findMin(numbers []int) int {
	index := 0
	for i := 1; i < len(numbers); i++ {
		if numbers[i] < numbers[index] {
			index = i
		}
	}
	return index
}

func main() {
	number1 := []int{2, 3, 4, 5, 1}
	number2 := []int{3, 4, 5, 1, 2}
	number3 := []int{4, 5, 1, 2, 3}
	min1 := minArray(number1)
	min2 := minArray(number2)
	min3 := minArray(number3)
	fmt.Printf("min1:%d, min2:%d, min3:%d\n", min1, min2, min3)

	min4 := minArray2(number1)
	min5 := minArray2(number2)
	min6 := minArray2(number3)
	fmt.Printf("min4:%d, min5:%d, min6:%d\n", min4, min5, min6)

	number4 := []int{1, 3, 5}         //这里个特殊测试用例:自身递增数组
	number5 := []int{1, 0, 1, 1, 1}   //这种特殊用例:{10, 1, 10, 10, 10}看成是{1,10,10,10,10}的旋转
	min7 := minArray2(number4)
	min8 := minArray2(number5)
	fmt.Printf("min7:%d, min8:%d\n", min7, min8)
}
