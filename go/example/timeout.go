package main

import (
	"fmt"
	"time"
)

//超时处理

func main() {
	c1 := make(chan string)
	go func() {
		time.Sleep(2 * time.Second)
		c1 <- "result1"
	}()

	select {
	case msg := <- c1:
		fmt.Printf("recv c1:%s\n", msg)
	case <- time.After(1 * time.Second):
		fmt.Printf("timeout 1s\n")
	}

	fmt.Printf("main exit c1:%v\n", c1)
}
