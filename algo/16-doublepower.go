package main

import "fmt"

//题目16:数值的整数次方
/*
实现函数double Power(double base, int exponent)，求base的exponent次方。
不得使用库函数，同时不需要考虑大数问题。
示例 1:

输入: 2.00000, 10
输出: 1024.00000
示例 2:

输入: 2.10000, 3
输出: 9.26100
示例 3:

输入: 2.00000, -2
输出: 0.25000
解释: 2-2 = 1/22 = 1/4 = 0.25
 */

func myPow(x float64, n int) float64 {
	if n == 0 {
		return 1
	}

	if n == 1 {
		return x
	}

	if n < 0 {
		x = 1 / x
		n = -n
	}

	tmp := myPow(x, n/2)   //将大问题转换为更小的问题
	if n & 1 == 0 {
		return tmp * tmp
	}

	return tmp*tmp*x
}

func pow2(x float64, n int) float64 {
	if n == 0 {
		return 1
	}

	if n == 1 {
		return x
	}

	if n&1 == 0 { //n为偶数
		return pow2(x, n>>1) * pow2(x, n>>1)
	} else {
		return x * pow2(x, (n-1)>>1) * pow2(x, (n-1)>>1)
	}
}

//耗时太长
func pow(x float64, n int) float64 {
	var result float64 = 1
	for i := 1; i <= n; i++ {
		result *= x
	}
	return result
}

func main() {
	res1 := myPow(2.000000, 10)
	res2 := myPow(-2.100000, -3)
	res3 := myPow(2.000000, -3)
	res4 := myPow(0.00001, 2147483647)
	fmt.Printf("res1:%f, res2:%f, res3:%f, res4:%f\n", res1, res2, res3, res4)
}
