package main

import (
	"fmt"
	"time"
)

//通道关闭

//关闭意味着不能再往这个通道发送值了
//同于通知

func main() {
	jobs := make(chan int, 5)
	done := make(chan bool)

	go func() {
		for i := 0; i < 5; i++ {
			jobs <- i
			fmt.Printf("sent %d to jobs\n", i)
			time.Sleep(1 * time.Second)
		}
		close(jobs)  //必须加上,忽然后面取值的时候会报错的

	}()

	go func() {
		for {
			if v, ok := <-jobs; ok {
				fmt.Printf("receive jobs msg:%d\n", v)
			} else {
				fmt.Println("receive all jobs msg")
				done <- true
				return
			}
		}
	}()

	fmt.Println("**************")
	<-done
	fmt.Println("main exit......")
}
