package main

import "fmt"

//题目5:替换字符串中的空格
//请实现一个函数，将一个字符串中的每个空格替换成“%20”。例如，当字符串为We Are Happy.则经过替换之后的字符串为 We%20Are%20Happy

//思路:首先字符串是不可变对象,需要开辟额外空间
//1.直接扫描整个字符串,如果出现空格,则从空格后的一位开始,往后移动2位。
//然后从空格位置开始依次插入% 2 0三个字符
func replaceSpace(str string) string {
	space := 0
	for _, c := range str {
		if c == ' ' {
			space++
		}
	}

	if space == 0 {
		return str
	}

	str1 := make([]byte, len(str)+2*space)
	copy(str1, str)

	times := 0 //需要搬动的次数
	for i, c := range str1 {
		if c == ' ' {
			//往后移动2位
			for j := len(str) - 1 + times*2; j > i; j-- { //注意这里条件,需要画图观察
				str1[j+2] = str1[j]
			}
			str1[i] = '%'
			str1[i+1] = '2'
			str1[i+2] = '0'
			times++
		}
	}

	return string(str1)
}

//2.给出的建议算法为,用双指针从后往前同时扫描
//这样可以将O(n*n)的算法降到O(n)

func replaceSpace2(str string) string {
	space := 0
	for _, c := range str {
		if c == ' ' {
			space++
		}
	}

	if space == 0 {
		return str
	}

	str1 := make([]byte, len(str)+2*space)
	copy(str1, str)

	p1, p2 := len(str)-1, len(str1)-1 //p1,p2指针分别指向原字符串的末尾以及新字符串的末尾
	for {
		if p1 >= p2 || p1 < 0 || p2 < 0 { //结束条件
			break
		}

		if str[p1] == ' ' { //如果p1指针遇到了空格,此时就从p2指针开始往前依次填入0,2,%
			str1[p2] = '0'
			p2--
			str1[p2] = '2'
			p2--
			str1[p2] = '%'
			p2--
		} else { //如果没有遇到空格,则p1将其内容拷贝到str1
			str1[p2] = str[p1]
			p2--
		}
		p1--
	}

	return string(str1)
}

func main() {
	str := "we Are  Happy "
	//fmt.Printf("解法1,原字符串:%s替换空格后为:%s\n", str, replaceSpace(str))
	fmt.Printf("解法2,原字符串:%s替换空格后为:%s\n", str, replaceSpace2(str))
}
