package main

import "fmt"

//通道方向,指定通道的类型

//参数规定了通道的方向:只能将消息加入pings
func ping(pings chan<- string, msg string) {
	pings <- msg
	fmt.Printf("ping...msg:%s\n", msg)
}

//参数规定通道的方向:只能从pongs中接收
func pong(pongs chan<- string, pings <-chan string) {
	msg := <- pings
	pongs <- msg
	fmt.Printf("pong....msg:%s\n", msg)
}


func main() {
	pings := make(chan string, 1)
	pongs := make(chan string, 1)
	ping(pings, "1")
	pong(pongs, pings)
}