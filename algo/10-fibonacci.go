package main

import "fmt"

//算法和数据操作
//题目10:求斐波那契数列的第n项目
/*
  F(0) = 0,   F(1) = 1
  F(N) = F(N - 1) + F(N - 2), 其中 N > 1.
  斐波那契数列由 0 和 1 开始，之后的斐波那契数就是由之前的两数相加而得出。

  答案需要取模 1e9+7（1000000007），如计算初始结果为：1000000008，请返回 1
 */

//思路1:递归的最简单形式
//但是这样写会超时时间限制
func fib(n int) int {
	if n == 0 {
		return 0
	}

	if n == 1 {
		return 1
	}

	return (fib(n-1)+fib(n-2))%(1e9+7)
}

//采用循环
func fib2(n int) int {
	if n == 0 {
		return 0
	}

	if n == 1 {
		return 1
	}

	f0, f1 := 0, 1
	fn := 0
	for i := 2; i <= n; i++ {
		fn = (f0 + f1) % (1e9+7)
		f0 = f1
		f1 = fn
	}

	return fn
}

func main() {
	num2 := fib2(95)
	fmt.Printf("num:%d\n", num2)
}