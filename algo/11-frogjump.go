package main

import "fmt"

//题目11:青蛙跳台阶
/*
一只青蛙一次可以跳上1级台阶，也可以跳上2级台阶。求该青蛙跳上一个 n 级的台阶总共有多少种跳法。
答案需要取模 1e9+7（1000000007），如计算初始结果为：1000000008，请返回
//是斐波那契数列的变种

 */

func numWays(n int) int {
	if n == 0 || n == 1{
		return 1
	}

	if n == 2 {
		return 2
	}

	f1, f2 := 1, 2
	fn := 0
	for i := 3; i <= n; i++ {
		fn = (f1 + f2) % (1e9+7)
		f1 = f2
		f2 = fn
	}

	return fn
}

func main() {
	num := numWays(5)
	fmt.Printf("num:%d\n", num)

}