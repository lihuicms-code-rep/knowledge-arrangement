package main

import "fmt"

//题目3:数组中重复的数字
/*
在一个长度为n的数组里的所有数字都在0到 n-1 的范围内。
数组中某些数字是重复的，但不知道有几个数字是重复的。也不知道每个数字重复几次。
请找出数组中重复的数字。 例如，如果输入长度为7的数组{2,3,1,0,2,5,3}，那么对应的重复的数字2,3

//也可以引申为,在这样的数组中判断是否有重复的数字,都可以用解法3思考
 */

//思路1:暴力解法,遍历整个数组
func repeatNumberInArr(arr []int) {
	if len(arr) < 1 {
		return
	}

	repeatNumber := []int{}
	for i := 0; i < len(arr); i++ {
		for j := i; j < len(arr); j++ {
			if i != j && arr[i] == arr[j] {
				repeatNumber = append(repeatNumber, arr[i])
			}
		}
	}

	fmt.Printf("原数组:%v中重复的数字为:%v\n", arr, repeatNumber)
}

//思路2:可以使用hash map来存储,没有就存储,有就是重复的
//时间复杂度为O(n),空间复杂度为O(n)
func repeatNumberInArr2(arr []int) {
	if len(arr) < 1 {
		return
	}

	numbers := make(map[int]int, len(arr))
	repeatNums := make(map[int]int, len(arr))
	for _, v := range arr {
		if _, ok := numbers[v]; ok { //说明已经存在map中
			if _, ok := repeatNums[v]; !ok {
				repeatNums[v] = v
			}
		} else {
			numbers[v] = v
		}
	}

	fmt.Printf("原数组:%v中重复的数为:", arr)
	for _, v := range repeatNums {
		fmt.Printf("%d ", v)
	}
	fmt.Println()
}

//思路3:通过题干信息给出的限制,数组为n,每个元素的值的范围都是在0~n-1的范围
//假设这个数组中没有重复的数字,那么就是这样的数组[0, 1, 2, ..., n-1]
//假如有重复的数字,那么每个位置上的值就不是自己的下标
func repeatNumberInArr3(arr []int) {
	if !checkParam(arr) {
		fmt.Println("参数有误")
		return
	}

	for i := 0; i < len(arr); {
		m := arr[i]    //待扫描比较的数
		if m == i {    //如果该数和所在下标一致,说明在自己位置上
			i++
			continue
		}

		if m == arr[m] {   //重复的情况
			fmt.Printf("重复的数:%d\n", m)
			i++
			continue
		} else {
			arr[i], arr[m] = arr[m], arr[i]   //如果不等,就把m放到它本应该在的位置
			i = 0
		}
	}

}

//针对测试用例,需要检查参数
func checkParam(arr []int) bool {
	if arr == nil || len(arr) < 1 {
		return false
	}

	for _, v := range arr {
		if v < 0 || v > len(arr) - 1 {
			return false
		}
	}

	return true
}

func main() {
	//arr := []int{2, 3, 1, 0, 2, 5, 3}
	//repeatNumberInArr(arr)
	//arr2 := []int{2, 3, 1, 0, 2, 5, 3, 3, 4, 2}
	//repeatNumberInArr2(arr2)
	//arr3 := []int{2, 3, 1, 0, 2, 5, 3}
	//repeatNumberInArr3(arr3)

	arr4 := []int{2, 3, 1, 0, 4, 6}
	repeatNumberInArr3(arr4)
}
