package main

import "fmt"

//题目29:顺时针打印矩阵
/*
 输入一个矩阵，按照从外向里以顺时针的顺序依次打印出每一个数字。
示例 1：
{
  {1, 2, 3},
  {4, 5, 6},
  {7, 8, 9},
}
输入：matrix = [[1,2,3],[4,5,6],[7,8,9]]
输出：[1,2,3,6,9,8,7,4,5]
示例 2：
{
  {1, 2, 3, 4},
  {5, 6, 7, 8},
  {9, 10, 11, 12},
}
输入：matrix = [[1,2,3,4],[5,6,7,8],[9,10,11,12]]
输出：[1,2,3,4,8,12,11,10,9,5,6,7]
 */

//整个矩阵的遍历过程是这样的：
//从左到右,从上到下,从右到左,从下到上
//模拟整个过程,不断的缩小边界值即可

func spiralOrder(matrix [][]int) []int {
	if matrix == nil {
		return nil
	}

	row := len(matrix)
	if row < 1 {
		return nil
	}
	col := len(matrix[0])
	if col < 1 {
		return nil
	}

	res := make([]int, row*col)

	//四个边界值
	top, bottom := 0, row-1
	left, right := 0, col-1
	index := 0

	for {
		//先从左到右
		for i := left; i <= right; i++ {
			res[index] = matrix[top][i]
			index++
		}
		//一行遍历完成
		top++
		if top > bottom {
			break
		}


		//再从上往下
		for i := top; i <= bottom; i++ {
			res[index] = matrix[i][right]
			index++
		}
		//一列遍历完成
		right--
		if right < left {
			break
		}

		//再从右往左
		for i := right; i >= left; i-- {
			res[index] = matrix[bottom][i]
			index++
		}

		//一行遍历完成
		bottom--
		if bottom < top {
			break
		}

		//再从下往上
		for i := bottom; i >= top; i-- {
			res[index] = matrix[i][left]
			index++
		}

		//一列遍历完成
		left++
		if left > right {
			break
		}
	}

	return res
}

func main() {
	matrix := [][]int{
		{1, 2, 3, 4, 5},
		{6, 7, 8, 9, 10},
		{11, 12, 13, 14, 15},
		{16, 17, 18, 19, 20},
		{21, 22, 23, 24, 25},
	}

	res := spiralOrder(matrix)
	fmt.Printf("res:%v\n", res)
}
