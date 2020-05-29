package main

import "fmt"

//指针
//函数调用的堆栈情况

func zeroVal(iVal int) {
	fmt.Printf("in zeroVal iVal:%d, addr:%p\n", iVal, &iVal)
	iVal = 0
}

func zeroPtr(iPtr *int) {
	fmt.Printf("in zeroPtr iPtr:%p, addr:%p\n", iPtr, &iPtr)

	*iPtr = 0
}

func main() {
	i := 1
	fmt.Printf("****1**** i value:%d, addr:%p\n", i, &i)
	//zeroVal(i)
	//fmt.Printf("****2**** i value:%d, addr:%p\n", i, &i)

	zeroPtr(&i)
	fmt.Printf("****3**** i value:%d, addr:%p\n", i, &i)
}
