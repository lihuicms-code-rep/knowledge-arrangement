package main

import (
	"fmt"
	"time"
)

//通道选择器
//同时等待多个通道操作
//协程,通道,选择器的结合是go的一个强大特性

func main() {
	c1 := make(chan string)
	c2 := make(chan string)

	go func() {
		time.Sleep(1 * time.Second)
		c1 <- "one"
	}()

	go func() {
		time.Sleep(2 * time.Second)
		c2 <- "two"
	}()

	i := 0
	for {
		select {
		case msg := <-c1:
			fmt.Printf("recv c1 msg:%s\n", msg)
		case msg := <-c2:
			fmt.Printf("recv c2 msg:%s\n", msg)
			goto label
		default:
			fmt.Println("can not recv any msg")
		    i++
		}
	}

label:
	fmt.Printf("main exit, c1:%v, c2:%v, i:%d\n", c1, c2, i)
}
