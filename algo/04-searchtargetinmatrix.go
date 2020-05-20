package main

import "fmt"

//题目04:搜索二维矩阵
/*
问题描述:
编写一个高效的算法来搜索 m x n 矩阵 matrix 中的一个目标值 target。该矩阵具有以下特性：
每行的元素从左到右升序排列。
每列的元素从上到下升序排列。
示例:

现有矩阵 matrix 如下：

[
  [1,   4,  7, 11, 15],
  [2,   5,  8, 12, 19],
  [3,   6,  9, 16, 22],
  [10, 13, 14, 17, 24],
  [18, 21, 23, 26, 30]
]
给定 target = 5，返回 true。
给定 target = 20，返回 false。
 */

//1.暴力解法,遍历每个元素比较即可,时间复杂度O(m*n)
//2.可以在每一行进行二分查找,那这个时间复杂度为O(n*logn)
//3.比较高效的做法:由于二维数组是比较有特点的,每行从左到右递增,每列从上到下递增
//可以选择一个初始的位置,比如右上角的元素作为基准元素,进行比较(这个时间复杂度为O(n))

//2,3的算法在leetcode中显示,占用内存空间都差不多,2似乎还快点

//二分查找
//不存在则返回-1
func binarySearch(arr []int, target int) int {
	if arr == nil || len(arr) < 1 {
		return -1
	}
	left, right := 0, len(arr)-1
	index := -1

	for {
		mid := (left + right) / 2
		if left > right {
			break
		}

		if arr[mid] == target {
			index = mid
			break
		}

		if arr[mid] > target {
			right = mid - 1
		}

		if arr[mid] < target {
			left = mid + 1
		}
	}

	return index
}

func searchTargetInMatrix(matrix [][]int, target int) bool {
	row := len(matrix)
	if row < 1 {
		return false
	}

	col := len(matrix[0])
	if col < 1 {
		return false
	}

	for i := 0; i < row; i++ {
		index := binarySearch(matrix[i], target)
		if index != -1 {
			return true
		}
	}

	return false
}


func searchTargetInMatrix2(matrix [][]int, target int) bool {
	row := len(matrix)  //特别注意对数组的判断,比如传进来的是[],先判断row,有row后再判断col
	if row < 1 {
		return false
	}

	col := len(matrix[0])
	if col < 1 {
		return false
	}

	i, j := 0, col-1
	isFind := false
	for {
		if i > row-1 || j < 0 { //退出条件
			break
		}

		if matrix[i][j] == target { //找到
			isFind = true
			break
		}

		if matrix[i][j] < target { //右上角的小于target,说明i所在这一行数据都是小于target
			i++
			continue
		}

		if matrix[i][j] > target { //右上角的大于target,说明j所在这一列数据都大于target
			j--
			continue
		}

	}

	return isFind
}

func testSearch1() {
	//顺便复习一下多维切片的使用
	matrix := [][]int{
		{1, 4, 7, 11, 15},
		{2, 5, 8, 12, 19},
		{3, 6, 9, 16, 22},
		{10, 13, 14, 17, 24},
		{18, 21, 23, 26, 30},
	}

	var matrix1 = [][]int{
		{1, 4, 7, 11, 15},
		{2, 5, 8, 12, 19},
		{3, 6, 9, 16, 22},
		{10, 13, 14, 17, 24},
		{18, 21, 23, 26, 30},
	}

	var matrix2 [][]int = make([][]int, 5)
	matrix2[0] = []int{1, 4, 7, 11, 15}
	matrix2[1] = []int{2, 5, 8, 12, 19}
	matrix2[2] = []int{3, 6, 9, 16, 22}
	matrix2[3] = []int{10, 13, 14, 17, 24}
	matrix2[4] = []int{18, 21, 23, 26, 30}
	//matrix2[5] = []int{18, 21, 23, 26, 30},注意采用len后,超过部分需要用append来扩容

	matrix3 := make([][]int, 0, 5)
	matrix3 = append(matrix3, []int{1, 4, 7, 11, 15})
	matrix3 = append(matrix3, []int{2, 5, 8, 12, 19})
	matrix3 = append(matrix3, []int{3, 6, 9, 16, 22})
	matrix3 = append(matrix3, []int{10, 13, 14, 17, 24})
	matrix3 = append(matrix3, []int{18, 21, 23, 26, 30})

	fmt.Printf("%d exits in matrix:%t\n", 5, searchTargetInMatrix(matrix, 5))
	fmt.Printf("%d exits in matrix:%t\n", 44, searchTargetInMatrix(matrix1, 44))
}

func testSearch2() {
	//顺便复习一下多维切片的使用
	matrix := [][]int{
		{1, 4, 7, 11, 15},
		{2, 5, 8, 12, 19},
		{3, 6, 9, 16, 22},
		{10, 13, 14, 17, 24},
		{18, 21, 23, 26, 30},
	}

	var matrix1 = [][]int{
		{1, 4, 7, 11, 15},
		{2, 5, 8, 12, 19},
		{3, 6, 9, 16, 22},
		{10, 13, 14, 17, 24},
		{18, 21, 23, 26, 30},
	}

	var matrix2 [][]int = make([][]int, 5)
	matrix2[0] = []int{1, 4, 7, 11, 15}
	matrix2[1] = []int{2, 5, 8, 12, 19}
	matrix2[2] = []int{3, 6, 9, 16, 22}
	matrix2[3] = []int{10, 13, 14, 17, 24}
	matrix2[4] = []int{18, 21, 23, 26, 30}
	//matrix2[5] = []int{18, 21, 23, 26, 30},注意采用len后,超过部分需要用append来扩容

	matrix3 := make([][]int, 0, 5)
	matrix3 = append(matrix3, []int{1, 4, 7, 11, 15})
	matrix3 = append(matrix3, []int{2, 5, 8, 12, 19})
	matrix3 = append(matrix3, []int{3, 6, 9, 16, 22})
	matrix3 = append(matrix3, []int{10, 13, 14, 17, 24})
	matrix3 = append(matrix3, []int{18, 21, 23, 26, 30})

	fmt.Printf("%d exits in matrix:%t\n", 5, searchTargetInMatrix2(matrix, 5))
	fmt.Printf("%d exits in matrix:%t\n", 44, searchTargetInMatrix2(matrix1, 44))
}

func main() {
	testSearch2()
}
