package main

import "fmt"

//题目15:二进制中1的个数
/*
  实现一个函数,输入一个整数,输出该数的二进制表示中1的个数
 */

//关于位运算需要知道
//与,或,异或,左移,右移
//左移m<<n,表示数m左移n位,只需要在m的末尾填上n个0
//右移m>>n,表示数m右移n位,如果是正数,在前边补充0,负数补充1

//有负数的情况还没解决

func bitOperation() {
	a := 10  //00001010
	b := -10 //10001010
	fmt.Printf("a二进制:%b\n", a)
	fmt.Printf("b二进制:%b\n", b)
	//注意在计算机中参数运算,要变成补码
	c := a << 2
	d := b << 2
	fmt.Printf("左移==>c二进制:%b, 十进制:%d\n", c, c)
	fmt.Printf("左移==>d二进制:%b, 十进制:%d\n", d, d)

	e := a >> 2
	f := b >> 2
	fmt.Printf("右移==>e二进制:%b, 十进制:%d\n", e, e)
	fmt.Printf("右移==>f二进制:%b, 十进制:%d\n", f, f)
}

//题目限定了是一个uint32类型的数,说明是>=0
func hammingWeight(num uint32) int {
	if num < 0 {
		return 0
	}
	count := 0
	for ; num > 0; num = num >> 1 {
		if num & 1 == 1 {
			count++
		}
	}
	return count
}


func main() {
	//bitOperation()
	var a uint32 = 10  //00001010
	count1 := hammingWeight(a)
	fmt.Printf("count1:%d\n", count1)
}
