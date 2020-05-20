package main

import (
	"fmt"
	"unsafe"
)

//go语言中的string,src/runtime/string.go
type stringStruct struct {
	str unsafe.Pointer //字符串首地址
	len int            //字符串长度
}

//[]byte和string之间的转换

//实现一个字符串Index方法
//t为非空串,若主串s中第pos个字符后存在于t相等的子串,
//则返回第一个这样的子串再s中的位置
//当pos=0时,这也是朴素的模式匹配算法

//时间复杂度分析,设主串s长度m,子串t长度为n
//最好情况,比如s:googlesss, t:google,那么O(1)内即可完成
//最坏情况,比如s:000000000000000000000000001(m-1个0,1个1)
//           t:0000001(n-1个0,1个1)
//每次比较到最后一个才发现不对,那么每次比较n次,需要比较m-n+1次,共O((m-n+1)*n)

//它的过程就是主串s的指针不断回溯的
func StrIndex(s, t string, pos int) int {
	if s == "" || t == "" {
		return -1
	}

	if len(s) < len(t) {
		return -2
	}

	sStart := pos      //t在主串s中第一次匹配到的下标
	sIndex := pos     //s串扫描指针
	tIndex := 0       //t串扫描指针
	tScanEnd := false //只有将t串扫描匹配完才算成功

	for {
		//如果扫描到t串末尾,则匹配成功
		if tIndex == len(t) {
			tScanEnd = true
			break
		}

		//如果扫描到s串末尾,则不论匹配与否都结束
		if sIndex == len(s) {
			break
		}

		if s[sIndex] == t[tIndex] {
			tIndex++
			sIndex++
		} else {
			sIndex = sStart + 1 //回溯
			sStart = sIndex     //记录s串扫描指针开始
			tIndex = 0          //如果不成立,t串匹配指针重回0
		}
	}

	if tScanEnd {
		return sStart
	} else {
		return -1
	}

}

//如何解决上面的重复0,1且最坏情况下的匹配问题
//KMP模式匹配算法

func main() {
	s := "googleokstq"
	t := "goog"
	index := StrIndex(s, t, 0)
	fmt.Println(index)

}
