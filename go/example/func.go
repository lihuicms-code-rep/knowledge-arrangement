package main

import (
	"fmt"
)

//变参函数
func sum(nums ...int) {
	fmt.Printf("num:%T, %v\n", nums, nums) //nums的类型就是:[]int
	total := 0
	for _, num := range nums {
		total += num
	}
	fmt.Printf("total:%d\n", total)
}

//闭包
//返回的是一个函数类型
//通过闭包来使用匿名函数

func intSeq() func() int {
	i := 0
	return func() int { //匿名函数
		i += 1
		return i
	}
}

func testVarParm() {
	sum(1, 2)
	sum(2)
	sum([]int{1, 2, 3, 4}...) //如果已经有了切片可以使用slice...
	arr := [4]int{1, 2, 3, 4}
	sum(arr[:]...)
}

func testCloure() {
	f := intSeq()
	fmt.Printf("f type:%T, value:%v\n", f, f)
	fmt.Println(f())
	fmt.Println(f())
	fmt.Println(f())
}

func main() {
	testCloure()
}
