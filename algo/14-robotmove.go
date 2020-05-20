package main

import "fmt"

//物体或者人在二维方格助攻运动,都可以按回溯法来结局
//题目14:机器人运动范围
/*
  地上有一个m行n列的方格，从坐标 [0,0] 到坐标 [m-1,n-1] 。
  一个机器人从坐标 [0, 0] 的格子开始移动，它每次可以向左、右、上、下移动一格（不能移动到方格外），
  也不能进入行坐标和列坐标的数位之和大于k的格子。
  例如，当k为18时，机器人能够进入方格 [35, 37] ，因为3+5+3+7=18。
  但它不能进入方格 [35, 38]，因为3+5+3+8=19。
  请问该机器人能够到达多少个格子？

  输入：m = 2, n = 3, k = 1
  输出：3

  输入：m = 3, n = 1, k = 0
  输出：1
 */

 //13和14题的套路基本一致
 //主要是内部函数的使用

func movingCount(m int, n int, k int) int {
	if m <= 0 || n <= 0 || k < 0 {
		return 0
	}

	if k == 0 {
		return 1
	}

	exists := make(map[string]bool, m*n)

   return movingCore(m, n, 0, 0, k, exists)
}

//从i行j列寻找
func movingCore(m,n,i,j,k int, exists map[string]bool) int {
	if i < 0 || j < 0 || i >= m || j >= n ||
		exists[fmt.Sprintf("%d:%d", i, j)]  ||
		getSum(i) + getSum(j) > k {
			return 0
	}
	exists[fmt.Sprintf("%d:%d", i, j)] = true

	sum := 1 + movingCore(m, n, i-1,j,k,exists) + movingCore(m, n, i+1,j,k,exists) +
		       movingCore(m, n, i,j-1,k,exists) + movingCore(m, n, i, j+1, k, exists)
	return sum
}

func getSum(num int) int {
	if num <= 0 {
		return 0
	}

	sum := 0
	for num > 0 {
		sum += num % 10
		num /= 10
	}

	return sum
}

func main() {
	count1 := movingCount(2, 3, 1)
	count2 := movingCount(3, 1, 0)
	fmt.Printf("count1:%d, count2:%d\n", count1, count2)
}
