package main

import (
	"fmt"
	"time"
)

//关于chan的一些理解
func main() {
	ch := make(chan int, 10)
	close(ch)
	val, ok := <-ch //从关闭的chan中读取不会panic,但是读出的是默认类型
	if !ok {
		fmt.Printf("chan close val:%v\n", val)
	}

	//典型用法
	//1.goroutine通信
	uch := make(chan int) //unBuffer chan
	go func(ch chan int) {
		fmt.Printf("in goroutine ch:%v\n", ch)
		ch <- 1
	}(uch)

	<-uch
	fmt.Println("main thread.......")

	//2.select,类似linux中IO多路复用,这里主要是提供了对多个channel的统一管理
	ch1 := make(chan string, 10)
	ch2 := make(chan string, 10)

	go func(ch1, ch2 chan string) {
		for i := 0; i < 10; i++ {
			ch1 <- fmt.Sprintf("chan1:%d", i)
			ch2 <- fmt.Sprintf("chan2:%d", i)
		}
		close(ch1)
		close(ch2)
		return
	}(ch1, ch2)

	for {
		select {
		case e, ok := <-ch1: //注意一下:select的break只能退到select这一层,不能跳出for,需要用label
			if !ok {
				goto exitLabel
			}
			fmt.Printf("recv ch1 content,%s\n", e)
		case e, ok := <-ch2:
			if !ok {
				goto exitLabel
			}
			fmt.Printf("recv ch2 content,%s\n", e)
		}
	}

exitLabel:
	fmt.Println("main thread2.......")

	//3.range chan
	//range chan直接从chan中取值,同时,一旦chan关闭,循环自动结束
	ch3 := make(chan string, 10)
	ch4 := make(chan int)
	producer := func(ch chan string) {
		for i := 0; i < 20; i++ {
			ch <- fmt.Sprintf("produce_%d", i)
		}
		close(ch)
		fmt.Println("producer goroutine return......")
		return
	}

	consumer := func(ch chan string) {
		for x := range ch {
			fmt.Printf("consume msg: %s\n", x)
		}
		fmt.Println("consumer goroutine return......")
		ch4 <- 1
		return
	}

	go producer(ch3)
	go consumer(ch3)

	<-ch4
	fmt.Println("main thread3.......")

	//4.超时控制
	ch5 := make(chan string, 10)
	ch6 := make(chan string, 10)

	go func(ch5, ch6 chan string) {
		for i := 0; i < 10; i++ {
			ch5 <- fmt.Sprintf("chan5:%d", i)
			ch6 <- fmt.Sprintf("chan6:%d", i)
		}
		//close(ch5)
		//close(ch6)
		return
	}(ch5, ch6)

	for {
		select {
		case e, ok := <-ch5: //注意一下:select的break只能退到select这一层,不能跳出for,需要用label
			if !ok {
				fmt.Println("ch5 can not recv msg")
				break
			}
			fmt.Printf("recv ch5 content,%s\n", e)
		case e, ok := <-ch6:
			if !ok {
				fmt.Println("ch6 can not recv msg")
				break
			}
			fmt.Printf("recv ch5 content,%s\n", e)
		case <-time.After(2 * time.Second):  //超时处理
			goto endLabel
		}

	}

endLabel:
	fmt.Println("end label exit....")

	//5.管道的方向,一般用在函数参数上限制


}

//只写
func chanWrite(ch chan<- string) {

}

//只读
func chanRead(chan <-chan string) {

}