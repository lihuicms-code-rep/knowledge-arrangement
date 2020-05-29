package main

import "fmt"

//协程

func main() {
	messages := make(chan string) //默认通道是不带缓冲的,也就是<-chan要先准备好接收, chan<-才会发送

	go func() {
		messages <- "ok"
	}()

	msg := <-messages
	fmt.Printf("msg:%s\n", msg)
}
