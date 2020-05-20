package main

import "fmt"

//题目17:打印从1到最大的n位数
/*
输入数字 n，按顺序打印出从 1 到最大的 n 位十进制数。比如输入 3，则打印出 1、2、3 一直到最大的 3 位数 999。
需要考虑大数问题
 */
func print1ToMaxOfDigits(n int) {
	if n <= 0 {
		return
	}

	number := make([]byte, n) //数字用字符数组表示
	for i := 0; i < n; i++ {
		number[i] = '0'
	}

	for {
		overFlow := Incrment(number)
		if overFlow {
			break
		}
		PrintNumber(number)
	}
}

//在表示数字的字符切片上+1
//模拟整数加法,这个算法很重要
func Incrment(number []byte) bool {
	var nTakeOver byte = 0 //进位标识符
	isOverFlow := false    //是否溢出n位数的最大

	for i := len(number) - 1; i >= 0; i-- {
		nSum := number[i] - '0' + nTakeOver
		if i == len(number)-1 { //每次+1都是个位加1
			nSum++
		}

		//判断是否发生了进位
		if nSum >= 10 {
			if i == 0 {
				isOverFlow = true
				break
			} else {
				nSum -= 10
				nTakeOver = 1
				number[i] = nSum + '0'
			}
		} else {
			number[i] = nSum + '0'
			break //不涉及进位,末尾加一次就退出
		}
	}

	return isOverFlow
}

//打印字符切片表示的数字
//从前往后碰到第一个不为0的才开始打印
func PrintNumber(number []byte) {
	isBeginZero := true
	for i := 0; i < len(number); i++ {
		if isBeginZero && number[i] == '0' {
			continue
		}
		isBeginZero = false
		fmt.Printf("%c", number[i])
	}
	fmt.Println()
}

func main() {
	print1ToMaxOfDigits(3)
}
