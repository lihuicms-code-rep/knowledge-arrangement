package main

import "fmt"

//通道缓冲

func main() {
	messages := make(chan string, 2) //带缓冲的通道,即使没有一个对应的并发接收方,也可以发送

	messages <- "1"
	messages <- "2"

	msg := <-messages
	fmt.Printf("msg:%s\n", msg)
}
